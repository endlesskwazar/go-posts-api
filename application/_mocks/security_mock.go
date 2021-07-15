// Code generated by MockGen. DO NOT EDIT.
// Source: infrustructure/security/token_security.go

// Package mock_security is a generated GoMock package.
package mock

import (
	entity "go-cource-api/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenSecurity is a mock of TokenSecurity interface.
type MockTokenSecurity struct {
	ctrl     *gomock.Controller
	recorder *MockTokenSecurityMockRecorder
}

// MockTokenSecurityMockRecorder is the mock recorder for MockTokenSecurity.
type MockTokenSecurityMockRecorder struct {
	mock *MockTokenSecurity
}

// NewMockTokenSecurity creates a new mock instance.
func NewMockTokenSecurity(ctrl *gomock.Controller) *MockTokenSecurity {
	mock := &MockTokenSecurity{ctrl: ctrl}
	mock.recorder = &MockTokenSecurityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenSecurity) EXPECT() *MockTokenSecurityMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockTokenSecurity) FindUserByEmail(email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockTokenSecurityMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockTokenSecurity)(nil).FindUserByEmail), email)
}

// GenerateToken mocks base method.
func (m *MockTokenSecurity) GenerateToken(arg0 entity.User) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenSecurityMockRecorder) GenerateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenSecurity)(nil).GenerateToken), arg0)
}

// HashPassword mocks base method.
func (m *MockTokenSecurity) HashPassword(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockTokenSecurityMockRecorder) HashPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockTokenSecurity)(nil).HashPassword), arg0)
}

// LoginUser mocks base method.
func (m *MockTokenSecurity) LoginUser(arg0, arg1 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockTokenSecurityMockRecorder) LoginUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockTokenSecurity)(nil).LoginUser), arg0, arg1)
}

// RegisterUser mocks base method.
func (m *MockTokenSecurity) RegisterUser(user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockTokenSecurityMockRecorder) RegisterUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockTokenSecurity)(nil).RegisterUser), user)
}

// VerifyPassword mocks base method.
func (m *MockTokenSecurity) VerifyPassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyPassword indicates an expected call of VerifyPassword.
func (mr *MockTokenSecurityMockRecorder) VerifyPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPassword", reflect.TypeOf((*MockTokenSecurity)(nil).VerifyPassword), arg0, arg1)
}
