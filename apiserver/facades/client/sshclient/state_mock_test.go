// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/sshclient (interfaces: Backend,Model,Broker,SSHMachine)
//
// Generated by this command:
//
//	mockgen -typed -package sshclient_test -destination state_mock_test.go github.com/juju/juju/apiserver/facades/client/sshclient Backend,Model,Broker,SSHMachine
//

// Package sshclient_test is a generated GoMock package.
package sshclient_test

import (
	context "context"
	reflect "reflect"

	sshclient "github.com/juju/juju/apiserver/facades/client/sshclient"
	network "github.com/juju/juju/core/network"
	cloudspec "github.com/juju/juju/environs/cloudspec"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// CloudSpec mocks base method.
func (m *MockBackend) CloudSpec(arg0 context.Context) (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSpec", arg0)
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec.
func (mr *MockBackendMockRecorder) CloudSpec(arg0 any) *MockBackendCloudSpecCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockBackend)(nil).CloudSpec), arg0)
	return &MockBackendCloudSpecCall{Call: call}
}

// MockBackendCloudSpecCall wrap *gomock.Call
type MockBackendCloudSpecCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendCloudSpecCall) Return(arg0 cloudspec.CloudSpec, arg1 error) *MockBackendCloudSpecCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendCloudSpecCall) Do(f func(context.Context) (cloudspec.CloudSpec, error)) *MockBackendCloudSpecCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendCloudSpecCall) DoAndReturn(f func(context.Context) (cloudspec.CloudSpec, error)) *MockBackendCloudSpecCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerTag mocks base method.
func (m *MockBackend) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockBackendMockRecorder) ControllerTag() *MockBackendControllerTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockBackend)(nil).ControllerTag))
	return &MockBackendControllerTagCall{Call: call}
}

// MockBackendControllerTagCall wrap *gomock.Call
type MockBackendControllerTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendControllerTagCall) Return(arg0 names.ControllerTag) *MockBackendControllerTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendControllerTagCall) Do(f func() names.ControllerTag) *MockBackendControllerTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendControllerTagCall) DoAndReturn(f func() names.ControllerTag) *MockBackendControllerTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineForEntity mocks base method.
func (m *MockBackend) GetMachineForEntity(arg0 string) (sshclient.SSHMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineForEntity", arg0)
	ret0, _ := ret[0].(sshclient.SSHMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineForEntity indicates an expected call of GetMachineForEntity.
func (mr *MockBackendMockRecorder) GetMachineForEntity(arg0 any) *MockBackendGetMachineForEntityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineForEntity", reflect.TypeOf((*MockBackend)(nil).GetMachineForEntity), arg0)
	return &MockBackendGetMachineForEntityCall{Call: call}
}

// MockBackendGetMachineForEntityCall wrap *gomock.Call
type MockBackendGetMachineForEntityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendGetMachineForEntityCall) Return(arg0 sshclient.SSHMachine, arg1 error) *MockBackendGetMachineForEntityCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendGetMachineForEntityCall) Do(f func(string) (sshclient.SSHMachine, error)) *MockBackendGetMachineForEntityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendGetMachineForEntityCall) DoAndReturn(f func(string) (sshclient.SSHMachine, error)) *MockBackendGetMachineForEntityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSSHHostKeys mocks base method.
func (m *MockBackend) GetSSHHostKeys(arg0 names.MachineTag) (state.SSHHostKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSHHostKeys", arg0)
	ret0, _ := ret[0].(state.SSHHostKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSSHHostKeys indicates an expected call of GetSSHHostKeys.
func (mr *MockBackendMockRecorder) GetSSHHostKeys(arg0 any) *MockBackendGetSSHHostKeysCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSHHostKeys", reflect.TypeOf((*MockBackend)(nil).GetSSHHostKeys), arg0)
	return &MockBackendGetSSHHostKeysCall{Call: call}
}

// MockBackendGetSSHHostKeysCall wrap *gomock.Call
type MockBackendGetSSHHostKeysCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendGetSSHHostKeysCall) Return(arg0 state.SSHHostKeys, arg1 error) *MockBackendGetSSHHostKeysCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendGetSSHHostKeysCall) Do(f func(names.MachineTag) (state.SSHHostKeys, error)) *MockBackendGetSSHHostKeysCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendGetSSHHostKeysCall) DoAndReturn(f func(names.MachineTag) (state.SSHHostKeys, error)) *MockBackendGetSSHHostKeysCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockBackend) Model() (sshclient.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(sshclient.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockBackendMockRecorder) Model() *MockBackendModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockBackend)(nil).Model))
	return &MockBackendModelCall{Call: call}
}

