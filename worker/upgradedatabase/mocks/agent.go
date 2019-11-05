// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/agent (interfaces: Agent,Config,ConfigSetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	agent "github.com/juju/juju/agent"
	api "github.com/juju/juju/api"
	params "github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/core/model"
	network "github.com/juju/juju/core/network"
	mongo "github.com/juju/juju/mongo"
	shell "github.com/juju/utils/shell"
	version "github.com/juju/version"
	names_v3 "gopkg.in/juju/names.v3"
)

// MockAgent is a mock of Agent interface
type MockAgent struct {
	ctrl     *gomock.Controller
	recorder *MockAgentMockRecorder
}

// MockAgentMockRecorder is the mock recorder for MockAgent
type MockAgentMockRecorder struct {
	mock *MockAgent
}

// NewMockAgent creates a new mock instance
func NewMockAgent(ctrl *gomock.Controller) *MockAgent {
	mock := &MockAgent{ctrl: ctrl}
	mock.recorder = &MockAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgent) EXPECT() *MockAgentMockRecorder {
	return m.recorder
}

// ChangeConfig mocks base method
func (m *MockAgent) ChangeConfig(arg0 agent.ConfigMutator) error {
	ret := m.ctrl.Call(m, "ChangeConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeConfig indicates an expected call of ChangeConfig
func (mr *MockAgentMockRecorder) ChangeConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeConfig", reflect.TypeOf((*MockAgent)(nil).ChangeConfig), arg0)
}

// CurrentConfig mocks base method
func (m *MockAgent) CurrentConfig() agent.Config {
	ret := m.ctrl.Call(m, "CurrentConfig")
	ret0, _ := ret[0].(agent.Config)
	return ret0
}

// CurrentConfig indicates an expected call of CurrentConfig
func (mr *MockAgentMockRecorder) CurrentConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentConfig", reflect.TypeOf((*MockAgent)(nil).CurrentConfig))
}

// MockConfig is a mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// APIAddresses mocks base method
func (m *MockConfig) APIAddresses() ([]string, error) {
	ret := m.ctrl.Call(m, "APIAddresses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIAddresses indicates an expected call of APIAddresses
func (mr *MockConfigMockRecorder) APIAddresses() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIAddresses", reflect.TypeOf((*MockConfig)(nil).APIAddresses))
}

// APIInfo mocks base method
func (m *MockConfig) APIInfo() (*api.Info, bool) {
	ret := m.ctrl.Call(m, "APIInfo")
	ret0, _ := ret[0].(*api.Info)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// APIInfo indicates an expected call of APIInfo
func (mr *MockConfigMockRecorder) APIInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIInfo", reflect.TypeOf((*MockConfig)(nil).APIInfo))
}

// CACert mocks base method
func (m *MockConfig) CACert() string {
	ret := m.ctrl.Call(m, "CACert")
	ret0, _ := ret[0].(string)
	return ret0
}

// CACert indicates an expected call of CACert
func (mr *MockConfigMockRecorder) CACert() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CACert", reflect.TypeOf((*MockConfig)(nil).CACert))
}

// Controller mocks base method
func (m *MockConfig) Controller() names_v3.ControllerTag {
	ret := m.ctrl.Call(m, "Controller")
	ret0, _ := ret[0].(names_v3.ControllerTag)
	return ret0
}

// Controller indicates an expected call of Controller
func (mr *MockConfigMockRecorder) Controller() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockConfig)(nil).Controller))
}

// DataDir mocks base method
func (m *MockConfig) DataDir() string {
	ret := m.ctrl.Call(m, "DataDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DataDir indicates an expected call of DataDir
func (mr *MockConfigMockRecorder) DataDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataDir", reflect.TypeOf((*MockConfig)(nil).DataDir))
}

// Dir mocks base method
func (m *MockConfig) Dir() string {
	ret := m.ctrl.Call(m, "Dir")
	ret0, _ := ret[0].(string)
	return ret0
}

// Dir indicates an expected call of Dir
func (mr *MockConfigMockRecorder) Dir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dir", reflect.TypeOf((*MockConfig)(nil).Dir))
}

