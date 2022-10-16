// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/btwiuse/quic/internal/protocol"
)

// MockPacketHandler is a mock of PacketHandler interface.
type MockPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerMockRecorder
}

// MockPacketHandlerMockRecorder is the mock recorder for MockPacketHandler.
type MockPacketHandlerMockRecorder struct {
	mock *MockPacketHandler
}

// NewMockPacketHandler creates a new mock instance.
func NewMockPacketHandler(ctrl *gomock.Controller) *MockPacketHandler {
	mock := &MockPacketHandler{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketHandler) EXPECT() *MockPacketHandlerMockRecorder {
	return m.recorder
}

// destroy mocks base method.
func (m *MockPacketHandler) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockPacketHandlerMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockPacketHandler)(nil).destroy), arg0)
}

// getPerspective mocks base method.
func (m *MockPacketHandler) getPerspective() protocol.Perspective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// getPerspective indicates an expected call of getPerspective.
func (mr *MockPacketHandlerMockRecorder) getPerspective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPerspective", reflect.TypeOf((*MockPacketHandler)(nil).getPerspective))
}

// handlePacket mocks base method.
func (m *MockPacketHandler) handlePacket(arg0 *receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockPacketHandlerMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockPacketHandler)(nil).handlePacket), arg0)
}

// shutdown mocks base method.
func (m *MockPacketHandler) shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "shutdown")
}

// shutdown indicates an expected call of shutdown.
func (mr *MockPacketHandlerMockRecorder) shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shutdown", reflect.TypeOf((*MockPacketHandler)(nil).shutdown))
}
