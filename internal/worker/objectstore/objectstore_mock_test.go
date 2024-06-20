// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/objectstore (interfaces: TrackedObjectStore,MetadataServiceGetter,MetadataService,ModelClaimGetter,ControllerConfigService)
//
// Generated by this command:
//
//	mockgen -typed -package objectstore -destination objectstore_mock_test.go github.com/juju/juju/internal/worker/objectstore TrackedObjectStore,MetadataServiceGetter,MetadataService,ModelClaimGetter,ControllerConfigService
//

// Package objectstore is a generated GoMock package.
package objectstore

import (
	context "context"
	io "io"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	objectstore "github.com/juju/juju/core/objectstore"
	objectstore0 "github.com/juju/juju/internal/objectstore"
	gomock "go.uber.org/mock/gomock"
)

// MockTrackedObjectStore is a mock of TrackedObjectStore interface.
type MockTrackedObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockTrackedObjectStoreMockRecorder
}

// MockTrackedObjectStoreMockRecorder is the mock recorder for MockTrackedObjectStore.
type MockTrackedObjectStoreMockRecorder struct {
	mock *MockTrackedObjectStore
}

// NewMockTrackedObjectStore creates a new mock instance.
func NewMockTrackedObjectStore(ctrl *gomock.Controller) *MockTrackedObjectStore {
	mock := &MockTrackedObjectStore{ctrl: ctrl}
	mock.recorder = &MockTrackedObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackedObjectStore) EXPECT() *MockTrackedObjectStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTrackedObjectStore) Get(arg0 context.Context, arg1 string) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockTrackedObjectStoreMockRecorder) Get(arg0, arg1 any) *MockTrackedObjectStoreGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrackedObjectStore)(nil).Get), arg0, arg1)
	return &MockTrackedObjectStoreGetCall{Call: call}
}

// MockTrackedObjectStoreGetCall wrap *gomock.Call
type MockTrackedObjectStoreGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStoreGetCall) Return(arg0 io.ReadCloser, arg1 int64, arg2 error) *MockTrackedObjectStoreGetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStoreGetCall) Do(f func(context.Context, string) (io.ReadCloser, int64, error)) *MockTrackedObjectStoreGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStoreGetCall) DoAndReturn(f func(context.Context, string) (io.ReadCloser, int64, error)) *MockTrackedObjectStoreGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Kill mocks base method.
func (m *MockTrackedObjectStore) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockTrackedObjectStoreMockRecorder) Kill() *MockTrackedObjectStoreKillCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockTrackedObjectStore)(nil).Kill))
	return &MockTrackedObjectStoreKillCall{Call: call}
}

// MockTrackedObjectStoreKillCall wrap *gomock.Call
type MockTrackedObjectStoreKillCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStoreKillCall) Return() *MockTrackedObjectStoreKillCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStoreKillCall) Do(f func()) *MockTrackedObjectStoreKillCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStoreKillCall) DoAndReturn(f func()) *MockTrackedObjectStoreKillCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Put mocks base method.
func (m *MockTrackedObjectStore) Put(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockTrackedObjectStoreMockRecorder) Put(arg0, arg1, arg2, arg3 any) *MockTrackedObjectStorePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockTrackedObjectStore)(nil).Put), arg0, arg1, arg2, arg3)
	return &MockTrackedObjectStorePutCall{Call: call}
}

// MockTrackedObjectStorePutCall wrap *gomock.Call
type MockTrackedObjectStorePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStorePutCall) Return(arg0 error) *MockTrackedObjectStorePutCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStorePutCall) Do(f func(context.Context, string, io.Reader, int64) error) *MockTrackedObjectStorePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStorePutCall) DoAndReturn(f func(context.Context, string, io.Reader, int64) error) *MockTrackedObjectStorePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PutAndCheckHash mocks base method.
func (m *MockTrackedObjectStore) PutAndCheckHash(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 int64, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutAndCheckHash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutAndCheckHash indicates an expected call of PutAndCheckHash.
func (mr *MockTrackedObjectStoreMockRecorder) PutAndCheckHash(arg0, arg1, arg2, arg3, arg4 any) *MockTrackedObjectStorePutAndCheckHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutAndCheckHash", reflect.TypeOf((*MockTrackedObjectStore)(nil).PutAndCheckHash), arg0, arg1, arg2, arg3, arg4)
	return &MockTrackedObjectStorePutAndCheckHashCall{Call: call}
}

