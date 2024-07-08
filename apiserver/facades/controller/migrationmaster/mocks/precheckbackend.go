// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/migration (interfaces: PrecheckBackend)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/precheckbackend.go github.com/juju/juju/internal/migration PrecheckBackend
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	migration "github.com/juju/juju/internal/migration"
	state "github.com/juju/juju/state"
	replicaset "github.com/juju/replicaset/v3"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockPrecheckBackend is a mock of PrecheckBackend interface.
type MockPrecheckBackend struct {
	ctrl     *gomock.Controller
	recorder *MockPrecheckBackendMockRecorder
}

// MockPrecheckBackendMockRecorder is the mock recorder for MockPrecheckBackend.
type MockPrecheckBackendMockRecorder struct {
	mock *MockPrecheckBackend
}

// NewMockPrecheckBackend creates a new mock instance.
func NewMockPrecheckBackend(ctrl *gomock.Controller) *MockPrecheckBackend {
	mock := &MockPrecheckBackend{ctrl: ctrl}
	mock.recorder = &MockPrecheckBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrecheckBackend) EXPECT() *MockPrecheckBackendMockRecorder {
	return m.recorder
}

// AgentVersion mocks base method.
func (m *MockPrecheckBackend) AgentVersion() (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion.
func (mr *MockPrecheckBackendMockRecorder) AgentVersion() *MockPrecheckBackendAgentVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockPrecheckBackend)(nil).AgentVersion))
	return &MockPrecheckBackendAgentVersionCall{Call: call}
}

// MockPrecheckBackendAgentVersionCall wrap *gomock.Call
type MockPrecheckBackendAgentVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAgentVersionCall) Return(arg0 version.Number, arg1 error) *MockPrecheckBackendAgentVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAgentVersionCall) Do(f func() (version.Number, error)) *MockPrecheckBackendAgentVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAgentVersionCall) DoAndReturn(f func() (version.Number, error)) *MockPrecheckBackendAgentVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllApplications mocks base method.
func (m *MockPrecheckBackend) AllApplications() ([]migration.PrecheckApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplications")
	ret0, _ := ret[0].([]migration.PrecheckApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplications indicates an expected call of AllApplications.
func (mr *MockPrecheckBackendMockRecorder) AllApplications() *MockPrecheckBackendAllApplicationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplications", reflect.TypeOf((*MockPrecheckBackend)(nil).AllApplications))
	return &MockPrecheckBackendAllApplicationsCall{Call: call}
}

// MockPrecheckBackendAllApplicationsCall wrap *gomock.Call
type MockPrecheckBackendAllApplicationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAllApplicationsCall) Return(arg0 []migration.PrecheckApplication, arg1 error) *MockPrecheckBackendAllApplicationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAllApplicationsCall) Do(f func() ([]migration.PrecheckApplication, error)) *MockPrecheckBackendAllApplicationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAllApplicationsCall) DoAndReturn(f func() ([]migration.PrecheckApplication, error)) *MockPrecheckBackendAllApplicationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllMachines mocks base method.
func (m *MockPrecheckBackend) AllMachines() ([]migration.PrecheckMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]migration.PrecheckMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines.
func (mr *MockPrecheckBackendMockRecorder) AllMachines() *MockPrecheckBackendAllMachinesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockPrecheckBackend)(nil).AllMachines))
	return &MockPrecheckBackendAllMachinesCall{Call: call}
}

// MockPrecheckBackendAllMachinesCall wrap *gomock.Call
type MockPrecheckBackendAllMachinesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAllMachinesCall) Return(arg0 []migration.PrecheckMachine, arg1 error) *MockPrecheckBackendAllMachinesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAllMachinesCall) Do(f func() ([]migration.PrecheckMachine, error)) *MockPrecheckBackendAllMachinesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAllMachinesCall) DoAndReturn(f func() ([]migration.PrecheckMachine, error)) *MockPrecheckBackendAllMachinesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllMachinesCount mocks base method.
func (m *MockPrecheckBackend) AllMachinesCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachinesCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachinesCount indicates an expected call of AllMachinesCount.
func (mr *MockPrecheckBackendMockRecorder) AllMachinesCount() *MockPrecheckBackendAllMachinesCountCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachinesCount", reflect.TypeOf((*MockPrecheckBackend)(nil).AllMachinesCount))
	return &MockPrecheckBackendAllMachinesCountCall{Call: call}
}

