// Code generated by MockGen. DO NOT EDIT.
// Source: ./device.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/scriptscat/cloudcat/internal/domain/sync/entity"
)

// MockDevice is a mock of Device interface.
type MockDevice struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceMockRecorder
}

// MockDeviceMockRecorder is the mock recorder for MockDevice.
type MockDeviceMockRecorder struct {
	mock *MockDevice
}

// NewMockDevice creates a new mock instance.
func NewMockDevice(ctrl *gomock.Controller) *MockDevice {
	mock := &MockDevice{ctrl: ctrl}
	mock.recorder = &MockDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevice) EXPECT() *MockDeviceMockRecorder {
	return m.recorder
}

// FindById mocks base method.
func (m *MockDevice) FindById(id int64) (*entity.SyncDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*entity.SyncDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockDeviceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockDevice)(nil).FindById), id)
}

// ListDevice mocks base method.
func (m *MockDevice) ListDevice(user int64) ([]*entity.SyncDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDevice", user)
	ret0, _ := ret[0].([]*entity.SyncDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDevice indicates an expected call of ListDevice.
func (mr *MockDeviceMockRecorder) ListDevice(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevice", reflect.TypeOf((*MockDevice)(nil).ListDevice), user)
}

// Save mocks base method.
func (m *MockDevice) Save(device *entity.SyncDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", device)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockDeviceMockRecorder) Save(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDevice)(nil).Save), device)
}

// UpdateSetting mocks base method.
func (m *MockDevice) UpdateSetting(id int64, setting string, settingtime int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSetting", id, setting, settingtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSetting indicates an expected call of UpdateSetting.
func (mr *MockDeviceMockRecorder) UpdateSetting(id, setting, settingtime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSetting", reflect.TypeOf((*MockDevice)(nil).UpdateSetting), id, setting, settingtime)
}