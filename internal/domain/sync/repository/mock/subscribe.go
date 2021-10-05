// Code generated by MockGen. DO NOT EDIT.
// Source: ./subscribe.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/scriptscat/cloudcat/internal/domain/sync/dto"
	entity "github.com/scriptscat/cloudcat/internal/domain/sync/entity"
)

// MockSubscribe is a mock of Subscribe interface.
type MockSubscribe struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribeMockRecorder
}

// MockSubscribeMockRecorder is the mock recorder for MockSubscribe.
type MockSubscribeMockRecorder struct {
	mock *MockSubscribe
}

// NewMockSubscribe creates a new mock instance.
func NewMockSubscribe(ctrl *gomock.Controller) *MockSubscribe {
	mock := &MockSubscribe{ctrl: ctrl}
	mock.recorder = &MockSubscribeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribe) EXPECT() *MockSubscribeMockRecorder {
	return m.recorder
}

// ActionList mocks base method.
func (m *MockSubscribe) ActionList(user, device, version int64) ([][]*dto.SyncSubscribe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionList", user, device, version)
	ret0, _ := ret[0].([][]*dto.SyncSubscribe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionList indicates an expected call of ActionList.
func (mr *MockSubscribeMockRecorder) ActionList(user, device, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionList", reflect.TypeOf((*MockSubscribe)(nil).ActionList), user, device, version)
}

// FindByUrl mocks base method.
func (m *MockSubscribe) FindByUrl(user, device int64, url string) (*entity.SyncSubscribe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUrl", user, device, url)
	ret0, _ := ret[0].(*entity.SyncSubscribe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUrl indicates an expected call of FindByUrl.
func (mr *MockSubscribeMockRecorder) FindByUrl(user, device, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUrl", reflect.TypeOf((*MockSubscribe)(nil).FindByUrl), user, device, url)
}

// LatestVersion mocks base method.
func (m *MockSubscribe) LatestVersion(user, device int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestVersion", user, device)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestVersion indicates an expected call of LatestVersion.
func (mr *MockSubscribeMockRecorder) LatestVersion(user, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestVersion", reflect.TypeOf((*MockSubscribe)(nil).LatestVersion), user, device)
}

// PushVersion mocks base method.
func (m *MockSubscribe) PushVersion(user, device int64, data []*dto.SyncSubscribe) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushVersion", user, device, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushVersion indicates an expected call of PushVersion.
func (mr *MockSubscribeMockRecorder) PushVersion(user, device, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushVersion", reflect.TypeOf((*MockSubscribe)(nil).PushVersion), user, device, data)
}

// Save mocks base method.
func (m *MockSubscribe) Save(entity *entity.SyncSubscribe) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockSubscribeMockRecorder) Save(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSubscribe)(nil).Save), entity)
}

// SetStatus mocks base method.
func (m *MockSubscribe) SetStatus(id int64, status int8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockSubscribeMockRecorder) SetStatus(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockSubscribe)(nil).SetStatus), id, status)
}
