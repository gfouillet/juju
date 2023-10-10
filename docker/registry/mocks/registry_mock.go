// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/docker/registry/internal (interfaces: Registry)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	docker "github.com/juju/juju/docker"
	tools "github.com/juju/juju/tools"
	gomock "go.uber.org/mock/gomock"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRegistry) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// GetArchitecture mocks base method.
func (m *MockRegistry) GetArchitecture(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchitecture", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchitecture indicates an expected call of GetArchitecture.
func (mr *MockRegistryMockRecorder) GetArchitecture(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchitecture", reflect.TypeOf((*MockRegistry)(nil).GetArchitecture), arg0, arg1)
}

// ImageRepoDetails mocks base method.
func (m *MockRegistry) ImageRepoDetails() docker.ImageRepoDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageRepoDetails")
	ret0, _ := ret[0].(docker.ImageRepoDetails)
	return ret0
}

// ImageRepoDetails indicates an expected call of ImageRepoDetails.
func (mr *MockRegistryMockRecorder) ImageRepoDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageRepoDetails", reflect.TypeOf((*MockRegistry)(nil).ImageRepoDetails))
}

// Ping mocks base method.
func (m *MockRegistry) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRegistryMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRegistry)(nil).Ping))
}

// RefreshAuth mocks base method.
func (m *MockRegistry) RefreshAuth() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAuth")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshAuth indicates an expected call of RefreshAuth.
func (mr *MockRegistryMockRecorder) RefreshAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAuth", reflect.TypeOf((*MockRegistry)(nil).RefreshAuth))
}

// ShouldRefreshAuth mocks base method.
func (m *MockRegistry) ShouldRefreshAuth() (bool, time.Duration) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldRefreshAuth")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(time.Duration)
	return ret0, ret1
}

// ShouldRefreshAuth indicates an expected call of ShouldRefreshAuth.
func (mr *MockRegistryMockRecorder) ShouldRefreshAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldRefreshAuth", reflect.TypeOf((*MockRegistry)(nil).ShouldRefreshAuth))
}

// String mocks base method.
func (m *MockRegistry) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockRegistryMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockRegistry)(nil).String))
}

// Tags mocks base method.
func (m *MockRegistry) Tags(arg0 string) (tools.Versions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags", arg0)
	ret0, _ := ret[0].(tools.Versions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags.
func (mr *MockRegistryMockRecorder) Tags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockRegistry)(nil).Tags), arg0)
}
