// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package machinemanager_test

import (
	stdcontext "context"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/juju/names/v5"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common"
	commonmocks "github.com/juju/juju/apiserver/common/mocks"
	"github.com/juju/juju/apiserver/common/storagecommon"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facades/client/machinemanager"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/machine"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/environs/config"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/storage"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/binarystorage"
	stateerrors "github.com/juju/juju/state/errors"
	coretesting "github.com/juju/juju/testing"
)

var _ = gc.Suite(&MachineManagerSuite{})

type MachineManagerSuite struct {
	authorizer *apiservertesting.FakeAuthorizer

	controllerConfigService *MockControllerConfigService
	machineService          *MockMachineService
	networkService          *MockNetworkService
}

func (s *MachineManagerSuite) SetUpTest(c *gc.C) {
	s.authorizer = &apiservertesting.FakeAuthorizer{Tag: names.NewUserTag("admin")}
}

func (s *MachineManagerSuite) TestNewMachineManagerAPINonClient(c *gc.C) {
	tag := names.NewUnitTag("mysql/0")
	s.authorizer = &apiservertesting.FakeAuthorizer{Tag: tag}

	ctrl := gomock.NewController(c)
	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.machineService = NewMockMachineService(ctrl)
	s.networkService = NewMockNetworkService(ctrl)

	_, err := machinemanager.NewMachineManagerAPI(
		s.controllerConfigService,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		machinemanager.ModelAuthorizer{
			Authorizer: s.authorizer,
			ModelTag:   names.ModelTag{},
		},
		apiservertesting.NoopModelCredentialInvalidatorGetter,
		common.NewResources(),
		nil,
		loggertesting.WrapCheckLog(c),
		s.networkService,
	)
	c.Assert(err, gc.ErrorMatches, "permission denied")
}

var _ = gc.Suite(&AddMachineManagerSuite{})

type AddMachineManagerSuite struct {
	authorizer    *apiservertesting.FakeAuthorizer
	st            *MockBackend
	storageAccess *MockStorageInterface
	pool          *MockPool
	api           *machinemanager.MachineManagerAPI
	model         *MockModel
	store         *MockObjectStore
	cloudService  *commonmocks.MockCloudService
	credService   *commonmocks.MockCredentialService

	controllerConfigService *MockControllerConfigService
	machineService          *MockMachineService
	networkService          *MockNetworkService
}

func (s *AddMachineManagerSuite) SetUpTest(c *gc.C) {
	s.authorizer = &apiservertesting.FakeAuthorizer{Tag: names.NewUserTag("admin")}
}

func (s *AddMachineManagerSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.pool = NewMockPool(ctrl)
	s.model = NewMockModel(ctrl)

	s.st = NewMockBackend(ctrl)
	s.st.EXPECT().GetBlockForType(state.ChangeBlock).Return(nil, false, nil).AnyTimes()

	s.storageAccess = NewMockStorageInterface(ctrl)
	s.cloudService = commonmocks.NewMockCloudService(ctrl)
	s.credService = commonmocks.NewMockCredentialService(ctrl)
	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.machineService = NewMockMachineService(ctrl)
	s.store = NewMockObjectStore(ctrl)
	s.networkService = NewMockNetworkService(ctrl)

	var err error
	s.api, err = machinemanager.NewMachineManagerAPI(
		s.controllerConfigService,
		s.st,
		s.cloudService,
		s.credService,
		s.machineService,
		s.store,
		nil,
		s.storageAccess,
		s.pool,
		machinemanager.ModelAuthorizer{
			Authorizer: s.authorizer,
		},
		apiservertesting.NoopModelCredentialInvalidatorGetter,
		common.NewResources(),
		nil,
		loggertesting.WrapCheckLog(c),
		s.networkService,
	)
	c.Assert(err, jc.ErrorIsNil)

	return ctrl
}

