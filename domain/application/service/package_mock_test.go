// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/application/service (interfaces: ApplicationState,CharmState,WatcherFactory)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/application/service ApplicationState,CharmState,WatcherFactory
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	changestream "github.com/juju/juju/core/changestream"
	charm "github.com/juju/juju/core/charm"
	watcher "github.com/juju/juju/core/watcher"
	application0 "github.com/juju/juju/domain/application"
	charm0 "github.com/juju/juju/domain/application/charm"
	storage "github.com/juju/juju/domain/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockApplicationState is a mock of ApplicationState interface.
type MockApplicationState struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStateMockRecorder
}

// MockApplicationStateMockRecorder is the mock recorder for MockApplicationState.
type MockApplicationStateMockRecorder struct {
	mock *MockApplicationState
}

// NewMockApplicationState creates a new mock instance.
func NewMockApplicationState(ctrl *gomock.Controller) *MockApplicationState {
	mock := &MockApplicationState{ctrl: ctrl}
	mock.recorder = &MockApplicationStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationState) EXPECT() *MockApplicationStateMockRecorder {
	return m.recorder
}

// AddUnits mocks base method.
func (m *MockApplicationState) AddUnits(arg0 context.Context, arg1 string, arg2 ...application0.AddUnitArg) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUnits", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUnits indicates an expected call of AddUnits.
func (mr *MockApplicationStateMockRecorder) AddUnits(arg0, arg1 any, arg2 ...any) *MockApplicationStateAddUnitsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockApplicationState)(nil).AddUnits), varargs...)
	return &MockApplicationStateAddUnitsCall{Call: call}
}

// MockApplicationStateAddUnitsCall wrap *gomock.Call
type MockApplicationStateAddUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateAddUnitsCall) Return(arg0 error) *MockApplicationStateAddUnitsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateAddUnitsCall) Do(f func(context.Context, string, ...application0.AddUnitArg) error) *MockApplicationStateAddUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateAddUnitsCall) DoAndReturn(f func(context.Context, string, ...application0.AddUnitArg) error) *MockApplicationStateAddUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateApplication mocks base method.
func (m *MockApplicationState) CreateApplication(arg0 context.Context, arg1 string, arg2 charm0.Charm, arg3 ...application0.AddUnitArg) (application.ID, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApplication", varargs...)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationStateMockRecorder) CreateApplication(arg0, arg1, arg2 any, arg3 ...any) *MockApplicationStateCreateApplicationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationState)(nil).CreateApplication), varargs...)
	return &MockApplicationStateCreateApplicationCall{Call: call}
}

// MockApplicationStateCreateApplicationCall wrap *gomock.Call
type MockApplicationStateCreateApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateCreateApplicationCall) Return(arg0 application.ID, arg1 error) *MockApplicationStateCreateApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateCreateApplicationCall) Do(f func(context.Context, string, charm0.Charm, ...application0.AddUnitArg) (application.ID, error)) *MockApplicationStateCreateApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateCreateApplicationCall) DoAndReturn(f func(context.Context, string, charm0.Charm, ...application0.AddUnitArg) (application.ID, error)) *MockApplicationStateCreateApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteApplication mocks base method.
func (m *MockApplicationState) DeleteApplication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockApplicationStateMockRecorder) DeleteApplication(arg0, arg1 any) *MockApplicationStateDeleteApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockApplicationState)(nil).DeleteApplication), arg0, arg1)
	return &MockApplicationStateDeleteApplicationCall{Call: call}
}

