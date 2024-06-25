// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/caasmodelconfigmanager (interfaces: CAASBroker)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/broker_mock.go github.com/juju/juju/internal/worker/caasmodelconfigmanager CAASBroker
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	docker "github.com/juju/juju/internal/docker"
	gomock "go.uber.org/mock/gomock"
)

// MockCAASBroker is a mock of CAASBroker interface.
type MockCAASBroker struct {
	ctrl     *gomock.Controller
	recorder *MockCAASBrokerMockRecorder
}

// MockCAASBrokerMockRecorder is the mock recorder for MockCAASBroker.
type MockCAASBrokerMockRecorder struct {
	mock *MockCAASBroker
}

// NewMockCAASBroker creates a new mock instance.
func NewMockCAASBroker(ctrl *gomock.Controller) *MockCAASBroker {
	mock := &MockCAASBroker{ctrl: ctrl}
	mock.recorder = &MockCAASBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASBroker) EXPECT() *MockCAASBrokerMockRecorder {
	return m.recorder
}

// EnsureImageRepoSecret mocks base method.
func (m *MockCAASBroker) EnsureImageRepoSecret(arg0 context.Context, arg1 docker.ImageRepoDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureImageRepoSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureImageRepoSecret indicates an expected call of EnsureImageRepoSecret.
func (mr *MockCAASBrokerMockRecorder) EnsureImageRepoSecret(arg0, arg1 any) *MockCAASBrokerEnsureImageRepoSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureImageRepoSecret", reflect.TypeOf((*MockCAASBroker)(nil).EnsureImageRepoSecret), arg0, arg1)
	return &MockCAASBrokerEnsureImageRepoSecretCall{Call: call}
}

// MockCAASBrokerEnsureImageRepoSecretCall wrap *gomock.Call
type MockCAASBrokerEnsureImageRepoSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCAASBrokerEnsureImageRepoSecretCall) Return(arg0 error) *MockCAASBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCAASBrokerEnsureImageRepoSecretCall) Do(f func(context.Context, docker.ImageRepoDetails) error) *MockCAASBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCAASBrokerEnsureImageRepoSecretCall) DoAndReturn(f func(context.Context, docker.ImageRepoDetails) error) *MockCAASBrokerEnsureImageRepoSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}