// Jobs mocks base method
func (m *MockConfig) Jobs() []model.MachineJob {
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].([]model.MachineJob)
	return ret0
}

// Jobs indicates an expected call of Jobs
func (mr *MockConfigMockRecorder) Jobs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockConfig)(nil).Jobs))
}

// LogDir mocks base method
func (m *MockConfig) LogDir() string {
	ret := m.ctrl.Call(m, "LogDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogDir indicates an expected call of LogDir
func (mr *MockConfigMockRecorder) LogDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDir", reflect.TypeOf((*MockConfig)(nil).LogDir))
}

// LoggingConfig mocks base method
func (m *MockConfig) LoggingConfig() string {
	ret := m.ctrl.Call(m, "LoggingConfig")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoggingConfig indicates an expected call of LoggingConfig
func (mr *MockConfigMockRecorder) LoggingConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoggingConfig", reflect.TypeOf((*MockConfig)(nil).LoggingConfig))
}

// MetricsSpoolDir mocks base method
func (m *MockConfig) MetricsSpoolDir() string {
	ret := m.ctrl.Call(m, "MetricsSpoolDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// MetricsSpoolDir indicates an expected call of MetricsSpoolDir
func (mr *MockConfigMockRecorder) MetricsSpoolDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricsSpoolDir", reflect.TypeOf((*MockConfig)(nil).MetricsSpoolDir))
}

// Model mocks base method
func (m *MockConfig) Model() names_v3.ModelTag {
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(names_v3.ModelTag)
	return ret0
}

// Model indicates an expected call of Model
func (mr *MockConfigMockRecorder) Model() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockConfig)(nil).Model))
}

// MongoInfo mocks base method
func (m *MockConfig) MongoInfo() (*mongo.MongoInfo, bool) {
	ret := m.ctrl.Call(m, "MongoInfo")
	ret0, _ := ret[0].(*mongo.MongoInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// MongoInfo indicates an expected call of MongoInfo
func (mr *MockConfigMockRecorder) MongoInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoInfo", reflect.TypeOf((*MockConfig)(nil).MongoInfo))
}

// MongoMemoryProfile mocks base method
func (m *MockConfig) MongoMemoryProfile() mongo.MemoryProfile {
	ret := m.ctrl.Call(m, "MongoMemoryProfile")
	ret0, _ := ret[0].(mongo.MemoryProfile)
	return ret0
}

// MongoMemoryProfile indicates an expected call of MongoMemoryProfile
func (mr *MockConfigMockRecorder) MongoMemoryProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoMemoryProfile", reflect.TypeOf((*MockConfig)(nil).MongoMemoryProfile))
}

// MongoVersion mocks base method
func (m *MockConfig) MongoVersion() mongo.Version {
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(mongo.Version)
	return ret0
}

// MongoVersion indicates an expected call of MongoVersion
func (mr *MockConfigMockRecorder) MongoVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockConfig)(nil).MongoVersion))
}

// Nonce mocks base method
func (m *MockConfig) Nonce() string {
	ret := m.ctrl.Call(m, "Nonce")
	ret0, _ := ret[0].(string)
	return ret0
}

// Nonce indicates an expected call of Nonce
func (mr *MockConfigMockRecorder) Nonce() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockConfig)(nil).Nonce))
}

