// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ev/evsmtp/smtp.go

// Package evsmtp is a generated GoMock package.
package evsmtp

import (
	evmail "github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRandomRCPT is a mock of RandomRCPT interface
type MockRandomRCPT struct {
	ctrl     *gomock.Controller
	recorder *MockRandomRCPTMockRecorder
}

// MockRandomRCPTMockRecorder is the mock recorder for MockRandomRCPT
type MockRandomRCPTMockRecorder struct {
	mock *MockRandomRCPT
}

// NewMockRandomRCPT creates a new mock instance
func NewMockRandomRCPT(ctrl *gomock.Controller) *MockRandomRCPT {
	mock := &MockRandomRCPT{ctrl: ctrl}
	mock.recorder = &MockRandomRCPTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRandomRCPT) EXPECT() *MockRandomRCPTMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockRandomRCPT) Call(email evmail.Address) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", email)
	ret0, _ := ret[0].([]error)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockRandomRCPTMockRecorder) Call(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRandomRCPT)(nil).Call), email)
}

// set mocks base method
func (m *MockRandomRCPT) set(fn RandomRCPTFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "set", fn)
}

// set indicates an expected call of set
func (mr *MockRandomRCPTMockRecorder) set(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "set", reflect.TypeOf((*MockRandomRCPT)(nil).set), fn)
}

// get mocks base method
func (m *MockRandomRCPT) get() RandomRCPTFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "get")
	ret0, _ := ret[0].(RandomRCPTFunc)
	return ret0
}

// get indicates an expected call of get
func (mr *MockRandomRCPTMockRecorder) get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "get", reflect.TypeOf((*MockRandomRCPT)(nil).get))
}

// MockChecker is a mock of Checker interface
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockChecker) Validate(mxs MXs, input Input) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", mxs, input)
	ret0, _ := ret[0].([]error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockCheckerMockRecorder) Validate(mxs, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockChecker)(nil).Validate), mxs, input)
}

// MockCheckerWithRandomRCPT is a mock of CheckerWithRandomRCPT interface
type MockCheckerWithRandomRCPT struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerWithRandomRCPTMockRecorder
}

// MockCheckerWithRandomRCPTMockRecorder is the mock recorder for MockCheckerWithRandomRCPT
type MockCheckerWithRandomRCPTMockRecorder struct {
	mock *MockCheckerWithRandomRCPT
}

// NewMockCheckerWithRandomRCPT creates a new mock instance
func NewMockCheckerWithRandomRCPT(ctrl *gomock.Controller) *MockCheckerWithRandomRCPT {
	mock := &MockCheckerWithRandomRCPT{ctrl: ctrl}
	mock.recorder = &MockCheckerWithRandomRCPTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckerWithRandomRCPT) EXPECT() *MockCheckerWithRandomRCPTMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockCheckerWithRandomRCPT) Validate(mxs MXs, input Input) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", mxs, input)
	ret0, _ := ret[0].([]error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockCheckerWithRandomRCPTMockRecorder) Validate(mxs, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockCheckerWithRandomRCPT)(nil).Validate), mxs, input)
}

// Call mocks base method
func (m *MockCheckerWithRandomRCPT) Call(email evmail.Address) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", email)
	ret0, _ := ret[0].([]error)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockCheckerWithRandomRCPTMockRecorder) Call(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockCheckerWithRandomRCPT)(nil).Call), email)
}

// set mocks base method
func (m *MockCheckerWithRandomRCPT) set(fn RandomRCPTFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "set", fn)
}

// set indicates an expected call of set
func (mr *MockCheckerWithRandomRCPTMockRecorder) set(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "set", reflect.TypeOf((*MockCheckerWithRandomRCPT)(nil).set), fn)
}

// get mocks base method
func (m *MockCheckerWithRandomRCPT) get() RandomRCPTFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "get")
	ret0, _ := ret[0].(RandomRCPTFunc)
	return ret0
}

// get indicates an expected call of get
func (mr *MockCheckerWithRandomRCPTMockRecorder) get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "get", reflect.TypeOf((*MockCheckerWithRandomRCPT)(nil).get))
}
