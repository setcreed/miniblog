// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/setcreed/miniblog/internal/miniblog/biz (interfaces: IBiz)
//
// Generated by this command:
//
//	mockgen -destination mock_biz.go -package biz github.com/setcreed/miniblog/internal/miniblog/biz IBiz
//

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	post "github.com/setcreed/miniblog/internal/miniblog/biz/post"
	user "github.com/setcreed/miniblog/internal/miniblog/biz/user"
)

// MockIBiz is a mock of IBiz interface.
type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

// MockIBizMockRecorder is the mock recorder for MockIBiz.
type MockIBizMockRecorder struct {
	mock *MockIBiz
}

// NewMockIBiz creates a new mock instance.
func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}

// Posts mocks base method.
func (m *MockIBiz) Posts() post.PostBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Posts")
	ret0, _ := ret[0].(post.PostBiz)
	return ret0
}

// Posts indicates an expected call of Posts.
func (mr *MockIBizMockRecorder) Posts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Posts", reflect.TypeOf((*MockIBiz)(nil).Posts))
}

// Users mocks base method.
func (m *MockIBiz) Users() user.UserBiz {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(user.UserBiz)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockIBizMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockIBiz)(nil).Users))
}