// MockPrecheckBackendAllMachinesCountCall wrap *gomock.Call
type MockPrecheckBackendAllMachinesCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAllMachinesCountCall) Return(arg0 int, arg1 error) *MockPrecheckBackendAllMachinesCountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAllMachinesCountCall) Do(f func() (int, error)) *MockPrecheckBackendAllMachinesCountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAllMachinesCountCall) DoAndReturn(f func() (int, error)) *MockPrecheckBackendAllMachinesCountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllModelUUIDs mocks base method.
func (m *MockPrecheckBackend) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockPrecheckBackendMockRecorder) AllModelUUIDs() *MockPrecheckBackendAllModelUUIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockPrecheckBackend)(nil).AllModelUUIDs))
	return &MockPrecheckBackendAllModelUUIDsCall{Call: call}
}

// MockPrecheckBackendAllModelUUIDsCall wrap *gomock.Call
type MockPrecheckBackendAllModelUUIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAllModelUUIDsCall) Return(arg0 []string, arg1 error) *MockPrecheckBackendAllModelUUIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAllModelUUIDsCall) Do(f func() ([]string, error)) *MockPrecheckBackendAllModelUUIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAllModelUUIDsCall) DoAndReturn(f func() ([]string, error)) *MockPrecheckBackendAllModelUUIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllRelations mocks base method.
func (m *MockPrecheckBackend) AllRelations() ([]migration.PrecheckRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRelations")
	ret0, _ := ret[0].([]migration.PrecheckRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRelations indicates an expected call of AllRelations.
func (mr *MockPrecheckBackendMockRecorder) AllRelations() *MockPrecheckBackendAllRelationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRelations", reflect.TypeOf((*MockPrecheckBackend)(nil).AllRelations))
	return &MockPrecheckBackendAllRelationsCall{Call: call}
}

// MockPrecheckBackendAllRelationsCall wrap *gomock.Call
type MockPrecheckBackendAllRelationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendAllRelationsCall) Return(arg0 []migration.PrecheckRelation, arg1 error) *MockPrecheckBackendAllRelationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendAllRelationsCall) Do(f func() ([]migration.PrecheckRelation, error)) *MockPrecheckBackendAllRelationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendAllRelationsCall) DoAndReturn(f func() ([]migration.PrecheckRelation, error)) *MockPrecheckBackendAllRelationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerBackend mocks base method.
func (m *MockPrecheckBackend) ControllerBackend() (migration.PrecheckBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerBackend")
	ret0, _ := ret[0].(migration.PrecheckBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerBackend indicates an expected call of ControllerBackend.
func (mr *MockPrecheckBackendMockRecorder) ControllerBackend() *MockPrecheckBackendControllerBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerBackend", reflect.TypeOf((*MockPrecheckBackend)(nil).ControllerBackend))
	return &MockPrecheckBackendControllerBackendCall{Call: call}
}

// MockPrecheckBackendControllerBackendCall wrap *gomock.Call
type MockPrecheckBackendControllerBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendControllerBackendCall) Return(arg0 migration.PrecheckBackend, arg1 error) *MockPrecheckBackendControllerBackendCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendControllerBackendCall) Do(f func() (migration.PrecheckBackend, error)) *MockPrecheckBackendControllerBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendControllerBackendCall) DoAndReturn(f func() (migration.PrecheckBackend, error)) *MockPrecheckBackendControllerBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsMigrationActive mocks base method.
func (m *MockPrecheckBackend) IsMigrationActive(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMigrationActive", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMigrationActive indicates an expected call of IsMigrationActive.
func (mr *MockPrecheckBackendMockRecorder) IsMigrationActive(arg0 any) *MockPrecheckBackendIsMigrationActiveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMigrationActive", reflect.TypeOf((*MockPrecheckBackend)(nil).IsMigrationActive), arg0)
	return &MockPrecheckBackendIsMigrationActiveCall{Call: call}
}

