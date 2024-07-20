// Code generated by MockGen. DO NOT EDIT.
// Source: ../service/user_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	entities "backend/internal/entities"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserServicer is a mock of UserServicer interface.
type MockUserServicer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServicerMockRecorder
}

// MockUserServicerMockRecorder is the mock recorder for MockUserServicer.
type MockUserServicerMockRecorder struct {
	mock *MockUserServicer
}

// NewMockUserServicer creates a new mock instance.
func NewMockUserServicer(ctrl *gomock.Controller) *MockUserServicer {
	mock := &MockUserServicer{ctrl: ctrl}
	mock.recorder = &MockUserServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServicer) EXPECT() *MockUserServicerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserServicer) CreateUser(ctx context.Context, user entities.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServicerMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServicer)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserServicer) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServicerMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServicer)(nil).DeleteUser), ctx, id)
}

// GetUserByID mocks base method.
func (m *MockUserServicer) GetUserByID(ctx context.Context, id int) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServicerMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserServicer)(nil).GetUserByID), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserServicer) GetUsers(ctx context.Context, limit, offset int) ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, limit, offset)
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServicerMockRecorder) GetUsers(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserServicer)(nil).GetUsers), ctx, limit, offset)
}

// UpdateUser mocks base method.
func (m *MockUserServicer) UpdateUser(ctx context.Context, user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServicerMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserServicer)(nil).UpdateUser), ctx, user)
}