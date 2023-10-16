// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/servicefactory (interfaces: ControllerServiceFactory,ModelServiceFactory,ServiceFactory,ServiceFactoryGetter)

// Package servicefactory is a generated GoMock package.
package servicefactory

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/autocert/service"
	service0 "github.com/juju/juju/domain/cloud/service"
	service1 "github.com/juju/juju/domain/controllerconfig/service"
	service2 "github.com/juju/juju/domain/controllernode/service"
	service3 "github.com/juju/juju/domain/credential/service"
	service4 "github.com/juju/juju/domain/externalcontroller/service"
	service5 "github.com/juju/juju/domain/model/service"
	service6 "github.com/juju/juju/domain/modelmanager/service"
	service7 "github.com/juju/juju/domain/upgrade/service"
	servicefactory "github.com/juju/juju/internal/servicefactory"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerServiceFactory is a mock of ControllerServiceFactory interface.
type MockControllerServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockControllerServiceFactoryMockRecorder
}

// MockControllerServiceFactoryMockRecorder is the mock recorder for MockControllerServiceFactory.
type MockControllerServiceFactoryMockRecorder struct {
	mock *MockControllerServiceFactory
}

// NewMockControllerServiceFactory creates a new mock instance.
func NewMockControllerServiceFactory(ctrl *gomock.Controller) *MockControllerServiceFactory {
	mock := &MockControllerServiceFactory{ctrl: ctrl}
	mock.recorder = &MockControllerServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerServiceFactory) EXPECT() *MockControllerServiceFactoryMockRecorder {
	return m.recorder
}

// AutocertCache mocks base method.
func (m *MockControllerServiceFactory) AutocertCache() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockControllerServiceFactoryMockRecorder) AutocertCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockControllerServiceFactory)(nil).AutocertCache))
}

// Cloud mocks base method.
func (m *MockControllerServiceFactory) Cloud() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockControllerServiceFactoryMockRecorder) Cloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockControllerServiceFactory)(nil).Cloud))
}

// ControllerConfig mocks base method.
func (m *MockControllerServiceFactory) ControllerConfig() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerConfig))
}

// ControllerNode mocks base method.
func (m *MockControllerServiceFactory) ControllerNode() *service2.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service2.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerNode))
}

// Credential mocks base method.
func (m *MockControllerServiceFactory) Credential() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockControllerServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockControllerServiceFactory)(nil).Credential))
}

// ExternalController mocks base method.
func (m *MockControllerServiceFactory) ExternalController() *service4.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service4.Service)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockControllerServiceFactoryMockRecorder) ExternalController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockControllerServiceFactory)(nil).ExternalController))
}

// Model mocks base method.
func (m *MockControllerServiceFactory) Model() *service5.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service5.Service)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockControllerServiceFactoryMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockControllerServiceFactory)(nil).Model))
}

// ModelManager mocks base method.
func (m *MockControllerServiceFactory) ModelManager() *service6.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelManager")
	ret0, _ := ret[0].(*service6.Service)
	return ret0
}

// ModelManager indicates an expected call of ModelManager.
func (mr *MockControllerServiceFactoryMockRecorder) ModelManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelManager", reflect.TypeOf((*MockControllerServiceFactory)(nil).ModelManager))
}

// Upgrade mocks base method.
func (m *MockControllerServiceFactory) Upgrade() *service7.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade")
	ret0, _ := ret[0].(*service7.Service)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockControllerServiceFactoryMockRecorder) Upgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockControllerServiceFactory)(nil).Upgrade))
}

// MockModelServiceFactory is a mock of ModelServiceFactory interface.
type MockModelServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockModelServiceFactoryMockRecorder
}

// MockModelServiceFactoryMockRecorder is the mock recorder for MockModelServiceFactory.
type MockModelServiceFactoryMockRecorder struct {
	mock *MockModelServiceFactory
}

// NewMockModelServiceFactory creates a new mock instance.
func NewMockModelServiceFactory(ctrl *gomock.Controller) *MockModelServiceFactory {
	mock := &MockModelServiceFactory{ctrl: ctrl}
	mock.recorder = &MockModelServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelServiceFactory) EXPECT() *MockModelServiceFactoryMockRecorder {
	return m.recorder
}

// TODO mocks base method.
func (m *MockModelServiceFactory) TODO() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TODO")
}