func (s *AddMachineManagerSuite) TestAddMachines(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	apiParams := make([]params.AddMachineParams, 2)
	for i := range apiParams {
		apiParams[i] = params.AddMachineParams{
			Base: &params.Base{Name: "ubuntu", Channel: "22.04"},
			Jobs: []model.MachineJob{model.JobHostUnits},
		}
	}
	apiParams[0].Disks = []storage.Directive{{Size: 1, Count: 2}, {Size: 2, Count: 1}}
	apiParams[1].Disks = []storage.Directive{{Size: 1, Count: 2, Pool: "three"}}

	m1 := NewMockMachine(ctrl)
	m1.EXPECT().Id().Return("666").AnyTimes()
	m2 := NewMockMachine(ctrl)
	m2.EXPECT().Id().Return("667/lxd/1").AnyTimes()

	s.st.EXPECT().AddOneMachine(state.MachineTemplate{
		Base: state.UbuntuBase("22.04"),
		Jobs: []state.MachineJob{state.JobHostUnits},
		Volumes: []state.HostVolumeParams{
			{
				Volume:     state.VolumeParams{Pool: "", Size: 1},
				Attachment: state.VolumeAttachmentParams{ReadOnly: false},
			},
			{
				Volume:     state.VolumeParams{Pool: "", Size: 1},
				Attachment: state.VolumeAttachmentParams{ReadOnly: false},
			},
			{
				Volume:     state.VolumeParams{Pool: "", Size: 2},
				Attachment: state.VolumeAttachmentParams{ReadOnly: false},
			},
		},
	}).Return(m1, nil)
	s.machineService.EXPECT().CreateMachine(gomock.Any(), machine.Name("666"))
	s.machineService.EXPECT().CreateMachine(gomock.Any(), machine.Name("667/lxd/1"))
	s.machineService.EXPECT().CreateMachine(gomock.Any(), machine.Name("667"))
	s.st.EXPECT().AddOneMachine(state.MachineTemplate{
		Base: state.UbuntuBase("22.04"),
		Jobs: []state.MachineJob{state.JobHostUnits},
		Volumes: []state.HostVolumeParams{
			{
				Volume:     state.VolumeParams{Pool: "three", Size: 1},
				Attachment: state.VolumeAttachmentParams{ReadOnly: false},
			},
			{
				Volume:     state.VolumeParams{Pool: "three", Size: 1},
				Attachment: state.VolumeAttachmentParams{ReadOnly: false},
			},
		},
	}).Return(m2, nil)
	s.networkService.EXPECT().GetAllSpaces(gomock.Any())

	machines, err := s.api.AddMachines(stdcontext.Background(), params.AddMachines{MachineParams: apiParams})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(machines.Machines, gc.HasLen, 2)
}

