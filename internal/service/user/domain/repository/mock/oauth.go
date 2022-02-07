// Code generated by MockGen. DO NOT EDIT.
// Source: ./oauth.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/scriptscat/cloudcat/internal/service/user/domain/entity"
)

// MockBBSOAuth is a mock of BBSOAuth interface.
type MockBBSOAuth struct {
	ctrl     *gomock.Controller
	recorder *MockBBSOAuthMockRecorder
}

// MockBBSOAuthMockRecorder is the mock recorder for MockBBSOAuth.
type MockBBSOAuthMockRecorder struct {
	mock *MockBBSOAuth
}

// NewMockBBSOAuth creates a new mock instance.
func NewMockBBSOAuth(ctrl *gomock.Controller) *MockBBSOAuth {
	mock := &MockBBSOAuth{ctrl: ctrl}
	mock.recorder = &MockBBSOAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBBSOAuth) EXPECT() *MockBBSOAuthMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBBSOAuth) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBBSOAuthMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBBSOAuth)(nil).Delete), id)
}

// FindByOpenid mocks base method.
func (m *MockBBSOAuth) FindByOpenid(openid string) (*entity.BbsOauthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOpenid", openid)
	ret0, _ := ret[0].(*entity.BbsOauthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOpenid indicates an expected call of FindByOpenid.
func (mr *MockBBSOAuthMockRecorder) FindByOpenid(openid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOpenid", reflect.TypeOf((*MockBBSOAuth)(nil).FindByOpenid), openid)
}

// FindByUid mocks base method.
func (m *MockBBSOAuth) FindByUid(uid int64) (*entity.BbsOauthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUid", uid)
	ret0, _ := ret[0].(*entity.BbsOauthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUid indicates an expected call of FindByUid.
func (mr *MockBBSOAuthMockRecorder) FindByUid(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUid", reflect.TypeOf((*MockBBSOAuth)(nil).FindByUid), uid)
}

// Save mocks base method.
func (m *MockBBSOAuth) Save(bbs *entity.BbsOauthUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", bbs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockBBSOAuthMockRecorder) Save(bbs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBBSOAuth)(nil).Save), bbs)
}

// MockWechatOAuth is a mock of WechatOAuth interface.
type MockWechatOAuth struct {
	ctrl     *gomock.Controller
	recorder *MockWechatOAuthMockRecorder
}

// MockWechatOAuthMockRecorder is the mock recorder for MockWechatOAuth.
type MockWechatOAuthMockRecorder struct {
	mock *MockWechatOAuth
}

// NewMockWechatOAuth creates a new mock instance.
func NewMockWechatOAuth(ctrl *gomock.Controller) *MockWechatOAuth {
	mock := &MockWechatOAuth{ctrl: ctrl}
	mock.recorder = &MockWechatOAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWechatOAuth) EXPECT() *MockWechatOAuthMockRecorder {
	return m.recorder
}

// BindCodeUid mocks base method.
func (m *MockWechatOAuth) BindCodeUid(code string, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindCodeUid", code, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindCodeUid indicates an expected call of BindCodeUid.
func (mr *MockWechatOAuthMockRecorder) BindCodeUid(code, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindCodeUid", reflect.TypeOf((*MockWechatOAuth)(nil).BindCodeUid), code, uid)
}

// DelCode mocks base method.
func (m *MockWechatOAuth) DelCode(code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelCode", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelCode indicates an expected call of DelCode.
func (mr *MockWechatOAuthMockRecorder) DelCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelCode", reflect.TypeOf((*MockWechatOAuth)(nil).DelCode), code)
}

// Delete mocks base method.
func (m *MockWechatOAuth) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWechatOAuthMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWechatOAuth)(nil).Delete), id)
}

// FindByOpenid mocks base method.
func (m *MockWechatOAuth) FindByOpenid(openid string) (*entity.WechatOauthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOpenid", openid)
	ret0, _ := ret[0].(*entity.WechatOauthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOpenid indicates an expected call of FindByOpenid.
func (mr *MockWechatOAuthMockRecorder) FindByOpenid(openid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOpenid", reflect.TypeOf((*MockWechatOAuth)(nil).FindByOpenid), openid)
}

// FindByUid mocks base method.
func (m *MockWechatOAuth) FindByUid(uid int64) (*entity.WechatOauthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUid", uid)
	ret0, _ := ret[0].(*entity.WechatOauthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUid indicates an expected call of FindByUid.
func (mr *MockWechatOAuthMockRecorder) FindByUid(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUid", reflect.TypeOf((*MockWechatOAuth)(nil).FindByUid), uid)
}

// FindCodeUid mocks base method.
func (m *MockWechatOAuth) FindCodeUid(code string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCodeUid", code)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCodeUid indicates an expected call of FindCodeUid.
func (mr *MockWechatOAuthMockRecorder) FindCodeUid(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCodeUid", reflect.TypeOf((*MockWechatOAuth)(nil).FindCodeUid), code)
}

// Save mocks base method.
func (m *MockWechatOAuth) Save(u *entity.WechatOauthUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockWechatOAuthMockRecorder) Save(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockWechatOAuth)(nil).Save), u)
}