// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/setcreed/miniblog/internal/miniblog/store (interfaces: IStore,UserStore,PostStore)
//
// Generated by this command:
//
//	mockgen -destination mock_store.go -package store github.com/setcreed/miniblog/internal/miniblog/store IStore,UserStore,PostStore
//

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	model "github.com/setcreed/miniblog/internal/pkg/model"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIStore is a mock of IStore interface.
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

// MockIStoreMockRecorder is the mock recorder for MockIStore.
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance.
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// DB mocks base method.
func (m *MockIStore) DB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockIStoreMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockIStore)(nil).DB))
}

// Posts mocks base method.
func (m *MockIStore) Posts() PostStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Posts")
	ret0, _ := ret[0].(PostStore)
	return ret0
}

// Posts indicates an expected call of Posts.
func (mr *MockIStoreMockRecorder) Posts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Posts", reflect.TypeOf((*MockIStore)(nil).Posts))
}

// Users mocks base method.
func (m *MockIStore) Users() UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(UserStore)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockIStoreMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockIStore)(nil).Users))
}

// MockUserStore is a mock of UserStore interface.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore.
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance.
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserStore) Create(arg0 context.Context, arg1 *model.UserM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserStoreMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUserStore) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserStoreMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockUserStore) Get(arg0 context.Context, arg1 string) (*model.UserM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.UserM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockUserStore) List(arg0 context.Context, arg1, arg2 int) (int64, []*model.UserM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.UserM)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockUserStoreMockRecorder) List(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockUserStore) Update(arg0 context.Context, arg1 *model.UserM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserStoreMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserStore)(nil).Update), arg0, arg1)
}

// MockPostStore is a mock of PostStore interface.
type MockPostStore struct {
	ctrl     *gomock.Controller
	recorder *MockPostStoreMockRecorder
}

// MockPostStoreMockRecorder is the mock recorder for MockPostStore.
type MockPostStoreMockRecorder struct {
	mock *MockPostStore
}

// NewMockPostStore creates a new mock instance.
func NewMockPostStore(ctrl *gomock.Controller) *MockPostStore {
	mock := &MockPostStore{ctrl: ctrl}
	mock.recorder = &MockPostStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostStore) EXPECT() *MockPostStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostStore) Create(arg0 context.Context, arg1 *model.PostM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPostStoreMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockPostStore) Delete(arg0 context.Context, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostStoreMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostStore)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockPostStore) Get(arg0 context.Context, arg1, arg2 string) (*model.PostM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.PostM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostStoreMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostStore)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPostStore) List(arg0 context.Context, arg1 string, arg2, arg3 int) (int64, []*model.PostM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.PostM)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockPostStoreMockRecorder) List(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostStore)(nil).List), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockPostStore) Update(arg0 context.Context, arg1 *model.PostM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostStoreMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostStore)(nil).Update), arg0, arg1)
}