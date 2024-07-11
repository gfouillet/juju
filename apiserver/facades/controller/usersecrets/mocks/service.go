// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/usersecrets (interfaces: SecretService)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/service.go github.com/juju/juju/apiserver/facades/controller/usersecrets SecretService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretService is a mock of SecretService interface.
type MockSecretService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceMockRecorder
}

// MockSecretServiceMockRecorder is the mock recorder for MockSecretService.
type MockSecretServiceMockRecorder struct {
	mock *MockSecretService
}

// NewMockSecretService creates a new mock instance.
func NewMockSecretService(ctrl *gomock.Controller) *MockSecretService {
	mock := &MockSecretService{ctrl: ctrl}
	mock.recorder = &MockSecretServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretService) EXPECT() *MockSecretServiceMockRecorder {
	return m.recorder
}

// DeleteObsoleteUserSecretRevisions mocks base method.
func (m *MockSecretService) DeleteObsoleteUserSecretRevisions(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObsoleteUserSecretRevisions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObsoleteUserSecretRevisions indicates an expected call of DeleteObsoleteUserSecretRevisions.
func (mr *MockSecretServiceMockRecorder) DeleteObsoleteUserSecretRevisions(arg0 any) *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObsoleteUserSecretRevisions", reflect.TypeOf((*MockSecretService)(nil).DeleteObsoleteUserSecretRevisions), arg0)
	return &MockSecretServiceDeleteObsoleteUserSecretRevisionsCall{Call: call}
}

// MockSecretServiceDeleteObsoleteUserSecretRevisionsCall wrap *gomock.Call
type MockSecretServiceDeleteObsoleteUserSecretRevisionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall) Return(arg0 error) *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall) Do(f func(context.Context) error) *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall) DoAndReturn(f func(context.Context) error) *MockSecretServiceDeleteObsoleteUserSecretRevisionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecret mocks base method.
func (m *MockSecretService) GetSecret(arg0 context.Context, arg1 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretServiceMockRecorder) GetSecret(arg0, arg1 any) *MockSecretServiceGetSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretService)(nil).GetSecret), arg0, arg1)
	return &MockSecretServiceGetSecretCall{Call: call}
}

// MockSecretServiceGetSecretCall wrap *gomock.Call
type MockSecretServiceGetSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceGetSecretCall) Return(arg0 *secrets.SecretMetadata, arg1 error) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceGetSecretCall) Do(f func(context.Context, *secrets.URI) (*secrets.SecretMetadata, error)) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceGetSecretCall) DoAndReturn(f func(context.Context, *secrets.URI) (*secrets.SecretMetadata, error)) *MockSecretServiceGetSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchObsoleteUserSecretsToPrune mocks base method.
func (m *MockSecretService) WatchObsoleteUserSecretsToPrune(arg0 context.Context) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchObsoleteUserSecretsToPrune", arg0)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchObsoleteUserSecretsToPrune indicates an expected call of WatchObsoleteUserSecretsToPrune.
func (mr *MockSecretServiceMockRecorder) WatchObsoleteUserSecretsToPrune(arg0 any) *MockSecretServiceWatchObsoleteUserSecretsToPruneCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchObsoleteUserSecretsToPrune", reflect.TypeOf((*MockSecretService)(nil).WatchObsoleteUserSecretsToPrune), arg0)
	return &MockSecretServiceWatchObsoleteUserSecretsToPruneCall{Call: call}
}

// MockSecretServiceWatchObsoleteUserSecretsToPruneCall wrap *gomock.Call
type MockSecretServiceWatchObsoleteUserSecretsToPruneCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSecretServiceWatchObsoleteUserSecretsToPruneCall) Return(arg0 watcher.Watcher[struct{}], arg1 error) *MockSecretServiceWatchObsoleteUserSecretsToPruneCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSecretServiceWatchObsoleteUserSecretsToPruneCall) Do(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockSecretServiceWatchObsoleteUserSecretsToPruneCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSecretServiceWatchObsoleteUserSecretsToPruneCall) DoAndReturn(f func(context.Context) (watcher.Watcher[struct{}], error)) *MockSecretServiceWatchObsoleteUserSecretsToPruneCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}