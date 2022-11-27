// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	storage "github.com/go-seidon/chariot/internal/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// RetrieveObject mocks base method.
func (m *MockStorage) RetrieveObject(ctx context.Context, p storage.RetrieveObjectParam) (*storage.RetrieveObjectResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveObject", ctx, p)
	ret0, _ := ret[0].(*storage.RetrieveObjectResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveObject indicates an expected call of RetrieveObject.
func (mr *MockStorageMockRecorder) RetrieveObject(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveObject", reflect.TypeOf((*MockStorage)(nil).RetrieveObject), ctx, p)
}

// UploadObject mocks base method.
func (m *MockStorage) UploadObject(ctx context.Context, p storage.UploadObjectParam) (*storage.UploadObjectResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadObject", ctx, p)
	ret0, _ := ret[0].(*storage.UploadObjectResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadObject indicates an expected call of UploadObject.
func (mr *MockStorageMockRecorder) UploadObject(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadObject", reflect.TypeOf((*MockStorage)(nil).UploadObject), ctx, p)
}
