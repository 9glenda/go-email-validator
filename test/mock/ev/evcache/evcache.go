// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ev/evcache/evcache.go

// Package mock_evcache is a generated GoMock package.
package mock_evcache

import (
	store "github.com/eko/gocache/store"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockInterface) Get(key interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), key)
}

// Set mocks base method
func (m *MockInterface) Set(key, object interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockInterfaceMockRecorder) Set(key, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInterface)(nil).Set), key, object)
}

// MockMarshaler is a mock of Marshaler interface
type MockMarshaler struct {
	ctrl     *gomock.Controller
	recorder *MockMarshalerMockRecorder
}

// MockMarshalerMockRecorder is the mock recorder for MockMarshaler
type MockMarshalerMockRecorder struct {
	mock *MockMarshaler
}

// NewMockMarshaler creates a new mock instance
func NewMockMarshaler(ctrl *gomock.Controller) *MockMarshaler {
	mock := &MockMarshaler{ctrl: ctrl}
	mock.recorder = &MockMarshalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarshaler) EXPECT() *MockMarshalerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMarshaler) Get(key, returnObj interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key, returnObj)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMarshalerMockRecorder) Get(key, returnObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMarshaler)(nil).Get), key, returnObj)
}

// Set mocks base method
func (m *MockMarshaler) Set(key, object interface{}, options *store.Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, object, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockMarshalerMockRecorder) Set(key, object, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMarshaler)(nil).Set), key, object, options)
}

// Delete mocks base method
func (m *MockMarshaler) Delete(key interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMarshalerMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMarshaler)(nil).Delete), key)
}

// Invalidate mocks base method
func (m *MockMarshaler) Invalidate(options store.InvalidateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidate", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate
func (mr *MockMarshalerMockRecorder) Invalidate(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockMarshaler)(nil).Invalidate), options)
}

// Clear mocks base method
func (m *MockMarshaler) Clear() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockMarshalerMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockMarshaler)(nil).Clear))
}