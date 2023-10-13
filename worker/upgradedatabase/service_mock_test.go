// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/upgradedatabase (interfaces: UpgradeService)

// Package upgradedatabase is a generated GoMock package.
package upgradedatabase

import (
	context "context"
	reflect "reflect"

	upgrade "github.com/juju/juju/core/upgrade"
	watcher "github.com/juju/juju/core/watcher"
	upgrade0 "github.com/juju/juju/domain/upgrade"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// ActiveUpgrade mocks base method.
func (m *MockUpgradeService) ActiveUpgrade(arg0 context.Context) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveUpgrade", arg0)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveUpgrade indicates an expected call of ActiveUpgrade.
func (mr *MockUpgradeServiceMockRecorder) ActiveUpgrade(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).ActiveUpgrade), arg0)
}

// CreateUpgrade mocks base method.
func (m *MockUpgradeService) CreateUpgrade(arg0 context.Context, arg1, arg2 version.Number) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpgrade indicates an expected call of CreateUpgrade.
func (mr *MockUpgradeServiceMockRecorder) CreateUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).CreateUpgrade), arg0, arg1, arg2)
}

// SetControllerDone mocks base method.
func (m *MockUpgradeService) SetControllerDone(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerDone", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerDone indicates an expected call of SetControllerDone.
func (mr *MockUpgradeServiceMockRecorder) SetControllerDone(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerDone", reflect.TypeOf((*MockUpgradeService)(nil).SetControllerDone), arg0, arg1, arg2)
}

// SetControllerReady mocks base method.
func (m *MockUpgradeService) SetControllerReady(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerReady indicates an expected call of SetControllerReady.
func (mr *MockUpgradeServiceMockRecorder) SetControllerReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerReady", reflect.TypeOf((*MockUpgradeService)(nil).SetControllerReady), arg0, arg1, arg2)
}

// SetDBUpgradeCompleted mocks base method.
func (m *MockUpgradeService) SetDBUpgradeCompleted(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeCompleted indicates an expected call of SetDBUpgradeCompleted.
func (mr *MockUpgradeServiceMockRecorder) SetDBUpgradeCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeCompleted", reflect.TypeOf((*MockUpgradeService)(nil).SetDBUpgradeCompleted), arg0, arg1)
}

// StartUpgrade mocks base method.
func (m *MockUpgradeService) StartUpgrade(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUpgrade indicates an expected call of StartUpgrade.
func (mr *MockUpgradeServiceMockRecorder) StartUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUpgrade", reflect.TypeOf((*MockUpgradeService)(nil).StartUpgrade), arg0, arg1)
}

// WatchForUpgradeReady mocks base method.
func (m *MockUpgradeService) WatchForUpgradeReady(arg0 context.Context, arg1 upgrade0.UUID) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForUpgradeReady", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForUpgradeReady indicates an expected call of WatchForUpgradeReady.
func (mr *MockUpgradeServiceMockRecorder) WatchForUpgradeReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForUpgradeReady", reflect.TypeOf((*MockUpgradeService)(nil).WatchForUpgradeReady), arg0, arg1)
}

// WatchForUpgradeState mocks base method.
func (m *MockUpgradeService) WatchForUpgradeState(arg0 context.Context, arg1 upgrade0.UUID, arg2 upgrade.State) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForUpgradeState", arg0, arg1, arg2)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForUpgradeState indicates an expected call of WatchForUpgradeState.
func (mr *MockUpgradeServiceMockRecorder) WatchForUpgradeState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForUpgradeState", reflect.TypeOf((*MockUpgradeService)(nil).WatchForUpgradeState), arg0, arg1, arg2)
}
