// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/leadership (interfaces: Reader)
//
// Generated by this command:
//
//	mockgen -typed -package sshclient_test -destination leadership_mock_test.go github.com/juju/juju/core/leadership Reader
//

// Package sshclient_test is a generated GoMock package.
package sshclient_test

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Leaders mocks base method.
func (m *MockReader) Leaders() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leaders")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leaders indicates an expected call of Leaders.
func (mr *MockReaderMockRecorder) Leaders() *MockReaderLeadersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leaders", reflect.TypeOf((*MockReader)(nil).Leaders))
	return &MockReaderLeadersCall{Call: call}
}

// MockReaderLeadersCall wrap *gomock.Call
type MockReaderLeadersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockReaderLeadersCall) Return(arg0 map[string]string, arg1 error) *MockReaderLeadersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockReaderLeadersCall) Do(f func() (map[string]string, error)) *MockReaderLeadersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockReaderLeadersCall) DoAndReturn(f func() (map[string]string, error)) *MockReaderLeadersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}