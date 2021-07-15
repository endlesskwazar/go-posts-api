// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/comment_repository.go

// Package mock_repository is a generated GoMock package.
package mock

import (
	entity "go-cource-api/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCommentRepository) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockCommentRepository) FindAll() ([]entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCommentRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCommentRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockCommentRepository) FindById(id int64) (*entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCommentRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCommentRepository)(nil).FindById), id)
}

// FindByPostId mocks base method.
func (m *MockCommentRepository) FindByPostId(id int64) ([]entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPostId", id)
	ret0, _ := ret[0].([]entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPostId indicates an expected call of FindByPostId.
func (mr *MockCommentRepositoryMockRecorder) FindByPostId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPostId", reflect.TypeOf((*MockCommentRepository)(nil).FindByPostId), id)
}

// Save mocks base method.
func (m *MockCommentRepository) Save(comment *entity.Comment) (*entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", comment)
	ret0, _ := ret[0].(*entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCommentRepositoryMockRecorder) Save(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCommentRepository)(nil).Save), comment)
}

// Update mocks base method.
func (m *MockCommentRepository) Update(arg0 *entity.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCommentRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentRepository)(nil).Update), arg0)
}