// TODO indicates an expected call of TODO.
func (mr *MockModelServiceFactoryMockRecorder) TODO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TODO", reflect.TypeOf((*MockModelServiceFactory)(nil).TODO))
}

// MockServiceFactory is a mock of ServiceFactory interface.
type MockServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFactoryMockRecorder
}

// MockServiceFactoryMockRecorder is the mock recorder for MockServiceFactory.
type MockServiceFactoryMockRecorder struct {
	mock *MockServiceFactory
}

// NewMockServiceFactory creates a new mock instance.
func NewMockServiceFactory(ctrl *gomock.Controller) *MockServiceFactory {
	mock := &MockServiceFactory{ctrl: ctrl}
	mock.recorder = &MockServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFactory) EXPECT() *MockServiceFactoryMockRecorder {
	return m.recorder
}

// AutocertCache mocks base method.
func (m *MockServiceFactory) AutocertCache() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockServiceFactoryMockRecorder) AutocertCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockServiceFactory)(nil).AutocertCache))
}

// Cloud mocks base method.
func (m *MockServiceFactory) Cloud() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockServiceFactoryMockRecorder) Cloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockServiceFactory)(nil).Cloud))
}

// ControllerConfig mocks base method.
func (m *MockServiceFactory) ControllerConfig() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockServiceFactoryMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockServiceFactory)(nil).ControllerConfig))
}

// ControllerNode mocks base method.
func (m *MockServiceFactory) ControllerNode() *service2.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service2.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockServiceFactoryMockRecorder) ControllerNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockServiceFactory)(nil).ControllerNode))
}

// Credential mocks base method.
func (m *MockServiceFactory) Credential() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockServiceFactory)(nil).Credential))
}

// ExternalController mocks base method.
func (m *MockServiceFactory) ExternalController() *service4.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service4.Service)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockServiceFactoryMockRecorder) ExternalController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockServiceFactory)(nil).ExternalController))
}

// Model mocks base method.
func (m *MockServiceFactory) Model() *service5.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service5.Service)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockServiceFactoryMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockServiceFactory)(nil).Model))
}

// ModelManager mocks base method.
func (m *MockServiceFactory) ModelManager() *service6.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelManager")
	ret0, _ := ret[0].(*service6.Service)
	return ret0
}

// ModelManager indicates an expected call of ModelManager.
func (mr *MockServiceFactoryMockRecorder) ModelManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelManager", reflect.TypeOf((*MockServiceFactory)(nil).ModelManager))
}

// TODO mocks base method.
func (m *MockServiceFactory) TODO() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TODO")
}

// TODO indicates an expected call of TODO.
func (mr *MockServiceFactoryMockRecorder) TODO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TODO", reflect.TypeOf((*MockServiceFactory)(nil).TODO))
}

// Upgrade mocks base method.
func (m *MockServiceFactory) Upgrade() *service7.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade")
	ret0, _ := ret[0].(*service7.Service)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockServiceFactoryMockRecorder) Upgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockServiceFactory)(nil).Upgrade))
}

// MockServiceFactoryGetter is a mock of ServiceFactoryGetter interface.
type MockServiceFactoryGetter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFactoryGetterMockRecorder
}

// MockServiceFactoryGetterMockRecorder is the mock recorder for MockServiceFactoryGetter.
type MockServiceFactoryGetterMockRecorder struct {
	mock *MockServiceFactoryGetter
}

// NewMockServiceFactoryGetter creates a new mock instance.
func NewMockServiceFactoryGetter(ctrl *gomock.Controller) *MockServiceFactoryGetter {
	mock := &MockServiceFactoryGetter{ctrl: ctrl}
	mock.recorder = &MockServiceFactoryGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFactoryGetter) EXPECT() *MockServiceFactoryGetterMockRecorder {
	return m.recorder
}

// FactoryForModel mocks base method.
func (m *MockServiceFactoryGetter) FactoryForModel(arg0 string) servicefactory.ServiceFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FactoryForModel", arg0)
	ret0, _ := ret[0].(servicefactory.ServiceFactory)
	return ret0
}

// FactoryForModel indicates an expected call of FactoryForModel.
func (mr *MockServiceFactoryGetterMockRecorder) FactoryForModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FactoryForModel", reflect.TypeOf((*MockServiceFactoryGetter)(nil).FactoryForModel), arg0)
}