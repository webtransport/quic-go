// Code generated by MockGen. DO NOT EDIT.
// Source: multiplexer.go

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logging "github.com/btwiuse/quic/logging"
)

// MockMultiplexer is a mock of Multiplexer interface.
type MockMultiplexer struct {
	ctrl     *gomock.Controller
	recorder *MockMultiplexerMockRecorder
}

// MockMultiplexerMockRecorder is the mock recorder for MockMultiplexer.
type MockMultiplexerMockRecorder struct {
	mock *MockMultiplexer
}

// NewMockMultiplexer creates a new mock instance.
func NewMockMultiplexer(ctrl *gomock.Controller) *MockMultiplexer {
	mock := &MockMultiplexer{ctrl: ctrl}
	mock.recorder = &MockMultiplexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiplexer) EXPECT() *MockMultiplexerMockRecorder {
	return m.recorder
}

// AddConn mocks base method.
func (m *MockMultiplexer) AddConn(c net.PacketConn, connIDLen int, statelessResetKey []byte, tracer logging.Tracer) (packetHandlerManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddConn", c, connIDLen, statelessResetKey, tracer)
	ret0, _ := ret[0].(packetHandlerManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddConn indicates an expected call of AddConn.
func (mr *MockMultiplexerMockRecorder) AddConn(c, connIDLen, statelessResetKey, tracer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConn", reflect.TypeOf((*MockMultiplexer)(nil).AddConn), c, connIDLen, statelessResetKey, tracer)
}

// RemoveConn mocks base method.
func (m *MockMultiplexer) RemoveConn(arg0 indexableConn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveConn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConn indicates an expected call of RemoveConn.
func (mr *MockMultiplexerMockRecorder) RemoveConn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConn", reflect.TypeOf((*MockMultiplexer)(nil).RemoveConn), arg0)
}
