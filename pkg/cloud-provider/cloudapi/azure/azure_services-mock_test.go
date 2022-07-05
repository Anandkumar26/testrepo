// // Copyright 2022 Antrea Authors.
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //      http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud-provider/cloudapi/azure/azure_services.go

// Package azure is a generated GoMock package.
package azure

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockazureServiceClientCreateInterface is a mock of azureServiceClientCreateInterface interface.
type MockazureServiceClientCreateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockazureServiceClientCreateInterfaceMockRecorder
}

// MockazureServiceClientCreateInterfaceMockRecorder is the mock recorder for MockazureServiceClientCreateInterface.
type MockazureServiceClientCreateInterfaceMockRecorder struct {
	mock *MockazureServiceClientCreateInterface
}

// NewMockazureServiceClientCreateInterface creates a new mock instance.
func NewMockazureServiceClientCreateInterface(ctrl *gomock.Controller) *MockazureServiceClientCreateInterface {
	mock := &MockazureServiceClientCreateInterface{ctrl: ctrl}
	mock.recorder = &MockazureServiceClientCreateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockazureServiceClientCreateInterface) EXPECT() *MockazureServiceClientCreateInterfaceMockRecorder {
	return m.recorder
}

// applicationSecurityGroups mocks base method.
func (m *MockazureServiceClientCreateInterface) applicationSecurityGroups(subscriptionID string) (azureAsgWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "applicationSecurityGroups", subscriptionID)
	ret0, _ := ret[0].(azureAsgWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// applicationSecurityGroups indicates an expected call of applicationSecurityGroups.
func (mr *MockazureServiceClientCreateInterfaceMockRecorder) applicationSecurityGroups(subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "applicationSecurityGroups", reflect.TypeOf((*MockazureServiceClientCreateInterface)(nil).applicationSecurityGroups), subscriptionID)
}

// networkInterfaces mocks base method.
func (m *MockazureServiceClientCreateInterface) networkInterfaces(subscriptionID string) (azureNwIntfWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "networkInterfaces", subscriptionID)
	ret0, _ := ret[0].(azureNwIntfWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// networkInterfaces indicates an expected call of networkInterfaces.
func (mr *MockazureServiceClientCreateInterfaceMockRecorder) networkInterfaces(subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "networkInterfaces", reflect.TypeOf((*MockazureServiceClientCreateInterface)(nil).networkInterfaces), subscriptionID)
}

// resourceGraph mocks base method.
func (m *MockazureServiceClientCreateInterface) resourceGraph() (azureResourceGraphWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "resourceGraph")
	ret0, _ := ret[0].(azureResourceGraphWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// resourceGraph indicates an expected call of resourceGraph.
func (mr *MockazureServiceClientCreateInterfaceMockRecorder) resourceGraph() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "resourceGraph", reflect.TypeOf((*MockazureServiceClientCreateInterface)(nil).resourceGraph))
}

// securityGroups mocks base method.
func (m *MockazureServiceClientCreateInterface) securityGroups(subscriptionID string) (azureNsgWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "securityGroups", subscriptionID)
	ret0, _ := ret[0].(azureNsgWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// securityGroups indicates an expected call of securityGroups.
func (mr *MockazureServiceClientCreateInterfaceMockRecorder) securityGroups(subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "securityGroups", reflect.TypeOf((*MockazureServiceClientCreateInterface)(nil).securityGroups), subscriptionID)
}

// virtualNetworks mocks base method.
func (m *MockazureServiceClientCreateInterface) virtualNetworks(subscriptionID string) (azureVirtualNetworksWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "virtualNetworks", subscriptionID)
	ret0, _ := ret[0].(azureVirtualNetworksWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// virtualNetworks indicates an expected call of virtualNetworks.
func (mr *MockazureServiceClientCreateInterfaceMockRecorder) virtualNetworks(subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "virtualNetworks", reflect.TypeOf((*MockazureServiceClientCreateInterface)(nil).virtualNetworks), subscriptionID)
}

// MockazureServicesHelper is a mock of azureServicesHelper interface.
type MockazureServicesHelper struct {
	ctrl     *gomock.Controller
	recorder *MockazureServicesHelperMockRecorder
}

// MockazureServicesHelperMockRecorder is the mock recorder for MockazureServicesHelper.
type MockazureServicesHelperMockRecorder struct {
	mock *MockazureServicesHelper
}

// NewMockazureServicesHelper creates a new mock instance.
func NewMockazureServicesHelper(ctrl *gomock.Controller) *MockazureServicesHelper {
	mock := &MockazureServicesHelper{ctrl: ctrl}
	mock.recorder = &MockazureServicesHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockazureServicesHelper) EXPECT() *MockazureServicesHelperMockRecorder {
	return m.recorder
}

// newServiceSdkConfigProvider mocks base method.
func (m *MockazureServicesHelper) newServiceSdkConfigProvider(accCfg *azureAccountConfig) (azureServiceClientCreateInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "newServiceSdkConfigProvider", accCfg)
	ret0, _ := ret[0].(azureServiceClientCreateInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// newServiceSdkConfigProvider indicates an expected call of newServiceSdkConfigProvider.
func (mr *MockazureServicesHelperMockRecorder) newServiceSdkConfigProvider(accCfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newServiceSdkConfigProvider", reflect.TypeOf((*MockazureServicesHelper)(nil).newServiceSdkConfigProvider), accCfg)
}