// MockApplicationStateDeleteApplicationCall wrap *gomock.Call
type MockApplicationStateDeleteApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateDeleteApplicationCall) Return(arg0 error) *MockApplicationStateDeleteApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateDeleteApplicationCall) Do(f func(context.Context, string) error) *MockApplicationStateDeleteApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateDeleteApplicationCall) DoAndReturn(f func(context.Context, string) error) *MockApplicationStateDeleteApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetStoragePoolByName mocks base method.
func (m *MockApplicationState) GetStoragePoolByName(arg0 context.Context, arg1 string) (storage.StoragePoolDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePoolByName", arg0, arg1)
	ret0, _ := ret[0].(storage.StoragePoolDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePoolByName indicates an expected call of GetStoragePoolByName.
func (mr *MockApplicationStateMockRecorder) GetStoragePoolByName(arg0, arg1 any) *MockApplicationStateGetStoragePoolByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolByName", reflect.TypeOf((*MockApplicationState)(nil).GetStoragePoolByName), arg0, arg1)
	return &MockApplicationStateGetStoragePoolByNameCall{Call: call}
}

// MockApplicationStateGetStoragePoolByNameCall wrap *gomock.Call
type MockApplicationStateGetStoragePoolByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateGetStoragePoolByNameCall) Return(arg0 storage.StoragePoolDetails, arg1 error) *MockApplicationStateGetStoragePoolByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateGetStoragePoolByNameCall) Do(f func(context.Context, string) (storage.StoragePoolDetails, error)) *MockApplicationStateGetStoragePoolByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateGetStoragePoolByNameCall) DoAndReturn(f func(context.Context, string) (storage.StoragePoolDetails, error)) *MockApplicationStateGetStoragePoolByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageDefaults mocks base method.
func (m *MockApplicationState) StorageDefaults(arg0 context.Context) (storage.StorageDefaults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageDefaults", arg0)
	ret0, _ := ret[0].(storage.StorageDefaults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageDefaults indicates an expected call of StorageDefaults.
func (mr *MockApplicationStateMockRecorder) StorageDefaults(arg0 any) *MockApplicationStateStorageDefaultsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageDefaults", reflect.TypeOf((*MockApplicationState)(nil).StorageDefaults), arg0)
	return &MockApplicationStateStorageDefaultsCall{Call: call}
}

// MockApplicationStateStorageDefaultsCall wrap *gomock.Call
type MockApplicationStateStorageDefaultsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateStorageDefaultsCall) Return(arg0 storage.StorageDefaults, arg1 error) *MockApplicationStateStorageDefaultsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateStorageDefaultsCall) Do(f func(context.Context) (storage.StorageDefaults, error)) *MockApplicationStateStorageDefaultsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateStorageDefaultsCall) DoAndReturn(f func(context.Context) (storage.StorageDefaults, error)) *MockApplicationStateStorageDefaultsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpsertApplication mocks base method.
func (m *MockApplicationState) UpsertApplication(arg0 context.Context, arg1 string, arg2 ...application0.AddUnitArg) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertApplication", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertApplication indicates an expected call of UpsertApplication.
func (mr *MockApplicationStateMockRecorder) UpsertApplication(arg0, arg1 any, arg2 ...any) *MockApplicationStateUpsertApplicationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertApplication", reflect.TypeOf((*MockApplicationState)(nil).UpsertApplication), varargs...)
	return &MockApplicationStateUpsertApplicationCall{Call: call}
}

// MockApplicationStateUpsertApplicationCall wrap *gomock.Call
type MockApplicationStateUpsertApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationStateUpsertApplicationCall) Return(arg0 error) *MockApplicationStateUpsertApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationStateUpsertApplicationCall) Do(f func(context.Context, string, ...application0.AddUnitArg) error) *MockApplicationStateUpsertApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationStateUpsertApplicationCall) DoAndReturn(f func(context.Context, string, ...application0.AddUnitArg) error) *MockApplicationStateUpsertApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCharmState is a mock of CharmState interface.
type MockCharmState struct {
	ctrl     *gomock.Controller
	recorder *MockCharmStateMockRecorder
}

// MockCharmStateMockRecorder is the mock recorder for MockCharmState.
type MockCharmStateMockRecorder struct {
	mock *MockCharmState
}

// NewMockCharmState creates a new mock instance.
func NewMockCharmState(ctrl *gomock.Controller) *MockCharmState {
	mock := &MockCharmState{ctrl: ctrl}
	mock.recorder = &MockCharmStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmState) EXPECT() *MockCharmStateMockRecorder {
	return m.recorder
}

