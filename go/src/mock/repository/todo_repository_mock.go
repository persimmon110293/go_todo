// Code generated by MockGen. DO NOT EDIT.
// Source: repository/todo_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	repository "main/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITodoRepository is a mock of ITodoRepository interface.
type MockITodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITodoRepositoryMockRecorder
}

// CreateTodo implements repository.ITodoRepository.
func (m *MockITodoRepository) CreateTodo(map[string]string) (*repository.Todo, error) {
	panic("unimplemented")
}

// DeleteTodoById implements repository.ITodoRepository.
func (m *MockITodoRepository) DeleteTodoById(string) error {
	panic("unimplemented")
}

// GetTodoById implements repository.ITodoRepository.
func (m *MockITodoRepository) GetTodoById(string) (*repository.Todo, error) {
	panic("unimplemented")
}

// UpdateTodoById implements repository.ITodoRepository.
func (m *MockITodoRepository) UpdateTodoById(string, map[string]string) (*repository.Todo, error) {
	panic("unimplemented")
}

// MockITodoRepositoryMockRecorder is the mock recorder for MockITodoRepository.
type MockITodoRepositoryMockRecorder struct {
	mock *MockITodoRepository
}

// NewMockITodoRepository creates a new mock instance.
func NewMockITodoRepository(ctrl *gomock.Controller) *MockITodoRepository {
	mock := &MockITodoRepository{ctrl: ctrl}
	mock.recorder = &MockITodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITodoRepository) EXPECT() *MockITodoRepositoryMockRecorder {
	return m.recorder
}

// GetAllTodo mocks base method.
func (m *MockITodoRepository) GetAllTodo() ([]repository.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodo")
	ret0, _ := ret[0].([]repository.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodo indicates an expected call of GetAllTodo.
func (mr *MockITodoRepositoryMockRecorder) GetAllTodo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodo", reflect.TypeOf((*MockITodoRepository)(nil).GetAllTodo))
}