// OldPassword mocks base method
func (m *MockConfig) OldPassword() string {
	ret := m.ctrl.Call(m, "OldPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// OldPassword indicates an expected call of OldPassword
func (mr *MockConfigMockRecorder) OldPassword() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OldPassword", reflect.TypeOf((*MockConfig)(nil).OldPassword))
}

// StateServingInfo mocks base method
func (m *MockConfig) StateServingInfo() (params.StateServingInfo, bool) {
	ret := m.ctrl.Call(m, "StateServingInfo")
	ret0, _ := ret[0].(params.StateServingInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// StateServingInfo indicates an expected call of StateServingInfo
func (mr *MockConfigMockRecorder) StateServingInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateServingInfo", reflect.TypeOf((*MockConfig)(nil).StateServingInfo))
}

// SystemIdentityPath mocks base method
func (m *MockConfig) SystemIdentityPath() string {
	ret := m.ctrl.Call(m, "SystemIdentityPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// SystemIdentityPath indicates an expected call of SystemIdentityPath
func (mr *MockConfigMockRecorder) SystemIdentityPath() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemIdentityPath", reflect.TypeOf((*MockConfig)(nil).SystemIdentityPath))
}

// Tag mocks base method
func (m *MockConfig) Tag() names_v3.Tag {
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names_v3.Tag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockConfigMockRecorder) Tag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockConfig)(nil).Tag))
}

// UpgradedToVersion mocks base method
func (m *MockConfig) UpgradedToVersion() version.Number {
	ret := m.ctrl.Call(m, "UpgradedToVersion")
	ret0, _ := ret[0].(version.Number)
	return ret0
}

// UpgradedToVersion indicates an expected call of UpgradedToVersion
func (mr *MockConfigMockRecorder) UpgradedToVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradedToVersion", reflect.TypeOf((*MockConfig)(nil).UpgradedToVersion))
}

// Value mocks base method
func (m *MockConfig) Value(arg0 string) string {
	ret := m.ctrl.Call(m, "Value", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockConfigMockRecorder) Value(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockConfig)(nil).Value), arg0)
}

// WriteCommands mocks base method
func (m *MockConfig) WriteCommands(arg0 shell.Renderer) ([]string, error) {
	ret := m.ctrl.Call(m, "WriteCommands", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteCommands indicates an expected call of WriteCommands
func (mr *MockConfigMockRecorder) WriteCommands(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCommands", reflect.TypeOf((*MockConfig)(nil).WriteCommands), arg0)
}

// MockConfigSetter is a mock of ConfigSetter interface
type MockConfigSetter struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSetterMockRecorder
}

// MockConfigSetterMockRecorder is the mock recorder for MockConfigSetter
type MockConfigSetterMockRecorder struct {
	mock *MockConfigSetter
}

// NewMockConfigSetter creates a new mock instance
func NewMockConfigSetter(ctrl *gomock.Controller) *MockConfigSetter {
	mock := &MockConfigSetter{ctrl: ctrl}
	mock.recorder = &MockConfigSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigSetter) EXPECT() *MockConfigSetterMockRecorder {
	return m.recorder
}

// APIAddresses mocks base method
func (m *MockConfigSetter) APIAddresses() ([]string, error) {
	ret := m.ctrl.Call(m, "APIAddresses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIAddresses indicates an expected call of APIAddresses
func (mr *MockConfigSetterMockRecorder) APIAddresses() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIAddresses", reflect.TypeOf((*MockConfigSetter)(nil).APIAddresses))
}

// APIInfo mocks base method
func (m *MockConfigSetter) APIInfo() (*api.Info, bool) {
	ret := m.ctrl.Call(m, "APIInfo")
	ret0, _ := ret[0].(*api.Info)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// APIInfo indicates an expected call of APIInfo
func (mr *MockConfigSetterMockRecorder) APIInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIInfo", reflect.TypeOf((*MockConfigSetter)(nil).APIInfo))
}

// CACert mocks base method
func (m *MockConfigSetter) CACert() string {
	ret := m.ctrl.Call(m, "CACert")
	ret0, _ := ret[0].(string)
	return ret0
}

// CACert indicates an expected call of CACert
func (mr *MockConfigSetterMockRecorder) CACert() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CACert", reflect.TypeOf((*MockConfigSetter)(nil).CACert))
}

// Clone mocks base method
func (m *MockConfigSetter) Clone() agent.Config {
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(agent.Config)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockConfigSetterMockRecorder) Clone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockConfigSetter)(nil).Clone))
}

// Controller mocks base method
func (m *MockConfigSetter) Controller() names_v3.ControllerTag {
	ret := m.ctrl.Call(m, "Controller")
	ret0, _ := ret[0].(names_v3.ControllerTag)
	return ret0
}

