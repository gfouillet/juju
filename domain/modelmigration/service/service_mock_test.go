// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/modelmigration/service (interfaces: InstanceProvider,ResourceProvider,State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination service_mock_test.go github.com/juju/juju/domain/modelmigration/service InstanceProvider,ResourceProvider,State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	set "github.com/juju/collections/set"
	envcontext "github.com/juju/juju/environs/envcontext"
	instances "github.com/juju/juju/environs/instances"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockInstanceProvider is a mock of InstanceProvider interface.
type MockInstanceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceProviderMockRecorder
}

// MockInstanceProviderMockRecorder is the mock recorder for MockInstanceProvider.
type MockInstanceProviderMockRecorder struct {
	mock *MockInstanceProvider
}

// NewMockInstanceProvider creates a new mock instance.
func NewMockInstanceProvider(ctrl *gomock.Controller) *MockInstanceProvider {
	mock := &MockInstanceProvider{ctrl: ctrl}
	mock.recorder = &MockInstanceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceProvider) EXPECT() *MockInstanceProviderMockRecorder {
	return m.recorder
}

// AllInstances mocks base method.
func (m *MockInstanceProvider) AllInstances(arg0 envcontext.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllInstances indicates an expected call of AllInstances.
func (mr *MockInstanceProviderMockRecorder) AllInstances(arg0 any) *MockInstanceProviderAllInstancesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllInstances", reflect.TypeOf((*MockInstanceProvider)(nil).AllInstances), arg0)
	return &MockInstanceProviderAllInstancesCall{Call: call}
}

// MockInstanceProviderAllInstancesCall wrap *gomock.Call
type MockInstanceProviderAllInstancesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceProviderAllInstancesCall) Return(arg0 []instances.Instance, arg1 error) *MockInstanceProviderAllInstancesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceProviderAllInstancesCall) Do(f func(envcontext.ProviderCallContext) ([]instances.Instance, error)) *MockInstanceProviderAllInstancesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceProviderAllInstancesCall) DoAndReturn(f func(envcontext.ProviderCallContext) ([]instances.Instance, error)) *MockInstanceProviderAllInstancesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockResourceProvider is a mock of ResourceProvider interface.
type MockResourceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockResourceProviderMockRecorder
}

// MockResourceProviderMockRecorder is the mock recorder for MockResourceProvider.
type MockResourceProviderMockRecorder struct {
	mock *MockResourceProvider
}

// NewMockResourceProvider creates a new mock instance.
func NewMockResourceProvider(ctrl *gomock.Controller) *MockResourceProvider {
	mock := &MockResourceProvider{ctrl: ctrl}
	mock.recorder = &MockResourceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceProvider) EXPECT() *MockResourceProviderMockRecorder {
	return m.recorder
}

// AdoptResources mocks base method.
func (m *MockResourceProvider) AdoptResources(arg0 envcontext.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockResourceProviderMockRecorder) AdoptResources(arg0, arg1, arg2 any) *MockResourceProviderAdoptResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockResourceProvider)(nil).AdoptResources), arg0, arg1, arg2)
	return &MockResourceProviderAdoptResourcesCall{Call: call}
}

// MockResourceProviderAdoptResourcesCall wrap *gomock.Call
type MockResourceProviderAdoptResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceProviderAdoptResourcesCall) Return(arg0 error) *MockResourceProviderAdoptResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceProviderAdoptResourcesCall) Do(f func(envcontext.ProviderCallContext, string, version.Number) error) *MockResourceProviderAdoptResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceProviderAdoptResourcesCall) DoAndReturn(f func(envcontext.ProviderCallContext, string, version.Number) error) *MockResourceProviderAdoptResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// GetAllInstanceIDs mocks base method.
func (m *MockState) GetAllInstanceIDs(arg0 context.Context) (set.Strings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInstanceIDs", arg0)
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInstanceIDs indicates an expected call of GetAllInstanceIDs.
func (mr *MockStateMockRecorder) GetAllInstanceIDs(arg0 any) *MockStateGetAllInstanceIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInstanceIDs", reflect.TypeOf((*MockState)(nil).GetAllInstanceIDs), arg0)
	return &MockStateGetAllInstanceIDsCall{Call: call}
}

// MockStateGetAllInstanceIDsCall wrap *gomock.Call
type MockStateGetAllInstanceIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetAllInstanceIDsCall) Return(arg0 set.Strings, arg1 error) *MockStateGetAllInstanceIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetAllInstanceIDsCall) Do(f func(context.Context) (set.Strings, error)) *MockStateGetAllInstanceIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetAllInstanceIDsCall) DoAndReturn(f func(context.Context) (set.Strings, error)) *MockStateGetAllInstanceIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetControllerUUID mocks base method.
func (m *MockState) GetControllerUUID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControllerUUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControllerUUID indicates an expected call of GetControllerUUID.
func (mr *MockStateMockRecorder) GetControllerUUID(arg0 any) *MockStateGetControllerUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControllerUUID", reflect.TypeOf((*MockState)(nil).GetControllerUUID), arg0)
	return &MockStateGetControllerUUIDCall{Call: call}
}

// MockStateGetControllerUUIDCall wrap *gomock.Call
type MockStateGetControllerUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetControllerUUIDCall) Return(arg0 string, arg1 error) *MockStateGetControllerUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetControllerUUIDCall) Do(f func(context.Context) (string, error)) *MockStateGetControllerUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetControllerUUIDCall) DoAndReturn(f func(context.Context) (string, error)) *MockStateGetControllerUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
