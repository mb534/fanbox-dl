// Code generated by MockGen. DO NOT EDIT.
// Source: file_storage.go

// Package fanbox is a generated GoMock package.
package fanbox

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileStorage is a mock of FileStorage interface.
type MockFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageMockRecorder
}

// MockFileStorageMockRecorder is the mock recorder for MockFileStorage.
type MockFileStorageMockRecorder struct {
	mock *MockFileStorage
}

// NewMockFileStorage creates a new mock instance.
func NewMockFileStorage(ctrl *gomock.Controller) *MockFileStorage {
	mock := &MockFileStorage{ctrl: ctrl}
	mock.recorder = &MockFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorage) EXPECT() *MockFileStorageMockRecorder {
	return m.recorder
}

// Exist mocks base method.
func (m *MockFileStorage) Exist(post Post, order int, file File) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", post, order, file)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exist indicates an expected call of Exist.
func (mr *MockFileStorageMockRecorder) Exist(post, order, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockFileStorage)(nil).Exist), post, order, file)
}

// Save mocks base method.
func (m *MockFileStorage) Save(post Post, order int, file File, r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", post, order, file, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockFileStorageMockRecorder) Save(post, order, file, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileStorage)(nil).Save), post, order, file, r)
}