// Controller indicates an expected call of Controller
func (mr *MockConfigSetterMockRecorder) Controller() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockConfigSetter)(nil).Controller))
}

// DataDir mocks base method
func (m *MockConfigSetter) DataDir() string {
	ret := m.ctrl.Call(m, "DataDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DataDir indicates an expected call of DataDir
func (mr *MockConfigSetterMockRecorder) DataDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataDir", reflect.TypeOf((*MockConfigSetter)(nil).DataDir))
}

// Dir mocks base method
func (m *MockConfigSetter) Dir() string {
	ret := m.ctrl.Call(m, "Dir")
	ret0, _ := ret[0].(string)
	return ret0
}

// Dir indicates an expected call of Dir
func (mr *MockConfigSetterMockRecorder) Dir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dir", reflect.TypeOf((*MockConfigSetter)(nil).Dir))
}

// Jobs mocks base method
func (m *MockConfigSetter) Jobs() []model.MachineJob {
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].([]model.MachineJob)
	return ret0
}

// Jobs indicates an expected call of Jobs
func (mr *MockConfigSetterMockRecorder) Jobs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockConfigSetter)(nil).Jobs))
}

// LogDir mocks base method
func (m *MockConfigSetter) LogDir() string {
	ret := m.ctrl.Call(m, "LogDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogDir indicates an expected call of LogDir
func (mr *MockConfigSetterMockRecorder) LogDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDir", reflect.TypeOf((*MockConfigSetter)(nil).LogDir))
}

// LoggingConfig mocks base method
func (m *MockConfigSetter) LoggingConfig() string {
	ret := m.ctrl.Call(m, "LoggingConfig")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoggingConfig indicates an expected call of LoggingConfig
func (mr *MockConfigSetterMockRecorder) LoggingConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoggingConfig", reflect.TypeOf((*MockConfigSetter)(nil).LoggingConfig))
}

// MetricsSpoolDir mocks base method
func (m *MockConfigSetter) MetricsSpoolDir() string {
	ret := m.ctrl.Call(m, "MetricsSpoolDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// MetricsSpoolDir indicates an expected call of MetricsSpoolDir
func (mr *MockConfigSetterMockRecorder) MetricsSpoolDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricsSpoolDir", reflect.TypeOf((*MockConfigSetter)(nil).MetricsSpoolDir))
}

// Model mocks base method
func (m *MockConfigSetter) Model() names_v3.ModelTag {
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(names_v3.ModelTag)
	return ret0
}

// Model indicates an expected call of Model
func (mr *MockConfigSetterMockRecorder) Model() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockConfigSetter)(nil).Model))
}

// MongoInfo mocks base method
func (m *MockConfigSetter) MongoInfo() (*mongo.MongoInfo, bool) {
	ret := m.ctrl.Call(m, "MongoInfo")
	ret0, _ := ret[0].(*mongo.MongoInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// MongoInfo indicates an expected call of MongoInfo
func (mr *MockConfigSetterMockRecorder) MongoInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoInfo", reflect.TypeOf((*MockConfigSetter)(nil).MongoInfo))
}

// MongoMemoryProfile mocks base method
func (m *MockConfigSetter) MongoMemoryProfile() mongo.MemoryProfile {
	ret := m.ctrl.Call(m, "MongoMemoryProfile")
	ret0, _ := ret[0].(mongo.MemoryProfile)
	return ret0
}

// MongoMemoryProfile indicates an expected call of MongoMemoryProfile
func (mr *MockConfigSetterMockRecorder) MongoMemoryProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoMemoryProfile", reflect.TypeOf((*MockConfigSetter)(nil).MongoMemoryProfile))
}

// MongoVersion mocks base method
func (m *MockConfigSetter) MongoVersion() mongo.Version {
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(mongo.Version)
	return ret0
}

// MongoVersion indicates an expected call of MongoVersion
func (mr *MockConfigSetterMockRecorder) MongoVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockConfigSetter)(nil).MongoVersion))
}