// MockTrackedObjectStorePutAndCheckHashCall wrap *gomock.Call
type MockTrackedObjectStorePutAndCheckHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStorePutAndCheckHashCall) Return(arg0 error) *MockTrackedObjectStorePutAndCheckHashCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStorePutAndCheckHashCall) Do(f func(context.Context, string, io.Reader, int64, string) error) *MockTrackedObjectStorePutAndCheckHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStorePutAndCheckHashCall) DoAndReturn(f func(context.Context, string, io.Reader, int64, string) error) *MockTrackedObjectStorePutAndCheckHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Remove mocks base method.
func (m *MockTrackedObjectStore) Remove(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTrackedObjectStoreMockRecorder) Remove(arg0, arg1 any) *MockTrackedObjectStoreRemoveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTrackedObjectStore)(nil).Remove), arg0, arg1)
	return &MockTrackedObjectStoreRemoveCall{Call: call}
}

// MockTrackedObjectStoreRemoveCall wrap *gomock.Call
type MockTrackedObjectStoreRemoveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStoreRemoveCall) Return(arg0 error) *MockTrackedObjectStoreRemoveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStoreRemoveCall) Do(f func(context.Context, string) error) *MockTrackedObjectStoreRemoveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStoreRemoveCall) DoAndReturn(f func(context.Context, string) error) *MockTrackedObjectStoreRemoveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *MockTrackedObjectStore) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockTrackedObjectStoreMockRecorder) Wait() *MockTrackedObjectStoreWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockTrackedObjectStore)(nil).Wait))
	return &MockTrackedObjectStoreWaitCall{Call: call}
}

// MockTrackedObjectStoreWaitCall wrap *gomock.Call
type MockTrackedObjectStoreWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTrackedObjectStoreWaitCall) Return(arg0 error) *MockTrackedObjectStoreWaitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTrackedObjectStoreWaitCall) Do(f func() error) *MockTrackedObjectStoreWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTrackedObjectStoreWaitCall) DoAndReturn(f func() error) *MockTrackedObjectStoreWaitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMetadataServiceGetter is a mock of MetadataServiceGetter interface.
type MockMetadataServiceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataServiceGetterMockRecorder
}

// MockMetadataServiceGetterMockRecorder is the mock recorder for MockMetadataServiceGetter.
type MockMetadataServiceGetterMockRecorder struct {
	mock *MockMetadataServiceGetter
}

// NewMockMetadataServiceGetter creates a new mock instance.
func NewMockMetadataServiceGetter(ctrl *gomock.Controller) *MockMetadataServiceGetter {
	mock := &MockMetadataServiceGetter{ctrl: ctrl}
	mock.recorder = &MockMetadataServiceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataServiceGetter) EXPECT() *MockMetadataServiceGetterMockRecorder {
	return m.recorder
}

// ForModelUUID mocks base method.
func (m *MockMetadataServiceGetter) ForModelUUID(arg0 string) MetadataService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForModelUUID", arg0)
	ret0, _ := ret[0].(MetadataService)
	return ret0
}

// ForModelUUID indicates an expected call of ForModelUUID.
func (mr *MockMetadataServiceGetterMockRecorder) ForModelUUID(arg0 any) *MockMetadataServiceGetterForModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForModelUUID", reflect.TypeOf((*MockMetadataServiceGetter)(nil).ForModelUUID), arg0)
	return &MockMetadataServiceGetterForModelUUIDCall{Call: call}
}