// MockPrecheckBackendIsMigrationActiveCall wrap *gomock.Call
type MockPrecheckBackendIsMigrationActiveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendIsMigrationActiveCall) Return(arg0 bool, arg1 error) *MockPrecheckBackendIsMigrationActiveCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendIsMigrationActiveCall) Do(f func(string) (bool, error)) *MockPrecheckBackendIsMigrationActiveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendIsMigrationActiveCall) DoAndReturn(f func(string) (bool, error)) *MockPrecheckBackendIsMigrationActiveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineCountForBase mocks base method.
func (m *MockPrecheckBackend) MachineCountForBase(arg0 ...state.Base) (map[string]int, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MachineCountForBase", varargs...)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineCountForBase indicates an expected call of MachineCountForBase.
func (mr *MockPrecheckBackendMockRecorder) MachineCountForBase(arg0 ...any) *MockPrecheckBackendMachineCountForBaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineCountForBase", reflect.TypeOf((*MockPrecheckBackend)(nil).MachineCountForBase), arg0...)
	return &MockPrecheckBackendMachineCountForBaseCall{Call: call}
}

// MockPrecheckBackendMachineCountForBaseCall wrap *gomock.Call
type MockPrecheckBackendMachineCountForBaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendMachineCountForBaseCall) Return(arg0 map[string]int, arg1 error) *MockPrecheckBackendMachineCountForBaseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendMachineCountForBaseCall) Do(f func(...state.Base) (map[string]int, error)) *MockPrecheckBackendMachineCountForBaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendMachineCountForBaseCall) DoAndReturn(f func(...state.Base) (map[string]int, error)) *MockPrecheckBackendMachineCountForBaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Model mocks base method.
func (m *MockPrecheckBackend) Model() (migration.PrecheckModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(migration.PrecheckModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockPrecheckBackendMockRecorder) Model() *MockPrecheckBackendModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockPrecheckBackend)(nil).Model))
	return &MockPrecheckBackendModelCall{Call: call}
}

// MockPrecheckBackendModelCall wrap *gomock.Call
type MockPrecheckBackendModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendModelCall) Return(arg0 migration.PrecheckModel, arg1 error) *MockPrecheckBackendModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendModelCall) Do(f func() (migration.PrecheckModel, error)) *MockPrecheckBackendModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendModelCall) DoAndReturn(f func() (migration.PrecheckModel, error)) *MockPrecheckBackendModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MongoCurrentStatus mocks base method.
func (m *MockPrecheckBackend) MongoCurrentStatus() (*replicaset.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoCurrentStatus")
	ret0, _ := ret[0].(*replicaset.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoCurrentStatus indicates an expected call of MongoCurrentStatus.
func (mr *MockPrecheckBackendMockRecorder) MongoCurrentStatus() *MockPrecheckBackendMongoCurrentStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoCurrentStatus", reflect.TypeOf((*MockPrecheckBackend)(nil).MongoCurrentStatus))
	return &MockPrecheckBackendMongoCurrentStatusCall{Call: call}
}

// MockPrecheckBackendMongoCurrentStatusCall wrap *gomock.Call
type MockPrecheckBackendMongoCurrentStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendMongoCurrentStatusCall) Return(arg0 *replicaset.Status, arg1 error) *MockPrecheckBackendMongoCurrentStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendMongoCurrentStatusCall) Do(f func() (*replicaset.Status, error)) *MockPrecheckBackendMongoCurrentStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendMongoCurrentStatusCall) DoAndReturn(f func() (*replicaset.Status, error)) *MockPrecheckBackendMongoCurrentStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NeedsCleanup mocks base method.
func (m *MockPrecheckBackend) NeedsCleanup() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsCleanup")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeedsCleanup indicates an expected call of NeedsCleanup.
func (mr *MockPrecheckBackendMockRecorder) NeedsCleanup() *MockPrecheckBackendNeedsCleanupCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsCleanup", reflect.TypeOf((*MockPrecheckBackend)(nil).NeedsCleanup))
	return &MockPrecheckBackendNeedsCleanupCall{Call: call}
}

// MockPrecheckBackendNeedsCleanupCall wrap *gomock.Call
type MockPrecheckBackendNeedsCleanupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPrecheckBackendNeedsCleanupCall) Return(arg0 bool, arg1 error) *MockPrecheckBackendNeedsCleanupCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPrecheckBackendNeedsCleanupCall) Do(f func() (bool, error)) *MockPrecheckBackendNeedsCleanupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPrecheckBackendNeedsCleanupCall) DoAndReturn(f func() (bool, error)) *MockPrecheckBackendNeedsCleanupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