// MockBackendModelCall wrap *gomock.Call
type MockBackendModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendModelCall) Return(arg0 sshclient.Model, arg1 error) *MockBackendModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendModelCall) Do(f func() (sshclient.Model, error)) *MockBackendModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendModelCall) DoAndReturn(f func() (sshclient.Model, error)) *MockBackendModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelTag mocks base method.
func (m *MockBackend) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockBackendMockRecorder) ModelTag() *MockBackendModelTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockBackend)(nil).ModelTag))
	return &MockBackendModelTagCall{Call: call}
}

// MockBackendModelTagCall wrap *gomock.Call
type MockBackendModelTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendModelTagCall) Return(arg0 names.ModelTag) *MockBackendModelTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendModelTagCall) Do(f func() names.ModelTag) *MockBackendModelTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendModelTagCall) DoAndReturn(f func() names.ModelTag) *MockBackendModelTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// ControllerUUID mocks base method.
func (m *MockModel) ControllerUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerUUID indicates an expected call of ControllerUUID.
func (mr *MockModelMockRecorder) ControllerUUID() *MockModelControllerUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerUUID", reflect.TypeOf((*MockModel)(nil).ControllerUUID))
	return &MockModelControllerUUIDCall{Call: call}
}

// MockModelControllerUUIDCall wrap *gomock.Call
type MockModelControllerUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelControllerUUIDCall) Return(arg0 string) *MockModelControllerUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelControllerUUIDCall) Do(f func() string) *MockModelControllerUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelControllerUUIDCall) DoAndReturn(f func() string) *MockModelControllerUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *MockModelTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
	return &MockModelTypeCall{Call: call}
}

// MockModelTypeCall wrap *gomock.Call
type MockModelTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelTypeCall) Return(arg0 state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelTypeCall) Do(f func() state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelTypeCall) DoAndReturn(f func() state.ModelType) *MockModelTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// GetSecretToken mocks base method.
func (m *MockBroker) GetSecretToken(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretToken indicates an expected call of GetSecretToken.
func (mr *MockBrokerMockRecorder) GetSecretToken(arg0, arg1 any) *MockBrokerGetSecretTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretToken", reflect.TypeOf((*MockBroker)(nil).GetSecretToken), arg0, arg1)
	return &MockBrokerGetSecretTokenCall{Call: call}
}

// MockBrokerGetSecretTokenCall wrap *gomock.Call
type MockBrokerGetSecretTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBrokerGetSecretTokenCall) Return(arg0 string, arg1 error) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBrokerGetSecretTokenCall) Do(f func(context.Context, string) (string, error)) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBrokerGetSecretTokenCall) DoAndReturn(f func(context.Context, string) (string, error)) *MockBrokerGetSecretTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockSSHMachine is a mock of SSHMachine interface.
type MockSSHMachine struct {
	ctrl     *gomock.Controller
	recorder *MockSSHMachineMockRecorder
}

// MockSSHMachineMockRecorder is the mock recorder for MockSSHMachine.
type MockSSHMachineMockRecorder struct {
	mock *MockSSHMachine
}

// NewMockSSHMachine creates a new mock instance.
func NewMockSSHMachine(ctrl *gomock.Controller) *MockSSHMachine {
	mock := &MockSSHMachine{ctrl: ctrl}
	mock.recorder = &MockSSHMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHMachine) EXPECT() *MockSSHMachineMockRecorder {
	return m.recorder
}

// Addresses mocks base method.
func (m *MockSSHMachine) Addresses() network.SpaceAddresses {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].(network.SpaceAddresses)
	return ret0
}

// Addresses indicates an expected call of Addresses.
func (mr *MockSSHMachineMockRecorder) Addresses() *MockSSHMachineAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockSSHMachine)(nil).Addresses))
	return &MockSSHMachineAddressesCall{Call: call}
}