// DeleteCharm mocks base method.
func (m *MockCharmState) DeleteCharm(arg0 context.Context, arg1 charm.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCharm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCharm indicates an expected call of DeleteCharm.
func (mr *MockCharmStateMockRecorder) DeleteCharm(arg0, arg1 any) *MockCharmStateDeleteCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCharm", reflect.TypeOf((*MockCharmState)(nil).DeleteCharm), arg0, arg1)
	return &MockCharmStateDeleteCharmCall{Call: call}
}

// MockCharmStateDeleteCharmCall wrap *gomock.Call
type MockCharmStateDeleteCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateDeleteCharmCall) Return(arg0 error) *MockCharmStateDeleteCharmCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateDeleteCharmCall) Do(f func(context.Context, charm.ID) error) *MockCharmStateDeleteCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateDeleteCharmCall) DoAndReturn(f func(context.Context, charm.ID) error) *MockCharmStateDeleteCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmActions mocks base method.
func (m *MockCharmState) GetCharmActions(arg0 context.Context, arg1 charm.ID) (charm0.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmActions", arg0, arg1)
	ret0, _ := ret[0].(charm0.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmActions indicates an expected call of GetCharmActions.
func (mr *MockCharmStateMockRecorder) GetCharmActions(arg0, arg1 any) *MockCharmStateGetCharmActionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmActions", reflect.TypeOf((*MockCharmState)(nil).GetCharmActions), arg0, arg1)
	return &MockCharmStateGetCharmActionsCall{Call: call}
}

// MockCharmStateGetCharmActionsCall wrap *gomock.Call
type MockCharmStateGetCharmActionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmActionsCall) Return(arg0 charm0.Actions, arg1 error) *MockCharmStateGetCharmActionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmActionsCall) Do(f func(context.Context, charm.ID) (charm0.Actions, error)) *MockCharmStateGetCharmActionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmActionsCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Actions, error)) *MockCharmStateGetCharmActionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmArchivePath mocks base method.
func (m *MockCharmState) GetCharmArchivePath(arg0 context.Context, arg1 charm.ID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmArchivePath", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmArchivePath indicates an expected call of GetCharmArchivePath.
func (mr *MockCharmStateMockRecorder) GetCharmArchivePath(arg0, arg1 any) *MockCharmStateGetCharmArchivePathCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmArchivePath", reflect.TypeOf((*MockCharmState)(nil).GetCharmArchivePath), arg0, arg1)
	return &MockCharmStateGetCharmArchivePathCall{Call: call}
}

// MockCharmStateGetCharmArchivePathCall wrap *gomock.Call
type MockCharmStateGetCharmArchivePathCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmArchivePathCall) Return(arg0 string, arg1 error) *MockCharmStateGetCharmArchivePathCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmArchivePathCall) Do(f func(context.Context, charm.ID) (string, error)) *MockCharmStateGetCharmArchivePathCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmArchivePathCall) DoAndReturn(f func(context.Context, charm.ID) (string, error)) *MockCharmStateGetCharmArchivePathCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmConfig mocks base method.
func (m *MockCharmState) GetCharmConfig(arg0 context.Context, arg1 charm.ID) (charm0.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmConfig", arg0, arg1)
	ret0, _ := ret[0].(charm0.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmConfig indicates an expected call of GetCharmConfig.
func (mr *MockCharmStateMockRecorder) GetCharmConfig(arg0, arg1 any) *MockCharmStateGetCharmConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmConfig", reflect.TypeOf((*MockCharmState)(nil).GetCharmConfig), arg0, arg1)
	return &MockCharmStateGetCharmConfigCall{Call: call}
}

// MockCharmStateGetCharmConfigCall wrap *gomock.Call
type MockCharmStateGetCharmConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmConfigCall) Return(arg0 charm0.Config, arg1 error) *MockCharmStateGetCharmConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmConfigCall) Do(f func(context.Context, charm.ID) (charm0.Config, error)) *MockCharmStateGetCharmConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmConfigCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Config, error)) *MockCharmStateGetCharmConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmIDByRevision mocks base method.
func (m *MockCharmState) GetCharmIDByRevision(arg0 context.Context, arg1 string, arg2 int) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmIDByRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmIDByRevision indicates an expected call of GetCharmIDByRevision.
func (mr *MockCharmStateMockRecorder) GetCharmIDByRevision(arg0, arg1, arg2 any) *MockCharmStateGetCharmIDByRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmIDByRevision", reflect.TypeOf((*MockCharmState)(nil).GetCharmIDByRevision), arg0, arg1, arg2)
	return &MockCharmStateGetCharmIDByRevisionCall{Call: call}
}

