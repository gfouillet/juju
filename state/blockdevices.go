// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"reflect"

	"github.com/juju/errors"
	"github.com/juju/mgo/v3"
	"github.com/juju/mgo/v3/bson"
	"github.com/juju/mgo/v3/txn"
	"github.com/juju/names/v5"
	jujutxn "github.com/juju/txn/v3"
)

// BlockDevice represents the state of a block device in the model.
type BlockDevice interface {
	// Machine returns the ID of the machine the block device is attached to.
	Machine() string

	// Info returns the block device's BlockDeviceInfo.
	Info() BlockDeviceInfo
}

// blockDevicesDoc records information about a machine's block devices.
type blockDevicesDoc struct {
	DocID        string            `bson:"_id"`
	ModelUUID    string            `bson:"model-uuid"`
	Machine      string            `bson:"machineid"`
	BlockDevices []BlockDeviceInfo `bson:"blockdevices"`
}

// BlockDeviceInfo describes information about a block device.
type BlockDeviceInfo struct {
	DeviceName     string   `bson:"devicename"`
	DeviceLinks    []string `bson:"devicelinks,omitempty"`
	Label          string   `bson:"label,omitempty"`
	UUID           string   `bson:"uuid,omitempty"`
	HardwareId     string   `bson:"hardwareid,omitempty"`
	WWN            string   `bson:"wwn,omitempty"`
	BusAddress     string   `bson:"busaddress,omitempty"`
	Size           uint64   `bson:"size"`
	FilesystemType string   `bson:"fstype,omitempty"`
	InUse          bool     `bson:"inuse"`
	MountPoint     string   `bson:"mountpoint,omitempty"`
	SerialId       string   `bson:"serialid,omitempty"`
}

// WatchBlockDevices returns a new NotifyWatcher watching for
// changes to block devices associated with the specified machine.
func (sb *storageBackend) WatchBlockDevices(machine names.MachineTag) NotifyWatcher {
	return newBlockDevicesWatcher(sb.mb, machine.Id())
}

// BlockDevices returns the BlockDeviceInfo for the specified machine.
func (sb *storageBackend) BlockDevices(machine names.MachineTag) ([]BlockDeviceInfo, error) {
	return getBlockDevices(sb.mb.db(), machine.Id())
}

func getBlockDevices(db Database, machineId string) ([]BlockDeviceInfo, error) {
	coll, cleanup := db.GetCollection(blockDevicesC)
	defer cleanup()

	var d blockDevicesDoc
	err := coll.FindId(machineId).One(&d)
	if err == mgo.ErrNotFound {
		return nil, errors.NotFoundf("block devices not found for machine %q", machineId)
	} else if err != nil {
		return nil, errors.Annotate(err, "cannot get block device details")
	}
	return d.BlockDevices, nil
}

// setMachineBlockDevices updates the blockdevices collection with the
// currently attached block devices. Previously recorded block devices
// not in the list will be removed.
func setMachineBlockDevices(st modelBackend, machineId string, newInfo []BlockDeviceInfo) error {
	db := st.db()
	buildTxn := func(attempt int) ([]txn.Op, error) {
		oldInfo, err := getBlockDevices(db, machineId)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if !blockDevicesChanged(oldInfo, newInfo) {
			return nil, jujutxn.ErrNoOperations
		}
		ops := []txn.Op{{
			C:      machinesC,
			Id:     machineId,
			Assert: notDeadDoc,
		}, {
			C:      blockDevicesC,
			Id:     machineId,
			Assert: txn.DocExists,
			Update: bson.D{{"$set", bson.D{{"blockdevices", newInfo}}}},
		}}
		return ops, nil
	}
	return db.Run(buildTxn)
}

func createMachineBlockDevicesOp(machineId string) txn.Op {
	return txn.Op{
		C:      blockDevicesC,
		Id:     machineId,
		Insert: &blockDevicesDoc{Machine: machineId},
		Assert: txn.DocMissing,
	}
}

func removeMachineBlockDevicesOp(machineId string) txn.Op {
	return txn.Op{
		C:      blockDevicesC,
		Id:     machineId,
		Remove: true,
	}
}

func blockDevicesChanged(oldDevices, newDevices []BlockDeviceInfo) bool {
	if len(oldDevices) != len(newDevices) {
		return true
	}
	for _, o := range oldDevices {
		var found bool
		for _, n := range newDevices {
			if reflect.DeepEqual(o, n) {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}
	return false
}