// MockSSHMachineAddressesCall wrap *gomock.Call
type MockSSHMachineAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSSHMachineAddressesCall) Return(arg0 network.SpaceAddresses) *MockSSHMachineAddressesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSSHMachineAddressesCall) Do(f func() network.SpaceAddresses) *MockSSHMachineAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSSHMachineAddressesCall) DoAndReturn(f func() network.SpaceAddresses) *MockSSHMachineAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllDeviceSpaceAddresses mocks base method.
func (m *MockSSHMachine) AllDeviceSpaceAddresses(arg0 context.Context) (network.SpaceAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllDeviceSpaceAddresses", arg0)
	ret0, _ := ret[0].(network.SpaceAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDeviceSpaceAddresses indicates an expected call of AllDeviceSpaceAddresses.
func (mr *MockSSHMachineMockRecorder) AllDeviceSpaceAddresses(arg0 any) *MockSSHMachineAllDeviceSpaceAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDeviceSpaceAddresses", reflect.TypeOf((*MockSSHMachine)(nil).AllDeviceSpaceAddresses), arg0)
	return &MockSSHMachineAllDeviceSpaceAddressesCall{Call: call}
}

// MockSSHMachineAllDeviceSpaceAddressesCall wrap *gomock.Call
type MockSSHMachineAllDeviceSpaceAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSSHMachineAllDeviceSpaceAddressesCall) Return(arg0 network.SpaceAddresses, arg1 error) *MockSSHMachineAllDeviceSpaceAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSSHMachineAllDeviceSpaceAddressesCall) Do(f func(context.Context) (network.SpaceAddresses, error)) *MockSSHMachineAllDeviceSpaceAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSSHMachineAllDeviceSpaceAddressesCall) DoAndReturn(f func(context.Context) (network.SpaceAddresses, error)) *MockSSHMachineAllDeviceSpaceAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineTag mocks base method.
func (m *MockSSHMachine) MachineTag() names.MachineTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineTag")
	ret0, _ := ret[0].(names.MachineTag)
	return ret0
}

// MachineTag indicates an expected call of MachineTag.
func (mr *MockSSHMachineMockRecorder) MachineTag() *MockSSHMachineMachineTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineTag", reflect.TypeOf((*MockSSHMachine)(nil).MachineTag))
	return &MockSSHMachineMachineTagCall{Call: call}
}

// MockSSHMachineMachineTagCall wrap *gomock.Call
type MockSSHMachineMachineTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSSHMachineMachineTagCall) Return(arg0 names.MachineTag) *MockSSHMachineMachineTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSSHMachineMachineTagCall) Do(f func() names.MachineTag) *MockSSHMachineMachineTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSSHMachineMachineTagCall) DoAndReturn(f func() names.MachineTag) *MockSSHMachineMachineTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrivateAddress mocks base method.
func (m *MockSSHMachine) PrivateAddress() (network.SpaceAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateAddress")
	ret0, _ := ret[0].(network.SpaceAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateAddress indicates an expected call of PrivateAddress.
func (mr *MockSSHMachineMockRecorder) PrivateAddress() *MockSSHMachinePrivateAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateAddress", reflect.TypeOf((*MockSSHMachine)(nil).PrivateAddress))
	return &MockSSHMachinePrivateAddressCall{Call: call}
}

// MockSSHMachinePrivateAddressCall wrap *gomock.Call
type MockSSHMachinePrivateAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSSHMachinePrivateAddressCall) Return(arg0 network.SpaceAddress, arg1 error) *MockSSHMachinePrivateAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSSHMachinePrivateAddressCall) Do(f func() (network.SpaceAddress, error)) *MockSSHMachinePrivateAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSSHMachinePrivateAddressCall) DoAndReturn(f func() (network.SpaceAddress, error)) *MockSSHMachinePrivateAddressCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PublicAddress mocks base method.
func (m *MockSSHMachine) PublicAddress() (network.SpaceAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicAddress")
	ret0, _ := ret[0].(network.SpaceAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicAddress indicates an expected call of PublicAddress.
func (mr *MockSSHMachineMockRecorder) PublicAddress() *MockSSHMachinePublicAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicAddress", reflect.TypeOf((*MockSSHMachine)(nil).PublicAddress))
	return &MockSSHMachinePublicAddressCall{Call: call}
}

// MockSSHMachinePublicAddressCall wrap *gomock.Call
type MockSSHMachinePublicAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSSHMachinePublicAddressCall) Return(arg0 network.SpaceAddress, arg1 error) *MockSSHMachinePublicAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSSHMachinePublicAddressCall) Do(f func() (network.SpaceAddress, error)) *MockSSHMachinePublicAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSSHMachinePublicAddressCall) DoAndReturn(f func() (network.SpaceAddress, error)) *MockSSHMachinePublicAddressCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