// Nonce mocks base method
func (m *MockConfigSetter) Nonce() string {
	ret := m.ctrl.Call(m, "Nonce")
	ret0, _ := ret[0].(string)
	return ret0
}

// Nonce indicates an expected call of Nonce
func (mr *MockConfigSetterMockRecorder) Nonce() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockConfigSetter)(nil).Nonce))
}

// OldPassword mocks base method
func (m *MockConfigSetter) OldPassword() string {
	ret := m.ctrl.Call(m, "OldPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// OldPassword indicates an expected call of OldPassword
func (mr *MockConfigSetterMockRecorder) OldPassword() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OldPassword", reflect.TypeOf((*MockConfigSetter)(nil).OldPassword))
}

// SetAPIHostPorts mocks base method
func (m *MockConfigSetter) SetAPIHostPorts(arg0 []network.HostPorts) {
	m.ctrl.Call(m, "SetAPIHostPorts", arg0)
}

// SetAPIHostPorts indicates an expected call of SetAPIHostPorts
func (mr *MockConfigSetterMockRecorder) SetAPIHostPorts(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAPIHostPorts", reflect.TypeOf((*MockConfigSetter)(nil).SetAPIHostPorts), arg0)
}

// SetCACert mocks base method
func (m *MockConfigSetter) SetCACert(arg0 string) {
	m.ctrl.Call(m, "SetCACert", arg0)
}

// SetCACert indicates an expected call of SetCACert
func (mr *MockConfigSetterMockRecorder) SetCACert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCACert", reflect.TypeOf((*MockConfigSetter)(nil).SetCACert), arg0)
}

// SetControllerAPIPort mocks base method
func (m *MockConfigSetter) SetControllerAPIPort(arg0 int) {
	m.ctrl.Call(m, "SetControllerAPIPort", arg0)
}

// SetControllerAPIPort indicates an expected call of SetControllerAPIPort
func (mr *MockConfigSetterMockRecorder) SetControllerAPIPort(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerAPIPort", reflect.TypeOf((*MockConfigSetter)(nil).SetControllerAPIPort), arg0)
}

// SetLoggingConfig mocks base method
func (m *MockConfigSetter) SetLoggingConfig(arg0 string) {
	m.ctrl.Call(m, "SetLoggingConfig", arg0)
}

// SetLoggingConfig indicates an expected call of SetLoggingConfig
func (mr *MockConfigSetterMockRecorder) SetLoggingConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoggingConfig", reflect.TypeOf((*MockConfigSetter)(nil).SetLoggingConfig), arg0)
}

// SetMongoMemoryProfile mocks base method
func (m *MockConfigSetter) SetMongoMemoryProfile(arg0 mongo.MemoryProfile) {
	m.ctrl.Call(m, "SetMongoMemoryProfile", arg0)
}

// SetMongoMemoryProfile indicates an expected call of SetMongoMemoryProfile
func (mr *MockConfigSetterMockRecorder) SetMongoMemoryProfile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMongoMemoryProfile", reflect.TypeOf((*MockConfigSetter)(nil).SetMongoMemoryProfile), arg0)
}

// SetMongoVersion mocks base method
func (m *MockConfigSetter) SetMongoVersion(arg0 mongo.Version) {
	m.ctrl.Call(m, "SetMongoVersion", arg0)
}

// SetMongoVersion indicates an expected call of SetMongoVersion
func (mr *MockConfigSetterMockRecorder) SetMongoVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMongoVersion", reflect.TypeOf((*MockConfigSetter)(nil).SetMongoVersion), arg0)
}

// SetOldPassword mocks base method
func (m *MockConfigSetter) SetOldPassword(arg0 string) {
	m.ctrl.Call(m, "SetOldPassword", arg0)
}

// SetOldPassword indicates an expected call of SetOldPassword
func (mr *MockConfigSetterMockRecorder) SetOldPassword(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOldPassword", reflect.TypeOf((*MockConfigSetter)(nil).SetOldPassword), arg0)
}