// MockMetadataServiceGetterForModelUUIDCall wrap *gomock.Call
type MockMetadataServiceGetterForModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMetadataServiceGetterForModelUUIDCall) Return(arg0 MetadataService) *MockMetadataServiceGetterForModelUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMetadataServiceGetterForModelUUIDCall) Do(f func(string) MetadataService) *MockMetadataServiceGetterForModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMetadataServiceGetterForModelUUIDCall) DoAndReturn(f func(string) MetadataService) *MockMetadataServiceGetterForModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMetadataService is a mock of MetadataService interface.
type MockMetadataService struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataServiceMockRecorder
}

// MockMetadataServiceMockRecorder is the mock recorder for MockMetadataService.
type MockMetadataServiceMockRecorder struct {
	mock *MockMetadataService
}

// NewMockMetadataService creates a new mock instance.
func NewMockMetadataService(ctrl *gomock.Controller) *MockMetadataService {
	mock := &MockMetadataService{ctrl: ctrl}
	mock.recorder = &MockMetadataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataService) EXPECT() *MockMetadataServiceMockRecorder {
	return m.recorder
}

// ObjectStore mocks base method.
func (m *MockMetadataService) ObjectStore() objectstore.ObjectStoreMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStore")
	ret0, _ := ret[0].(objectstore.ObjectStoreMetadata)
	return ret0
}

// ObjectStore indicates an expected call of ObjectStore.
func (mr *MockMetadataServiceMockRecorder) ObjectStore() *MockMetadataServiceObjectStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStore", reflect.TypeOf((*MockMetadataService)(nil).ObjectStore))
	return &MockMetadataServiceObjectStoreCall{Call: call}
}

// MockMetadataServiceObjectStoreCall wrap *gomock.Call
type MockMetadataServiceObjectStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMetadataServiceObjectStoreCall) Return(arg0 objectstore.ObjectStoreMetadata) *MockMetadataServiceObjectStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMetadataServiceObjectStoreCall) Do(f func() objectstore.ObjectStoreMetadata) *MockMetadataServiceObjectStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMetadataServiceObjectStoreCall) DoAndReturn(f func() objectstore.ObjectStoreMetadata) *MockMetadataServiceObjectStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelClaimGetter is a mock of ModelClaimGetter interface.
type MockModelClaimGetter struct {
	ctrl     *gomock.Controller
	recorder *MockModelClaimGetterMockRecorder
}

// MockModelClaimGetterMockRecorder is the mock recorder for MockModelClaimGetter.
type MockModelClaimGetterMockRecorder struct {
	mock *MockModelClaimGetter
}

// NewMockModelClaimGetter creates a new mock instance.
func NewMockModelClaimGetter(ctrl *gomock.Controller) *MockModelClaimGetter {
	mock := &MockModelClaimGetter{ctrl: ctrl}
	mock.recorder = &MockModelClaimGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelClaimGetter) EXPECT() *MockModelClaimGetterMockRecorder {
	return m.recorder
}

// ForModelUUID mocks base method.
func (m *MockModelClaimGetter) ForModelUUID(arg0 string) (objectstore0.Claimer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForModelUUID", arg0)
	ret0, _ := ret[0].(objectstore0.Claimer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForModelUUID indicates an expected call of ForModelUUID.
func (mr *MockModelClaimGetterMockRecorder) ForModelUUID(arg0 any) *MockModelClaimGetterForModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForModelUUID", reflect.TypeOf((*MockModelClaimGetter)(nil).ForModelUUID), arg0)
	return &MockModelClaimGetterForModelUUIDCall{Call: call}
}

// MockModelClaimGetterForModelUUIDCall wrap *gomock.Call
type MockModelClaimGetterForModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelClaimGetterForModelUUIDCall) Return(arg0 objectstore0.Claimer, arg1 error) *MockModelClaimGetterForModelUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelClaimGetterForModelUUIDCall) Do(f func(string) (objectstore0.Claimer, error)) *MockModelClaimGetterForModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelClaimGetterForModelUUIDCall) DoAndReturn(f func(string) (objectstore0.Claimer, error)) *MockModelClaimGetterForModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}