// Code generated by MockGen. DO NOT EDIT.
// Source: store/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	store "github.com/eko/gocache/v3/store"
	gomock "github.com/golang/mock/gomock"
)

// MockStoreInterface is a mock of StoreInterface interface.
type MockStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStoreInterfaceMockRecorder
}

// MockStoreInterfaceMockRecorder is the mock recorder for MockStoreInterface.
type MockStoreInterfaceMockRecorder struct {
	mock *MockStoreInterface
}

// NewMockStoreInterface creates a new mock instance.
func NewMockStoreInterface(ctrl *gomock.Controller) *MockStoreInterface {
	mock := &MockStoreInterface{ctrl: ctrl}
	mock.recorder = &MockStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreInterface) EXPECT() *MockStoreInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockStoreInterface) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockStoreInterfaceMockRecorder) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockStoreInterface)(nil).Clear), ctx)
}

// Delete mocks base method.
func (m *MockStoreInterface) Delete(ctx context.Context, key any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreInterfaceMockRecorder) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStoreInterface)(nil).Delete), ctx, key)
}

// Get mocks base method.
func (m *MockStoreInterface) Get(ctx context.Context, key any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStoreInterface)(nil).Get), ctx, key)
}

// GetType mocks base method.
func (m *MockStoreInterface) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockStoreInterfaceMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockStoreInterface)(nil).GetType))
}

// GetWithTTL mocks base method.
func (m *MockStoreInterface) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithTTL", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithTTL indicates an expected call of GetWithTTL.
func (mr *MockStoreInterfaceMockRecorder) GetWithTTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTTL", reflect.TypeOf((*MockStoreInterface)(nil).GetWithTTL), ctx, key)
}

// Invalidate mocks base method.
func (m *MockStoreInterface) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidate", ctx, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockStoreInterfaceMockRecorder) Invalidate(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockStoreInterface)(nil).Invalidate), ctx, options)
}

// Set mocks base method.
func (m *MockStoreInterface) Set(ctx context.Context, key, value any, options *store.Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStoreInterfaceMockRecorder) Set(ctx, key, value, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStoreInterface)(nil).Set), ctx, key, value, options)
}