// MockCharmStateGetCharmIDByRevisionCall wrap *gomock.Call
type MockCharmStateGetCharmIDByRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmIDByRevisionCall) Return(arg0 charm.ID, arg1 error) *MockCharmStateGetCharmIDByRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmIDByRevisionCall) Do(f func(context.Context, string, int) (charm.ID, error)) *MockCharmStateGetCharmIDByRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmIDByRevisionCall) DoAndReturn(f func(context.Context, string, int) (charm.ID, error)) *MockCharmStateGetCharmIDByRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmLXDProfile mocks base method.
func (m *MockCharmState) GetCharmLXDProfile(arg0 context.Context, arg1 charm.ID) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmLXDProfile", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmLXDProfile indicates an expected call of GetCharmLXDProfile.
func (mr *MockCharmStateMockRecorder) GetCharmLXDProfile(arg0, arg1 any) *MockCharmStateGetCharmLXDProfileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmLXDProfile", reflect.TypeOf((*MockCharmState)(nil).GetCharmLXDProfile), arg0, arg1)
	return &MockCharmStateGetCharmLXDProfileCall{Call: call}
}

// MockCharmStateGetCharmLXDProfileCall wrap *gomock.Call
type MockCharmStateGetCharmLXDProfileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmLXDProfileCall) Return(arg0 []byte, arg1 error) *MockCharmStateGetCharmLXDProfileCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmLXDProfileCall) Do(f func(context.Context, charm.ID) ([]byte, error)) *MockCharmStateGetCharmLXDProfileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmLXDProfileCall) DoAndReturn(f func(context.Context, charm.ID) ([]byte, error)) *MockCharmStateGetCharmLXDProfileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmManifest mocks base method.
func (m *MockCharmState) GetCharmManifest(arg0 context.Context, arg1 charm.ID) (charm0.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmManifest", arg0, arg1)
	ret0, _ := ret[0].(charm0.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmManifest indicates an expected call of GetCharmManifest.
func (mr *MockCharmStateMockRecorder) GetCharmManifest(arg0, arg1 any) *MockCharmStateGetCharmManifestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmManifest", reflect.TypeOf((*MockCharmState)(nil).GetCharmManifest), arg0, arg1)
	return &MockCharmStateGetCharmManifestCall{Call: call}
}

// MockCharmStateGetCharmManifestCall wrap *gomock.Call
type MockCharmStateGetCharmManifestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmManifestCall) Return(arg0 charm0.Manifest, arg1 error) *MockCharmStateGetCharmManifestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmManifestCall) Do(f func(context.Context, charm.ID) (charm0.Manifest, error)) *MockCharmStateGetCharmManifestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmManifestCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Manifest, error)) *MockCharmStateGetCharmManifestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmMetadata mocks base method.
func (m *MockCharmState) GetCharmMetadata(arg0 context.Context, arg1 charm.ID) (charm0.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmMetadata", arg0, arg1)
	ret0, _ := ret[0].(charm0.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmMetadata indicates an expected call of GetCharmMetadata.
func (mr *MockCharmStateMockRecorder) GetCharmMetadata(arg0, arg1 any) *MockCharmStateGetCharmMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmMetadata", reflect.TypeOf((*MockCharmState)(nil).GetCharmMetadata), arg0, arg1)
	return &MockCharmStateGetCharmMetadataCall{Call: call}
}

// MockCharmStateGetCharmMetadataCall wrap *gomock.Call
type MockCharmStateGetCharmMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateGetCharmMetadataCall) Return(arg0 charm0.Metadata, arg1 error) *MockCharmStateGetCharmMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateGetCharmMetadataCall) Do(f func(context.Context, charm.ID) (charm0.Metadata, error)) *MockCharmStateGetCharmMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateGetCharmMetadataCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Metadata, error)) *MockCharmStateGetCharmMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsCharmAvailable mocks base method.
func (m *MockCharmState) IsCharmAvailable(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCharmAvailable", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCharmAvailable indicates an expected call of IsCharmAvailable.
func (mr *MockCharmStateMockRecorder) IsCharmAvailable(arg0, arg1 any) *MockCharmStateIsCharmAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCharmAvailable", reflect.TypeOf((*MockCharmState)(nil).IsCharmAvailable), arg0, arg1)
	return &MockCharmStateIsCharmAvailableCall{Call: call}
}