func (s *AddMachineManagerSuite) TestAddMachinesStateError(c *gc.C) {
	defer s.setup(c).Finish()

	s.st.EXPECT().AddOneMachine(gomock.Any()).Return(nil, errors.New("boom"))
	s.networkService.EXPECT().GetAllSpaces(gomock.Any())

	results, err := s.api.AddMachines(stdcontext.Background(), params.AddMachines{
		MachineParams: []params.AddMachineParams{{
			Base: &params.Base{Name: "ubuntu", Channel: "22.04"},
		}},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.AddMachinesResults{
		Machines: []params.AddMachinesResult{{
			Error: &params.Error{Message: "boom", Code: ""},
		}},
	})
}

var _ = gc.Suite(&DestroyMachineManagerSuite{})

type DestroyMachineManagerSuite struct {
	testing.CleanupSuite
	authorizer    *apiservertesting.FakeAuthorizer
	st            *MockBackend
	storageAccess *MockStorageInterface
	leadership    *MockLeadership
	store         *MockObjectStore
	cloudService  *commonmocks.MockCloudService
	credService   *commonmocks.MockCredentialService
	api           *machinemanager.MachineManagerAPI

	controllerConfigService *MockControllerConfigService
	machineService          *MockMachineService
	networkService          *MockNetworkService
}

func (s *DestroyMachineManagerSuite) SetUpTest(c *gc.C) {
	s.CleanupSuite.SetUpTest(c)
	s.authorizer = &apiservertesting.FakeAuthorizer{Tag: names.NewUserTag("admin")}
	s.PatchValue(&machinemanager.ClassifyDetachedStorage, mockedClassifyDetachedStorage)
}

func (s *DestroyMachineManagerSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.st = NewMockBackend(ctrl)
	s.st.EXPECT().GetBlockForType(state.RemoveBlock).Return(nil, false, nil).AnyTimes()
	s.st.EXPECT().GetBlockForType(state.ChangeBlock).Return(nil, false, nil).AnyTimes()

	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.machineService = NewMockMachineService(ctrl)
	s.store = NewMockObjectStore(ctrl)
	s.networkService = NewMockNetworkService(ctrl)

	s.storageAccess = NewMockStorageInterface(ctrl)
	s.storageAccess.EXPECT().VolumeAccess().Return(nil).AnyTimes()
	s.storageAccess.EXPECT().FilesystemAccess().Return(nil).AnyTimes()

	s.cloudService = commonmocks.NewMockCloudService(ctrl)
	s.credService = commonmocks.NewMockCredentialService(ctrl)
	s.leadership = NewMockLeadership(ctrl)

	var err error
	s.api, err = machinemanager.NewMachineManagerAPI(
		s.controllerConfigService,
		s.st,
		s.cloudService,
		s.credService,
		s.machineService,
		s.store,
		nil,
		s.storageAccess,
		nil,
		machinemanager.ModelAuthorizer{
			Authorizer: s.authorizer,
		},
		nil,
		nil,
		s.leadership,
		loggertesting.WrapCheckLog(c),
		s.networkService,
	)
	c.Assert(err, jc.ErrorIsNil)

	return ctrl
}

func (s *DestroyMachineManagerSuite) expectUnpinAppLeaders(id string) {
	machineTag := names.NewMachineTag(id)

	s.leadership.EXPECT().GetMachineApplicationNames(gomock.Any(), id).Return([]string{"foo-app-1"}, nil)
	s.leadership.EXPECT().UnpinApplicationLeadersByName(gomock.Any(), machineTag, []string{"foo-app-1"}).Return(params.PinApplicationsResults{}, nil)
}

func (s *DestroyMachineManagerSuite) expectDestroyMachine(ctrl *gomock.Controller, units []machinemanager.Unit, containers []string, attemptDestroy, keep, force bool) *MockMachine {
	machine := NewMockMachine(ctrl)
	if keep {
		machine.EXPECT().SetKeepInstance(true).Return(nil)
	}

	machine.EXPECT().Containers().Return(containers, nil)

	if units == nil {
		units = []machinemanager.Unit{
			s.expectDestroyUnit(ctrl, "foo/0", true, nil),
			s.expectDestroyUnit(ctrl, "foo/1", false, nil),
			s.expectDestroyUnit(ctrl, "foo/2", false, nil),
		}
	}
	machine.EXPECT().Units().Return(units, nil)

	if attemptDestroy {
		if force {
			machine.EXPECT().ForceDestroy(gomock.Any()).Return(nil)
		} else {
			if len(containers) > 0 {
				machine.EXPECT().Destroy(gomock.Any()).Return(stateerrors.NewHasContainersError("0", containers))
			} else if len(units) > 0 {
				machine.EXPECT().Destroy(gomock.Any()).Return(stateerrors.NewHasAssignedUnitsError("0", []string{"foo/0", "foo/1", "foo/2"}))
			} else {
				machine.EXPECT().Destroy(gomock.Any()).Return(nil)
			}
		}
	}
	return machine
}

func (s *DestroyMachineManagerSuite) expectDestroyUnit(ctrl *gomock.Controller, name string, hasStorage bool, retrievalErr error) *MockUnit {
	unitTag := names.NewUnitTag(name)
	unit := NewMockUnit(ctrl)
	unit.EXPECT().UnitTag().Return(unitTag).AnyTimes()
	if retrievalErr != nil {
		s.storageAccess.EXPECT().UnitStorageAttachments(unitTag).Return(nil, retrievalErr)
	} else if !hasStorage {
		s.storageAccess.EXPECT().UnitStorageAttachments(unitTag).Return([]state.StorageAttachment{}, nil)
	} else {
		s.storageAccess.EXPECT().UnitStorageAttachments(unitTag).Return([]state.StorageAttachment{
			s.expectDestroyStorage(ctrl, "disks/0", true),
			s.expectDestroyStorage(ctrl, "disks/1", false),
		}, nil)
	}
	return unit
}

func (s *DestroyMachineManagerSuite) expectDestroyStorage(ctrl *gomock.Controller, id string, detachable bool) *MockStorageAttachment {
	storageInstanceTag := names.NewStorageTag(id)
	storageAttachment := NewMockStorageAttachment(ctrl)
	storageAttachment.EXPECT().StorageInstance().Return(storageInstanceTag)

	storageInstance := NewMockStorageInstance(ctrl)
	storageInstance.EXPECT().StorageTag().Return(storageInstanceTag).AnyTimes()
	s.storageAccess.EXPECT().StorageInstance(storageInstanceTag).Return(storageInstance, nil)

	return storageAttachment
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineFailedAllStorageRetrieval(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	units := []machinemanager.Unit{
		s.expectDestroyUnit(ctrl, "foo/0", false, errors.New("kaboom")),
		s.expectDestroyUnit(ctrl, "foo/1", false, errors.New("kaboom")),
		s.expectDestroyUnit(ctrl, "foo/2", false, errors.New("kaboom")),
	}
	machine0 := s.expectDestroyMachine(ctrl, units, nil, false, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	noWait := 0 * time.Second
	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		MachineTags: []string{"machine-0"},
		MaxWait:     &noWait,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Error: apiservererrors.ServerError(errors.New("getting storage for unit foo/0: kaboom\ngetting storage for unit foo/1: kaboom\ngetting storage for unit foo/2: kaboom")),
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineFailedSomeUnitStorageRetrieval(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	units := []machinemanager.Unit{
		s.expectDestroyUnit(ctrl, "foo/0", false, nil),
		s.expectDestroyUnit(ctrl, "foo/1", false, errors.New("kaboom")),
		s.expectDestroyUnit(ctrl, "foo/2", false, nil),
	}
	machine0 := s.expectDestroyMachine(ctrl, units, nil, false, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	noWait := 0 * time.Second
	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		MachineTags: []string{"machine-0"},
		MaxWait:     &noWait,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Error: apiservererrors.ServerError(errors.New("getting storage for unit foo/1: kaboom")),
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineFailedSomeStorageRetrievalManyMachines(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.expectUnpinAppLeaders("1")

	units0 := []machinemanager.Unit{
		s.expectDestroyUnit(ctrl, "foo/1", false, errors.New("kaboom")),
	}
	machine0 := s.expectDestroyMachine(ctrl, units0, nil, false, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	units1 := []machinemanager.Unit{}
	machine1 := s.expectDestroyMachine(ctrl, units1, nil, true, false, false)
	s.st.EXPECT().Machine("1").Return(machine1, nil)

	noWait := 0 * time.Second
	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		MachineTags: []string{"machine-0", "machine-1"},
		MaxWait:     &noWait,
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{
			{Error: apiservererrors.ServerError(errors.New("getting storage for unit foo/1: kaboom"))},
			{Info: &params.DestroyMachineInfo{
				MachineId: "1",
			}},
		},
	})
}

func (s *DestroyMachineManagerSuite) TestForceDestroyMachineFailedSomeStorageRetrievalManyMachines(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.expectUnpinAppLeaders("0")
	s.expectUnpinAppLeaders("1")

	units0 := []machinemanager.Unit{
		s.expectDestroyUnit(ctrl, "foo/1", false, errors.New("kaboom")),
	}
	machine0 := s.expectDestroyMachine(ctrl, units0, nil, true, false, true)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	units1 := []machinemanager.Unit{
		s.expectDestroyUnit(ctrl, "bar/0", true, nil),
	}
	machine1 := s.expectDestroyMachine(ctrl, units1, nil, true, false, true)
	s.st.EXPECT().Machine("1").Return(machine1, nil)

	noWait := 0 * time.Second
	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		Force:       true,
		MachineTags: []string{"machine-0", "machine-1"},
		MaxWait:     &noWait,
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{
			{Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-1"},
				},
			}},
			{Info: &params.DestroyMachineInfo{
				MachineId: "1",
				DestroyedUnits: []params.Entity{
					{"unit-bar-0"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
			}},
		},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineDryRun(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	machine0 := s.expectDestroyMachine(ctrl, nil, nil, false, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		MachineTags: []string{"machine-0"},
		DryRun:      true,
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-0"},
					{"unit-foo-1"},
					{"unit-foo-2"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
			},
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineWithContainersDryRun(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	machine0 := s.expectDestroyMachine(ctrl, nil, []string{"0/lxd/0"}, false, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)
	container0 := s.expectDestroyMachine(ctrl, nil, nil, false, false, false)
	s.st.EXPECT().Machine("0/lxd/0").Return(container0, nil)

	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		MachineTags: []string{"machine-0"},
		DryRun:      true,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-0"},
					{"unit-foo-1"},
					{"unit-foo-2"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
				DestroyedContainers: []params.DestroyMachineResult{{
					Info: &params.DestroyMachineInfo{
						MachineId: "0/lxd/0",
						DestroyedUnits: []params.Entity{
							{"unit-foo-0"},
							{"unit-foo-1"},
							{"unit-foo-2"},
						},
						DetachedStorage: []params.Entity{
							{"storage-disks-0"},
						},
						DestroyedStorage: []params.Entity{
							{"storage-disks-1"},
						},
					},
				}},
			},
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineWithParamsNoWait(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.expectUnpinAppLeaders("0")

	machine0 := s.expectDestroyMachine(ctrl, nil, nil, true, true, true)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	noWait := 0 * time.Second
	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		Keep:        true,
		Force:       true,
		MachineTags: []string{"machine-0"},
		MaxWait:     &noWait,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-0"},
					{"unit-foo-1"},
					{"unit-foo-2"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
			},
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineWithParamsNilWait(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.expectUnpinAppLeaders("0")

	machine0 := s.expectDestroyMachine(ctrl, nil, nil, true, true, true)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		Keep:        true,
		Force:       true,
		MachineTags: []string{"machine-0"},
		// This will use max wait of system default for delay between cleanup operations.
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-0"},
					{"unit-foo-1"},
					{"unit-foo-2"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
			},
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineWithContainers(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.leadership.EXPECT().GetMachineApplicationNames(gomock.Any(), "0").Return([]string{"foo-app-1"}, nil)

	machine0 := s.expectDestroyMachine(ctrl, nil, []string{"0/lxd/0"}, true, false, false)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		Force:       false,
		MachineTags: []string{"machine-0"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Error: apiservererrors.ServerError(stateerrors.NewHasContainersError("0", []string{"0/lxd/0"})),
		}},
	})
}

func (s *DestroyMachineManagerSuite) TestDestroyMachineWithContainersWithForce(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.expectUnpinAppLeaders("0")

	s.expectUnpinAppLeaders("0/lxd/0")

	machine0 := s.expectDestroyMachine(ctrl, nil, []string{"0/lxd/0"}, true, false, true)
	s.st.EXPECT().Machine("0").Return(machine0, nil)
	container0 := s.expectDestroyMachine(ctrl, nil, nil, true, false, true)
	s.st.EXPECT().Machine("0/lxd/0").Return(container0, nil)

	results, err := s.api.DestroyMachineWithParams(stdcontext.Background(), params.DestroyMachinesParams{
		Force:       true,
		MachineTags: []string{"machine-0"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.DestroyMachineResults{
		Results: []params.DestroyMachineResult{{
			Info: &params.DestroyMachineInfo{
				MachineId: "0",
				DestroyedUnits: []params.Entity{
					{"unit-foo-0"},
					{"unit-foo-1"},
					{"unit-foo-2"},
				},
				DetachedStorage: []params.Entity{
					{"storage-disks-0"},
				},
				DestroyedStorage: []params.Entity{
					{"storage-disks-1"},
				},
				DestroyedContainers: []params.DestroyMachineResult{{
					Info: &params.DestroyMachineInfo{
						MachineId: "0/lxd/0",
						DestroyedUnits: []params.Entity{
							{"unit-foo-0"},
							{"unit-foo-1"},
							{"unit-foo-2"},
						},
						DetachedStorage: []params.Entity{
							{"storage-disks-0"},
						},
						DestroyedStorage: []params.Entity{
							{"storage-disks-1"},
						},
					},
				}},
			},
		}},
	})
}

// Alternate placing storage instaces in detached, then destroyed
func mockedClassifyDetachedStorage(
	_ storagecommon.VolumeAccess,
	_ storagecommon.FilesystemAccess,
	storage []state.StorageInstance,
) ([]params.Entity, []params.Entity, error) {
	destroyed := make([]params.Entity, 0)
	detached := make([]params.Entity, 0)
	for i, stor := range storage {
		if i%2 == 0 {
			detached = append(detached, params.Entity{stor.StorageTag().String()})
		} else {
			destroyed = append(destroyed, params.Entity{stor.StorageTag().String()})
		}
	}
	return destroyed, detached, nil
}

var _ = gc.Suite(&ProvisioningMachineManagerSuite{})

type ProvisioningMachineManagerSuite struct {
	authorizer   *apiservertesting.FakeAuthorizer
	st           *MockBackend
	ctrlSt       *MockControllerBackend
	pool         *MockPool
	model        *MockModel
	store        *MockObjectStore
	cloudService *commonmocks.MockCloudService
	credService  *commonmocks.MockCredentialService
	api          *machinemanager.MachineManagerAPI

	controllerConfigService *MockControllerConfigService
	machineService          *MockMachineService
	networkService          *MockNetworkService
}

func (s *ProvisioningMachineManagerSuite) SetUpTest(c *gc.C) {
	s.authorizer = &apiservertesting.FakeAuthorizer{Tag: names.NewUserTag("admin")}
}

func (s *ProvisioningMachineManagerSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.st = NewMockBackend(ctrl)

	s.ctrlSt = NewMockControllerBackend(ctrl)
	s.ctrlSt.EXPECT().ControllerTag().Return(coretesting.ControllerTag).AnyTimes()

	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil).AnyTimes()
	s.machineService = NewMockMachineService(ctrl)

	s.pool = NewMockPool(ctrl)
	s.pool.EXPECT().SystemState().Return(s.ctrlSt, nil).AnyTimes()

	s.model = NewMockModel(ctrl)
	s.model.EXPECT().UUID().Return("uuid").AnyTimes()
	s.model.EXPECT().ModelTag().Return(coretesting.ModelTag).AnyTimes()
	s.st.EXPECT().Model().Return(s.model, nil).AnyTimes()

	s.cloudService = commonmocks.NewMockCloudService(ctrl)
	s.credService = commonmocks.NewMockCredentialService(ctrl)
	s.networkService = NewMockNetworkService(ctrl)

	s.store = NewMockObjectStore(ctrl)

	var err error
	s.api, err = machinemanager.NewMachineManagerAPI(
		s.controllerConfigService,
		s.st,
		s.cloudService,
		s.credService,
		s.machineService,
		s.store,
		nil,
		nil,
		s.pool,
		machinemanager.ModelAuthorizer{
			Authorizer: s.authorizer,
		},
		apiservertesting.NoopModelCredentialInvalidatorGetter,
		common.NewResources(),
		nil,
		loggertesting.WrapCheckLog(c),
		s.networkService,
	)
	c.Assert(err, jc.ErrorIsNil)
	return ctrl
}

func (s *ProvisioningMachineManagerSuite) expectProvisioningMachine(ctrl *gomock.Controller, arch *string) *MockMachine {
	machine := NewMockMachine(ctrl)
	machine.EXPECT().Base().Return(state.Base{OS: "ubuntu", Channel: "20.04/stable"}).AnyTimes()
	machine.EXPECT().Tag().Return(names.NewMachineTag("0")).AnyTimes()
	machine.EXPECT().HardwareCharacteristics().Return(&instance.HardwareCharacteristics{Arch: arch}, nil)
	if arch != nil {
		machine.EXPECT().SetPassword(gomock.Any()).Return(nil)
	}

	return machine
}

func (s *ProvisioningMachineManagerSuite) expectProvisioningStorageCloser(ctrl *gomock.Controller) *MockStorageCloser {
	storageCloser := NewMockStorageCloser(ctrl)
	storageCloser.EXPECT().AllMetadata().Return([]binarystorage.Metadata{{
		Version: "2.6.6-ubuntu-amd64",
	}}, nil)
	storageCloser.EXPECT().Close().Return(nil)

	return storageCloser
}

func (s *ProvisioningMachineManagerSuite) TestProvisioningScript(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.model.EXPECT().Config().Return(config.New(config.UseDefaults, coretesting.FakeConfig().Merge(coretesting.Attrs{
		"agent-version":            "2.6.6",
		"enable-os-upgrade":        true,
		"enable-os-refresh-update": true,
	}))).Times(2)

	arch := "amd64"
	machine0 := s.expectProvisioningMachine(ctrl, &arch)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	storageCloser := s.expectProvisioningStorageCloser(ctrl)
	s.st.EXPECT().ToolsStorage(gomock.Any()).Return(storageCloser, nil)

	s.ctrlSt.EXPECT().APIHostPortsForAgents(gomock.Any()).Return([]network.SpaceHostPorts{{{
		SpaceAddress: network.NewSpaceAddress("0.2.4.6", network.WithScope(network.ScopeCloudLocal)),
		NetPort:      1,
	}}}, nil).Times(2)

	result, err := s.api.ProvisioningScript(stdcontext.Background(), params.ProvisioningScriptParams{
		MachineId: "0",
		Nonce:     "nonce",
	})
	c.Assert(err, jc.ErrorIsNil)
	scriptLines := strings.Split(result.Script, "\n")
	provisioningScriptLines := strings.Split(result.Script, "\n")
	c.Assert(scriptLines, gc.HasLen, len(provisioningScriptLines))
	for i, line := range scriptLines {
		if strings.Contains(line, "oldpassword") {
			continue
		}
		c.Assert(line, gc.Equals, provisioningScriptLines[i])
	}
}

func (s *ProvisioningMachineManagerSuite) TestProvisioningScriptNoArch(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.model.EXPECT().Config().Return(config.New(config.UseDefaults, coretesting.FakeConfig().Merge(coretesting.Attrs{
		"agent-version":            "2.6.6",
		"enable-os-upgrade":        false,
		"enable-os-refresh-update": false,
	})))

	machine0 := s.expectProvisioningMachine(ctrl, nil)
	s.st.EXPECT().Machine("0").Return(machine0, nil)
	_, err := s.api.ProvisioningScript(stdcontext.Background(), params.ProvisioningScriptParams{
		MachineId: "0",
		Nonce:     "nonce",
	})
	c.Assert(err, gc.ErrorMatches, `getting instance config: arch is not set for "machine-0"`)
}

func (s *ProvisioningMachineManagerSuite) TestProvisioningScriptDisablePackageCommands(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.model.EXPECT().Config().Return(config.New(config.UseDefaults, coretesting.FakeConfig().Merge(coretesting.Attrs{
		"agent-version":            "2.6.6",
		"enable-os-upgrade":        false,
		"enable-os-refresh-update": false,
	}))).Times(2)

	arch := "amd64"
	machine0 := s.expectProvisioningMachine(ctrl, &arch)
	s.st.EXPECT().Machine("0").Return(machine0, nil)

	storageCloser := s.expectProvisioningStorageCloser(ctrl)
	s.st.EXPECT().ToolsStorage(gomock.Any()).Return(storageCloser, nil)

	s.ctrlSt.EXPECT().APIHostPortsForAgents(gomock.Any()).Return([]network.SpaceHostPorts{{{
		SpaceAddress: network.NewSpaceAddress("0.2.4.6", network.WithScope(network.ScopeCloudLocal)),
		NetPort:      1,
	}}}, nil).Times(2)

	result, err := s.api.ProvisioningScript(stdcontext.Background(), params.ProvisioningScriptParams{
		MachineId: "0",
		Nonce:     "nonce",
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result.Script, gc.Not(jc.Contains), "apt-get update")
	c.Assert(result.Script, gc.Not(jc.Contains), "apt-get upgrade")
}

type statusMatcher struct {
	c        *gc.C
	expected status.StatusInfo
}

func (m statusMatcher) Matches(x interface{}) bool {
	obtained, ok := x.(status.StatusInfo)
	m.c.Assert(ok, jc.IsTrue)
	if !ok {
		return false
	}

	m.c.Assert(obtained.Since, gc.NotNil)
	obtained.Since = nil
	m.c.Assert(obtained, jc.DeepEquals, m.expected)
	return true
}

func (m statusMatcher) String() string {
	return "Match the status.StatusInfo value"
}

func (s *ProvisioningMachineManagerSuite) TestRetryProvisioning(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.st.EXPECT().GetBlockForType(state.ChangeBlock).Return(nil, false, nil).AnyTimes()

	machine0 := NewMockMachine(ctrl)
	machine0.EXPECT().Id().Return("0")
	machine0.EXPECT().InstanceStatus().Return(status.StatusInfo{Status: "provisioning error"}, nil)
	machine0.EXPECT().SetInstanceStatus(statusMatcher{c: c, expected: status.StatusInfo{
		Status: status.ProvisioningError,
		Data:   map[string]interface{}{"transient": true},
	}}).Return(nil)
	machine1 := NewMockMachine(ctrl)
	machine1.EXPECT().Id().Return("1")
	s.st.EXPECT().AllMachines().Return([]machinemanager.Machine{machine0, machine1}, nil)

	results, err := s.api.RetryProvisioning(stdcontext.Background(), params.RetryProvisioningArgs{
		Machines: []string{"machine-0"},
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.ErrorResults{})
}

func (s *ProvisioningMachineManagerSuite) TestRetryProvisioningAll(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.st.EXPECT().GetBlockForType(state.ChangeBlock).Return(nil, false, nil).AnyTimes()

	machine0 := NewMockMachine(ctrl)
	machine0.EXPECT().InstanceStatus().Return(status.StatusInfo{Status: "provisioning error"}, nil)
	machine0.EXPECT().SetInstanceStatus(statusMatcher{c: c, expected: status.StatusInfo{
		Status: status.ProvisioningError,
		Data:   map[string]interface{}{"transient": true},
	}}).Return(nil)
	machine1 := NewMockMachine(ctrl)
	machine1.EXPECT().InstanceStatus().Return(status.StatusInfo{Status: "pending"}, nil)
	s.st.EXPECT().AllMachines().Return([]machinemanager.Machine{machine0, machine1}, nil)

	results, err := s.api.RetryProvisioning(stdcontext.Background(), params.RetryProvisioningArgs{
		All: true,
	})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.ErrorResults{})
}
