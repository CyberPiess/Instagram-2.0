// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	user "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
	gomock "github.com/golang/mock/gomock"
)

// MockuserService is a mock of userService interface.
type MockuserService struct {
	ctrl     *gomock.Controller
	recorder *MockuserServiceMockRecorder
}

// VerifyToken implements application.userService.
func (*MockuserService) VerifyToken(tokenString string) (*user.MyJWTClaims, error) {
	panic("unimplemented")
}

// MockuserServiceMockRecorder is the mock recorder for MockuserService.
type MockuserServiceMockRecorder struct {
	mock *MockuserService
}

// NewMockuserService creates a new mock instance.
func NewMockuserService(ctrl *gomock.Controller) *MockuserService {
	mock := &MockuserService{ctrl: ctrl}
	mock.recorder = &MockuserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserService) EXPECT() *MockuserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockuserService) CreateUser(newUser user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockuserServiceMockRecorder) CreateUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockuserService)(nil).CreateUser), newUser)
}

// LoginUser mocks base method.
func (m *MockuserService) LoginUser(logUser *user.LoginUserReq) (*user.LoginUserRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", logUser)
	ret0, _ := ret[0].(*user.LoginUserRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockuserServiceMockRecorder) LoginUser(logUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockuserService)(nil).LoginUser), logUser)
}

// VerifyData mocks base method.
func (m *MockuserService) VerifyData(newUser user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyData", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyData indicates an expected call of VerifyData.
func (mr *MockuserServiceMockRecorder) VerifyData(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyData", reflect.TypeOf((*MockuserService)(nil).VerifyData), newUser)
}