// MockCharmStateIsCharmAvailableCall wrap *gomock.Call
type MockCharmStateIsCharmAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateIsCharmAvailableCall) Return(arg0 bool, arg1 error) *MockCharmStateIsCharmAvailableCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateIsCharmAvailableCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsCharmAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateIsCharmAvailableCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsCharmAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsControllerCharm mocks base method.
func (m *MockCharmState) IsControllerCharm(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsControllerCharm", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsControllerCharm indicates an expected call of IsControllerCharm.
func (mr *MockCharmStateMockRecorder) IsControllerCharm(arg0, arg1 any) *MockCharmStateIsControllerCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsControllerCharm", reflect.TypeOf((*MockCharmState)(nil).IsControllerCharm), arg0, arg1)
	return &MockCharmStateIsControllerCharmCall{Call: call}
}

// MockCharmStateIsControllerCharmCall wrap *gomock.Call
type MockCharmStateIsControllerCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateIsControllerCharmCall) Return(arg0 bool, arg1 error) *MockCharmStateIsControllerCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateIsControllerCharmCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsControllerCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateIsControllerCharmCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsControllerCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsSubordinateCharm mocks base method.
func (m *MockCharmState) IsSubordinateCharm(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubordinateCharm", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubordinateCharm indicates an expected call of IsSubordinateCharm.
func (mr *MockCharmStateMockRecorder) IsSubordinateCharm(arg0, arg1 any) *MockCharmStateIsSubordinateCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubordinateCharm", reflect.TypeOf((*MockCharmState)(nil).IsSubordinateCharm), arg0, arg1)
	return &MockCharmStateIsSubordinateCharmCall{Call: call}
}

// MockCharmStateIsSubordinateCharmCall wrap *gomock.Call
type MockCharmStateIsSubordinateCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateIsSubordinateCharmCall) Return(arg0 bool, arg1 error) *MockCharmStateIsSubordinateCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateIsSubordinateCharmCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsSubordinateCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateIsSubordinateCharmCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateIsSubordinateCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReserveCharmRevision mocks base method.
func (m *MockCharmState) ReserveCharmRevision(arg0 context.Context, arg1 charm.ID, arg2 int) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveCharmRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveCharmRevision indicates an expected call of ReserveCharmRevision.
func (mr *MockCharmStateMockRecorder) ReserveCharmRevision(arg0, arg1, arg2 any) *MockCharmStateReserveCharmRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveCharmRevision", reflect.TypeOf((*MockCharmState)(nil).ReserveCharmRevision), arg0, arg1, arg2)
	return &MockCharmStateReserveCharmRevisionCall{Call: call}
}

// MockCharmStateReserveCharmRevisionCall wrap *gomock.Call
type MockCharmStateReserveCharmRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateReserveCharmRevisionCall) Return(arg0 charm.ID, arg1 error) *MockCharmStateReserveCharmRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateReserveCharmRevisionCall) Do(f func(context.Context, charm.ID, int) (charm.ID, error)) *MockCharmStateReserveCharmRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateReserveCharmRevisionCall) DoAndReturn(f func(context.Context, charm.ID, int) (charm.ID, error)) *MockCharmStateReserveCharmRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCharm mocks base method.
func (m *MockCharmState) SetCharm(arg0 context.Context, arg1 charm0.Charm, arg2 charm0.SetStateArgs) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockCharmStateMockRecorder) SetCharm(arg0, arg1, arg2 any) *MockCharmStateSetCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockCharmState)(nil).SetCharm), arg0, arg1, arg2)
	return &MockCharmStateSetCharmCall{Call: call}
}

