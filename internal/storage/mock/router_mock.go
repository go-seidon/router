// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/router/router.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	storage "github.com/go-seidon/chariot/internal/storage"
	router "github.com/go-seidon/chariot/internal/storage/router"
	gomock "github.com/golang/mock/gomock"
)

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// CreateStorage mocks base method.
func (m *MockRouter) CreateStorage(ctx context.Context, p router.CreateStorageParam) (storage.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStorage", ctx, p)
	ret0, _ := ret[0].(storage.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStorage indicates an expected call of CreateStorage.
func (mr *MockRouterMockRecorder) CreateStorage(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStorage", reflect.TypeOf((*MockRouter)(nil).CreateStorage), ctx, p)
}
