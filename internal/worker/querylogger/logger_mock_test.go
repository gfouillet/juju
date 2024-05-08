// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/logger (interfaces: Logger)
//
// Generated by this command:
//
//	mockgen -typed -package querylogger -destination logger_mock_test.go github.com/juju/juju/core/logger Logger
//

// Package querylogger is a generated GoMock package.
package querylogger

import (
	reflect "reflect"

	logger "github.com/juju/juju/core/logger"
	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Child mocks base method.
func (m *MockLogger) Child(arg0 string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Child", arg0)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Child indicates an expected call of Child.
func (mr *MockLoggerMockRecorder) Child(arg0 any) *MockLoggerChildCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Child", reflect.TypeOf((*MockLogger)(nil).Child), arg0)
	return &MockLoggerChildCall{Call: call}
}

// MockLoggerChildCall wrap *gomock.Call
type MockLoggerChildCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerChildCall) Return(arg0 logger.Logger) *MockLoggerChildCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerChildCall) Do(f func(string) logger.Logger) *MockLoggerChildCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerChildCall) DoAndReturn(f func(string) logger.Logger) *MockLoggerChildCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ChildWithTags mocks base method.
func (m *MockLogger) ChildWithTags(arg0 string, arg1 ...string) logger.Logger {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChildWithTags", varargs...)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// ChildWithTags indicates an expected call of ChildWithTags.
func (mr *MockLoggerMockRecorder) ChildWithTags(arg0 any, arg1 ...any) *MockLoggerChildWithTagsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildWithTags", reflect.TypeOf((*MockLogger)(nil).ChildWithTags), varargs...)
	return &MockLoggerChildWithTagsCall{Call: call}
}

// MockLoggerChildWithTagsCall wrap *gomock.Call
type MockLoggerChildWithTagsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerChildWithTagsCall) Return(arg0 logger.Logger) *MockLoggerChildWithTagsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerChildWithTagsCall) Do(f func(string, ...string) logger.Logger) *MockLoggerChildWithTagsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerChildWithTagsCall) DoAndReturn(f func(string, ...string) logger.Logger) *MockLoggerChildWithTagsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Criticalf mocks base method.
func (m *MockLogger) Criticalf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Criticalf", varargs...)
}

// Criticalf indicates an expected call of Criticalf.
func (mr *MockLoggerMockRecorder) Criticalf(arg0 any, arg1 ...any) *MockLoggerCriticalfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Criticalf", reflect.TypeOf((*MockLogger)(nil).Criticalf), varargs...)
	return &MockLoggerCriticalfCall{Call: call}
}

// MockLoggerCriticalfCall wrap *gomock.Call
type MockLoggerCriticalfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerCriticalfCall) Return() *MockLoggerCriticalfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerCriticalfCall) Do(f func(string, ...any)) *MockLoggerCriticalfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerCriticalfCall) DoAndReturn(f func(string, ...any)) *MockLoggerCriticalfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 any, arg1 ...any) *MockLoggerDebugfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
	return &MockLoggerDebugfCall{Call: call}
}

// MockLoggerDebugfCall wrap *gomock.Call
type MockLoggerDebugfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerDebugfCall) Return() *MockLoggerDebugfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerDebugfCall) Do(f func(string, ...any)) *MockLoggerDebugfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerDebugfCall) DoAndReturn(f func(string, ...any)) *MockLoggerDebugfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *MockLoggerErrorfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
	return &MockLoggerErrorfCall{Call: call}
}

// MockLoggerErrorfCall wrap *gomock.Call
type MockLoggerErrorfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerErrorfCall) Return() *MockLoggerErrorfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerErrorfCall) Do(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerErrorfCall) DoAndReturn(f func(string, ...any)) *MockLoggerErrorfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetChildByName mocks base method.
func (m *MockLogger) GetChildByName(arg0 string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildByName", arg0)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// GetChildByName indicates an expected call of GetChildByName.
func (mr *MockLoggerMockRecorder) GetChildByName(arg0 any) *MockLoggerGetChildByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildByName", reflect.TypeOf((*MockLogger)(nil).GetChildByName), arg0)
	return &MockLoggerGetChildByNameCall{Call: call}
}

// MockLoggerGetChildByNameCall wrap *gomock.Call
type MockLoggerGetChildByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerGetChildByNameCall) Return(arg0 logger.Logger) *MockLoggerGetChildByNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerGetChildByNameCall) Do(f func(string) logger.Logger) *MockLoggerGetChildByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerGetChildByNameCall) DoAndReturn(f func(string) logger.Logger) *MockLoggerGetChildByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 any, arg1 ...any) *MockLoggerInfofCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
	return &MockLoggerInfofCall{Call: call}
}

// MockLoggerInfofCall wrap *gomock.Call
type MockLoggerInfofCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerInfofCall) Return() *MockLoggerInfofCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerInfofCall) Do(f func(string, ...any)) *MockLoggerInfofCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerInfofCall) DoAndReturn(f func(string, ...any)) *MockLoggerInfofCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsDebugEnabled mocks base method.
func (m *MockLogger) IsDebugEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebugEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebugEnabled indicates an expected call of IsDebugEnabled.
func (mr *MockLoggerMockRecorder) IsDebugEnabled() *MockLoggerIsDebugEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebugEnabled", reflect.TypeOf((*MockLogger)(nil).IsDebugEnabled))
	return &MockLoggerIsDebugEnabledCall{Call: call}
}