// SetPassword mocks base method
func (m *MockConfigSetter) SetPassword(arg0 string) {
	m.ctrl.Call(m, "SetPassword", arg0)
}

// SetPassword indicates an expected call of SetPassword
func (mr *MockConfigSetterMockRecorder) SetPassword(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockConfigSetter)(nil).SetPassword), arg0)
}

// SetStateServingInfo mocks base method
func (m *MockConfigSetter) SetStateServingInfo(arg0 params.StateServingInfo) {
	m.ctrl.Call(m, "SetStateServingInfo", arg0)
}

// SetStateServingInfo indicates an expected call of SetStateServingInfo
func (mr *MockConfigSetterMockRecorder) SetStateServingInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateServingInfo", reflect.TypeOf((*MockConfigSetter)(nil).SetStateServingInfo), arg0)
}

// SetUpgradedToVersion mocks base method
func (m *MockConfigSetter) SetUpgradedToVersion(arg0 version.Number) {
	m.ctrl.Call(m, "SetUpgradedToVersion", arg0)
}

// SetUpgradedToVersion indicates an expected call of SetUpgradedToVersion
func (mr *MockConfigSetterMockRecorder) SetUpgradedToVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpgradedToVersion", reflect.TypeOf((*MockConfigSetter)(nil).SetUpgradedToVersion), arg0)
}

// SetValue mocks base method
func (m *MockConfigSetter) SetValue(arg0, arg1 string) {
	m.ctrl.Call(m, "SetValue", arg0, arg1)
}

// SetValue indicates an expected call of SetValue
func (mr *MockConfigSetterMockRecorder) SetValue(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValue", reflect.TypeOf((*MockConfigSetter)(nil).SetValue), arg0, arg1)
}

// StateServingInfo mocks base method
func (m *MockConfigSetter) StateServingInfo() (params.StateServingInfo, bool) {
	ret := m.ctrl.Call(m, "StateServingInfo")
	ret0, _ := ret[0].(params.StateServingInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// StateServingInfo indicates an expected call of StateServingInfo
func (mr *MockConfigSetterMockRecorder) StateServingInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateServingInfo", reflect.TypeOf((*MockConfigSetter)(nil).StateServingInfo))
}

// SystemIdentityPath mocks base method
func (m *MockConfigSetter) SystemIdentityPath() string {
	ret := m.ctrl.Call(m, "SystemIdentityPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// SystemIdentityPath indicates an expected call of SystemIdentityPath
func (mr *MockConfigSetterMockRecorder) SystemIdentityPath() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemIdentityPath", reflect.TypeOf((*MockConfigSetter)(nil).SystemIdentityPath))
}

// Tag mocks base method
func (m *MockConfigSetter) Tag() names_v3.Tag {
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names_v3.Tag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockConfigSetterMockRecorder) Tag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockConfigSetter)(nil).Tag))
}

// UpgradedToVersion mocks base method
func (m *MockConfigSetter) UpgradedToVersion() version.Number {
	ret := m.ctrl.Call(m, "UpgradedToVersion")
	ret0, _ := ret[0].(version.Number)
	return ret0
}

// UpgradedToVersion indicates an expected call of UpgradedToVersion
func (mr *MockConfigSetterMockRecorder) UpgradedToVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradedToVersion", reflect.TypeOf((*MockConfigSetter)(nil).UpgradedToVersion))
}

// Value mocks base method
func (m *MockConfigSetter) Value(arg0 string) string {
	ret := m.ctrl.Call(m, "Value", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockConfigSetterMockRecorder) Value(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockConfigSetter)(nil).Value), arg0)
}

// WriteCommands mocks base method
func (m *MockConfigSetter) WriteCommands(arg0 shell.Renderer) ([]string, error) {
	ret := m.ctrl.Call(m, "WriteCommands", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteCommands indicates an expected call of WriteCommands
func (mr *MockConfigSetterMockRecorder) WriteCommands(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCommands", reflect.TypeOf((*MockConfigSetter)(nil).WriteCommands), arg0)
}