// MockCharmStateSetCharmCall wrap *gomock.Call
type MockCharmStateSetCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateSetCharmCall) Return(arg0 charm.ID, arg1 error) *MockCharmStateSetCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateSetCharmCall) Do(f func(context.Context, charm0.Charm, charm0.SetStateArgs) (charm.ID, error)) *MockCharmStateSetCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateSetCharmCall) DoAndReturn(f func(context.Context, charm0.Charm, charm0.SetStateArgs) (charm.ID, error)) *MockCharmStateSetCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCharmAvailable mocks base method.
func (m *MockCharmState) SetCharmAvailable(arg0 context.Context, arg1 charm.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharmAvailable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharmAvailable indicates an expected call of SetCharmAvailable.
func (mr *MockCharmStateMockRecorder) SetCharmAvailable(arg0, arg1 any) *MockCharmStateSetCharmAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharmAvailable", reflect.TypeOf((*MockCharmState)(nil).SetCharmAvailable), arg0, arg1)
	return &MockCharmStateSetCharmAvailableCall{Call: call}
}

// MockCharmStateSetCharmAvailableCall wrap *gomock.Call
type MockCharmStateSetCharmAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateSetCharmAvailableCall) Return(arg0 error) *MockCharmStateSetCharmAvailableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateSetCharmAvailableCall) Do(f func(context.Context, charm.ID) error) *MockCharmStateSetCharmAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateSetCharmAvailableCall) DoAndReturn(f func(context.Context, charm.ID) error) *MockCharmStateSetCharmAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SupportsContainers mocks base method.
func (m *MockCharmState) SupportsContainers(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsContainers", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportsContainers indicates an expected call of SupportsContainers.
func (mr *MockCharmStateMockRecorder) SupportsContainers(arg0, arg1 any) *MockCharmStateSupportsContainersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsContainers", reflect.TypeOf((*MockCharmState)(nil).SupportsContainers), arg0, arg1)
	return &MockCharmStateSupportsContainersCall{Call: call}
}

// MockCharmStateSupportsContainersCall wrap *gomock.Call
type MockCharmStateSupportsContainersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmStateSupportsContainersCall) Return(arg0 bool, arg1 error) *MockCharmStateSupportsContainersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmStateSupportsContainersCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateSupportsContainersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmStateSupportsContainersCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockCharmStateSupportsContainersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockWatcherFactory is a mock of WatcherFactory interface.
type MockWatcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherFactoryMockRecorder
}

// MockWatcherFactoryMockRecorder is the mock recorder for MockWatcherFactory.
type MockWatcherFactoryMockRecorder struct {
	mock *MockWatcherFactory
}

// NewMockWatcherFactory creates a new mock instance.
func NewMockWatcherFactory(ctrl *gomock.Controller) *MockWatcherFactory {
	mock := &MockWatcherFactory{ctrl: ctrl}
	mock.recorder = &MockWatcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherFactory) EXPECT() *MockWatcherFactoryMockRecorder {
	return m.recorder
}

// NewUUIDsWatcher mocks base method.
func (m *MockWatcherFactory) NewUUIDsWatcher(arg0 string, arg1 changestream.ChangeType) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUIDsWatcher", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUUIDsWatcher indicates an expected call of NewUUIDsWatcher.
func (mr *MockWatcherFactoryMockRecorder) NewUUIDsWatcher(arg0, arg1 any) *MockWatcherFactoryNewUUIDsWatcherCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUIDsWatcher", reflect.TypeOf((*MockWatcherFactory)(nil).NewUUIDsWatcher), arg0, arg1)
	return &MockWatcherFactoryNewUUIDsWatcherCall{Call: call}
}

// MockWatcherFactoryNewUUIDsWatcherCall wrap *gomock.Call
type MockWatcherFactoryNewUUIDsWatcherCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherFactoryNewUUIDsWatcherCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherFactoryNewUUIDsWatcherCall) Do(f func(string, changestream.ChangeType) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherFactoryNewUUIDsWatcherCall) DoAndReturn(f func(string, changestream.ChangeType) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