// MockLoggerIsDebugEnabledCall wrap *gomock.Call
type MockLoggerIsDebugEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsDebugEnabledCall) Return(arg0 bool) *MockLoggerIsDebugEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsDebugEnabledCall) Do(f func() bool) *MockLoggerIsDebugEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsDebugEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsDebugEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsErrorEnabled mocks base method.
func (m *MockLogger) IsErrorEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsErrorEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsErrorEnabled indicates an expected call of IsErrorEnabled.
func (mr *MockLoggerMockRecorder) IsErrorEnabled() *MockLoggerIsErrorEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsErrorEnabled", reflect.TypeOf((*MockLogger)(nil).IsErrorEnabled))
	return &MockLoggerIsErrorEnabledCall{Call: call}
}

// MockLoggerIsErrorEnabledCall wrap *gomock.Call
type MockLoggerIsErrorEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsErrorEnabledCall) Return(arg0 bool) *MockLoggerIsErrorEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsErrorEnabledCall) Do(f func() bool) *MockLoggerIsErrorEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsErrorEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsErrorEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsInfoEnabled mocks base method.
func (m *MockLogger) IsInfoEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInfoEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInfoEnabled indicates an expected call of IsInfoEnabled.
func (mr *MockLoggerMockRecorder) IsInfoEnabled() *MockLoggerIsInfoEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInfoEnabled", reflect.TypeOf((*MockLogger)(nil).IsInfoEnabled))
	return &MockLoggerIsInfoEnabledCall{Call: call}
}

// MockLoggerIsInfoEnabledCall wrap *gomock.Call
type MockLoggerIsInfoEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsInfoEnabledCall) Return(arg0 bool) *MockLoggerIsInfoEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsInfoEnabledCall) Do(f func() bool) *MockLoggerIsInfoEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsInfoEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsInfoEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsLevelEnabled mocks base method.
func (m *MockLogger) IsLevelEnabled(arg0 logger.Level) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLevelEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLevelEnabled indicates an expected call of IsLevelEnabled.
func (mr *MockLoggerMockRecorder) IsLevelEnabled(arg0 any) *MockLoggerIsLevelEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLevelEnabled", reflect.TypeOf((*MockLogger)(nil).IsLevelEnabled), arg0)
	return &MockLoggerIsLevelEnabledCall{Call: call}
}

// MockLoggerIsLevelEnabledCall wrap *gomock.Call
type MockLoggerIsLevelEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsLevelEnabledCall) Return(arg0 bool) *MockLoggerIsLevelEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsLevelEnabledCall) Do(f func(logger.Level) bool) *MockLoggerIsLevelEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsLevelEnabledCall) DoAndReturn(f func(logger.Level) bool) *MockLoggerIsLevelEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *MockLoggerIsTraceEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
	return &MockLoggerIsTraceEnabledCall{Call: call}
}

// MockLoggerIsTraceEnabledCall wrap *gomock.Call
type MockLoggerIsTraceEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsTraceEnabledCall) Return(arg0 bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsTraceEnabledCall) Do(f func() bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsTraceEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsTraceEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsWarningEnabled mocks base method.
func (m *MockLogger) IsWarningEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWarningEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWarningEnabled indicates an expected call of IsWarningEnabled.
func (mr *MockLoggerMockRecorder) IsWarningEnabled() *MockLoggerIsWarningEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWarningEnabled", reflect.TypeOf((*MockLogger)(nil).IsWarningEnabled))
	return &MockLoggerIsWarningEnabledCall{Call: call}
}

// MockLoggerIsWarningEnabledCall wrap *gomock.Call
type MockLoggerIsWarningEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerIsWarningEnabledCall) Return(arg0 bool) *MockLoggerIsWarningEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerIsWarningEnabledCall) Do(f func() bool) *MockLoggerIsWarningEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerIsWarningEnabledCall) DoAndReturn(f func() bool) *MockLoggerIsWarningEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Logf mocks base method.
func (m *MockLogger) Logf(arg0 logger.Level, arg1 string, arg2 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(arg0, arg1 any, arg2 ...any) *MockLoggerLogfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
	return &MockLoggerLogfCall{Call: call}
}

// MockLoggerLogfCall wrap *gomock.Call
type MockLoggerLogfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerLogfCall) Return() *MockLoggerLogfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerLogfCall) Do(f func(logger.Level, string, ...any)) *MockLoggerLogfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerLogfCall) DoAndReturn(f func(logger.Level, string, ...any)) *MockLoggerLogfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *MockLoggerTracefCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
	return &MockLoggerTracefCall{Call: call}
}

// MockLoggerTracefCall wrap *gomock.Call
type MockLoggerTracefCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerTracefCall) Return() *MockLoggerTracefCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerTracefCall) Do(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerTracefCall) DoAndReturn(f func(string, ...any)) *MockLoggerTracefCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 any, arg1 ...any) *MockLoggerWarningfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
	return &MockLoggerWarningfCall{Call: call}
}

// MockLoggerWarningfCall wrap *gomock.Call
type MockLoggerWarningfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLoggerWarningfCall) Return() *MockLoggerWarningfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLoggerWarningfCall) Do(f func(string, ...any)) *MockLoggerWarningfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLoggerWarningfCall) DoAndReturn(f func(string, ...any)) *MockLoggerWarningfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
