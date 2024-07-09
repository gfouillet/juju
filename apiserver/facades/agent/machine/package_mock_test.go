// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/machine (interfaces: NetworkService,MachineService)
//
// Generated by this command:
//
//	mockgen -typed -package machine_test -destination package_mock_test.go github.com/juju/juju/apiserver/facades/agent/machine NetworkService,MachineService
//

// Package machine_test is a generated GoMock package.
package machine_test

import (
	context "context"
	reflect "reflect"

	machine "github.com/juju/juju/core/machine"
	network "github.com/juju/juju/core/network"
	gomock "go.uber.org/mock/gomock"
)

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// AddSubnet mocks base method.
func (m *MockNetworkService) AddSubnet(arg0 context.Context, arg1 network.SubnetInfo) (network.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubnet", arg0, arg1)
	ret0, _ := ret[0].(network.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockNetworkServiceMockRecorder) AddSubnet(arg0, arg1 any) *MockNetworkServiceAddSubnetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockNetworkService)(nil).AddSubnet), arg0, arg1)
	return &MockNetworkServiceAddSubnetCall{Call: call}
}

// MockNetworkServiceAddSubnetCall wrap *gomock.Call
type MockNetworkServiceAddSubnetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceAddSubnetCall) Return(arg0 network.Id, arg1 error) *MockNetworkServiceAddSubnetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceAddSubnetCall) Do(f func(context.Context, network.SubnetInfo) (network.Id, error)) *MockNetworkServiceAddSubnetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceAddSubnetCall) DoAndReturn(f func(context.Context, network.SubnetInfo) (network.Id, error)) *MockNetworkServiceAddSubnetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllSubnets mocks base method.
func (m *MockNetworkService) GetAllSubnets(arg0 context.Context) (network.SubnetInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubnets", arg0)
	ret0, _ := ret[0].(network.SubnetInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubnets indicates an expected call of GetAllSubnets.
func (mr *MockNetworkServiceMockRecorder) GetAllSubnets(arg0 any) *MockNetworkServiceGetAllSubnetsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubnets", reflect.TypeOf((*MockNetworkService)(nil).GetAllSubnets), arg0)
	return &MockNetworkServiceGetAllSubnetsCall{Call: call}
}

// MockNetworkServiceGetAllSubnetsCall wrap *gomock.Call
type MockNetworkServiceGetAllSubnetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSubnetsCall) Return(arg0 network.SubnetInfos, arg1 error) *MockNetworkServiceGetAllSubnetsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSubnetsCall) Do(f func(context.Context) (network.SubnetInfos, error)) *MockNetworkServiceGetAllSubnetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSubnetsCall) DoAndReturn(f func(context.Context) (network.SubnetInfos, error)) *MockNetworkServiceGetAllSubnetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// EnsureDeadMachine mocks base method.
func (m *MockMachineService) EnsureDeadMachine(arg0 context.Context, arg1 machine.Name) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDeadMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureDeadMachine indicates an expected call of EnsureDeadMachine.
func (mr *MockMachineServiceMockRecorder) EnsureDeadMachine(arg0, arg1 any) *MockMachineServiceEnsureDeadMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDeadMachine", reflect.TypeOf((*MockMachineService)(nil).EnsureDeadMachine), arg0, arg1)
	return &MockMachineServiceEnsureDeadMachineCall{Call: call}
}

// MockMachineServiceEnsureDeadMachineCall wrap *gomock.Call
type MockMachineServiceEnsureDeadMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceEnsureDeadMachineCall) Return(arg0 error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceEnsureDeadMachineCall) Do(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceEnsureDeadMachineCall) DoAndReturn(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsController mocks base method.
func (m *MockMachineService) IsController(arg0 context.Context, arg1 machine.Name) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsController", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsController indicates an expected call of IsController.
func (mr *MockMachineServiceMockRecorder) IsController(arg0, arg1 any) *MockMachineServiceIsControllerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsController", reflect.TypeOf((*MockMachineService)(nil).IsController), arg0, arg1)
	return &MockMachineServiceIsControllerCall{Call: call}
}

// MockMachineServiceIsControllerCall wrap *gomock.Call
type MockMachineServiceIsControllerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceIsControllerCall) Return(arg0 bool, arg1 error) *MockMachineServiceIsControllerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceIsControllerCall) Do(f func(context.Context, machine.Name) (bool, error)) *MockMachineServiceIsControllerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceIsControllerCall) DoAndReturn(f func(context.Context, machine.Name) (bool, error)) *MockMachineServiceIsControllerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
