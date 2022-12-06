// Code generated by MockGen. DO NOT EDIT.
// Source: internal/queueing/queueing.go

// Package mock_queueing is a generated GoMock package.
package mock_queueing

import (
	context "context"
	reflect "reflect"

	queueing "github.com/go-seidon/chariot/internal/queueing"
	gomock "github.com/golang/mock/gomock"
)

// MockQueueing is a mock of Queueing interface.
type MockQueueing struct {
	ctrl     *gomock.Controller
	recorder *MockQueueingMockRecorder
}

// MockQueueingMockRecorder is the mock recorder for MockQueueing.
type MockQueueingMockRecorder struct {
	mock *MockQueueing
}

// NewMockQueueing creates a new mock instance.
func NewMockQueueing(ctrl *gomock.Controller) *MockQueueing {
	mock := &MockQueueing{ctrl: ctrl}
	mock.recorder = &MockQueueingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueing) EXPECT() *MockQueueingMockRecorder {
	return m.recorder
}

// BindQueue mocks base method.
func (m *MockQueueing) BindQueue(ctx context.Context, p queueing.BindQueueParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindQueue", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindQueue indicates an expected call of BindQueue.
func (mr *MockQueueingMockRecorder) BindQueue(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQueue", reflect.TypeOf((*MockQueueing)(nil).BindQueue), ctx, p)
}

// DeclareExchange mocks base method.
func (m *MockQueueing) DeclareExchange(ctx context.Context, p queueing.DeclareExchangeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclareExchange", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclareExchange indicates an expected call of DeclareExchange.
func (mr *MockQueueingMockRecorder) DeclareExchange(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareExchange", reflect.TypeOf((*MockQueueing)(nil).DeclareExchange), ctx, p)
}

// DeclareQueue mocks base method.
func (m *MockQueueing) DeclareQueue(ctx context.Context, p queueing.DeclareQueueParam) (*queueing.DeclareQueueResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclareQueue", ctx, p)
	ret0, _ := ret[0].(*queueing.DeclareQueueResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeclareQueue indicates an expected call of DeclareQueue.
func (mr *MockQueueingMockRecorder) DeclareQueue(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareQueue", reflect.TypeOf((*MockQueueing)(nil).DeclareQueue), ctx, p)
}

// Init mocks base method.
func (m *MockQueueing) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockQueueingMockRecorder) Init(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockQueueing)(nil).Init), ctx)
}

// Publish mocks base method.
func (m *MockQueueing) Publish(ctx context.Context, p queueing.PublishParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockQueueingMockRecorder) Publish(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockQueueing)(nil).Publish), ctx, p)
}

// Subscribe mocks base method.
func (m *MockQueueing) Subscribe(ctx context.Context, p queueing.SubscribeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockQueueingMockRecorder) Subscribe(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockQueueing)(nil).Subscribe), ctx, p)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockManager) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockManagerMockRecorder) Init(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockManager)(nil).Init), ctx)
}

// MockExchange is a mock of Exchange interface.
type MockExchange struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeMockRecorder
}

// MockExchangeMockRecorder is the mock recorder for MockExchange.
type MockExchangeMockRecorder struct {
	mock *MockExchange
}

// NewMockExchange creates a new mock instance.
func NewMockExchange(ctrl *gomock.Controller) *MockExchange {
	mock := &MockExchange{ctrl: ctrl}
	mock.recorder = &MockExchangeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchange) EXPECT() *MockExchangeMockRecorder {
	return m.recorder
}

// BindQueue mocks base method.
func (m *MockExchange) BindQueue(ctx context.Context, p queueing.BindQueueParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindQueue", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindQueue indicates an expected call of BindQueue.
func (mr *MockExchangeMockRecorder) BindQueue(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQueue", reflect.TypeOf((*MockExchange)(nil).BindQueue), ctx, p)
}

// DeclareExchange mocks base method.
func (m *MockExchange) DeclareExchange(ctx context.Context, p queueing.DeclareExchangeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclareExchange", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclareExchange indicates an expected call of DeclareExchange.
func (mr *MockExchangeMockRecorder) DeclareExchange(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareExchange", reflect.TypeOf((*MockExchange)(nil).DeclareExchange), ctx, p)
}

// MockPubsub is a mock of Pubsub interface.
type MockPubsub struct {
	ctrl     *gomock.Controller
	recorder *MockPubsubMockRecorder
}

// MockPubsubMockRecorder is the mock recorder for MockPubsub.
type MockPubsubMockRecorder struct {
	mock *MockPubsub
}

// NewMockPubsub creates a new mock instance.
func NewMockPubsub(ctrl *gomock.Controller) *MockPubsub {
	mock := &MockPubsub{ctrl: ctrl}
	mock.recorder = &MockPubsubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPubsub) EXPECT() *MockPubsubMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockPubsub) Publish(ctx context.Context, p queueing.PublishParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockPubsubMockRecorder) Publish(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubsub)(nil).Publish), ctx, p)
}

// Subscribe mocks base method.
func (m *MockPubsub) Subscribe(ctx context.Context, p queueing.SubscribeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockPubsubMockRecorder) Subscribe(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPubsub)(nil).Subscribe), ctx, p)
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockMessage) Ack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *MockMessageMockRecorder) Ack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockMessage)(nil).Ack))
}

// GetBody mocks base method.
func (m *MockMessage) GetBody() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBody indicates an expected call of GetBody.
func (mr *MockMessageMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockMessage)(nil).GetBody))
}

// GetId mocks base method.
func (m *MockMessage) GetId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockMessageMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockMessage)(nil).GetId))
}

// Nack mocks base method.
func (m *MockMessage) Nack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack.
func (mr *MockMessageMockRecorder) Nack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockMessage)(nil).Nack))
}

// MockQueue is a mock of Queue interface.
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue.
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance.
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// DeclareQueue mocks base method.
func (m *MockQueue) DeclareQueue(ctx context.Context, p queueing.DeclareQueueParam) (*queueing.DeclareQueueResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclareQueue", ctx, p)
	ret0, _ := ret[0].(*queueing.DeclareQueueResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeclareQueue indicates an expected call of DeclareQueue.
func (mr *MockQueueMockRecorder) DeclareQueue(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareQueue", reflect.TypeOf((*MockQueue)(nil).DeclareQueue), ctx, p)
}
