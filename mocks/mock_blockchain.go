// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/juno/blockchain (interfaces: Reader)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/NethermindEth/juno/core"
	felt "github.com/NethermindEth/juno/core/felt"
	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockReader) BlockByHash(arg0 *felt.Felt) (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", arg0)
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockReaderMockRecorder) BlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockReader)(nil).BlockByHash), arg0)
}

// BlockByNumber mocks base method.
func (m *MockReader) BlockByNumber(arg0 uint64) (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", arg0)
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockReaderMockRecorder) BlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockReader)(nil).BlockByNumber), arg0)
}

// BlockHeaderByHash mocks base method.
func (m *MockReader) BlockHeaderByHash(arg0 *felt.Felt) (*core.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHeaderByHash", arg0)
	ret0, _ := ret[0].(*core.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHeaderByHash indicates an expected call of BlockHeaderByHash.
func (mr *MockReaderMockRecorder) BlockHeaderByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHeaderByHash", reflect.TypeOf((*MockReader)(nil).BlockHeaderByHash), arg0)
}

// BlockHeaderByNumber mocks base method.
func (m *MockReader) BlockHeaderByNumber(arg0 uint64) (*core.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHeaderByNumber", arg0)
	ret0, _ := ret[0].(*core.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHeaderByNumber indicates an expected call of BlockHeaderByNumber.
func (mr *MockReaderMockRecorder) BlockHeaderByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHeaderByNumber", reflect.TypeOf((*MockReader)(nil).BlockHeaderByNumber), arg0)
}

// Head mocks base method.
func (m *MockReader) Head() (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockReaderMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockReader)(nil).Head))
}

// HeadState mocks base method.
func (m *MockReader) HeadState() (core.StateReader, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadState")
	ret0, _ := ret[0].(core.StateReader)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// HeadState indicates an expected call of HeadState.
func (mr *MockReaderMockRecorder) HeadState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadState", reflect.TypeOf((*MockReader)(nil).HeadState))
}

// HeadsHeader mocks base method.
func (m *MockReader) HeadsHeader() (*core.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadsHeader")
	ret0, _ := ret[0].(*core.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadsHeader indicates an expected call of HeadsHeader.
func (mr *MockReaderMockRecorder) HeadsHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadsHeader", reflect.TypeOf((*MockReader)(nil).HeadsHeader))
}

// Height mocks base method.
func (m *MockReader) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height.
func (mr *MockReaderMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockReader)(nil).Height))
}

// Receipt mocks base method.
func (m *MockReader) Receipt(arg0 *felt.Felt) (*core.TransactionReceipt, *felt.Felt, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receipt", arg0)
	ret0, _ := ret[0].(*core.TransactionReceipt)
	ret1, _ := ret[1].(*felt.Felt)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Receipt indicates an expected call of Receipt.
func (mr *MockReaderMockRecorder) Receipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receipt", reflect.TypeOf((*MockReader)(nil).Receipt), arg0)
}

// StateAtBlockHash mocks base method.
func (m *MockReader) StateAtBlockHash(arg0 *felt.Felt) (core.StateReader, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtBlockHash", arg0)
	ret0, _ := ret[0].(core.StateReader)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StateAtBlockHash indicates an expected call of StateAtBlockHash.
func (mr *MockReaderMockRecorder) StateAtBlockHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtBlockHash", reflect.TypeOf((*MockReader)(nil).StateAtBlockHash), arg0)
}

// StateAtBlockNumber mocks base method.
func (m *MockReader) StateAtBlockNumber(arg0 uint64) (core.StateReader, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtBlockNumber", arg0)
	ret0, _ := ret[0].(core.StateReader)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StateAtBlockNumber indicates an expected call of StateAtBlockNumber.
func (mr *MockReaderMockRecorder) StateAtBlockNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtBlockNumber", reflect.TypeOf((*MockReader)(nil).StateAtBlockNumber), arg0)
}

// StateUpdateByHash mocks base method.
func (m *MockReader) StateUpdateByHash(arg0 *felt.Felt) (*core.StateUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateUpdateByHash", arg0)
	ret0, _ := ret[0].(*core.StateUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateUpdateByHash indicates an expected call of StateUpdateByHash.
func (mr *MockReaderMockRecorder) StateUpdateByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateUpdateByHash", reflect.TypeOf((*MockReader)(nil).StateUpdateByHash), arg0)
}

// StateUpdateByNumber mocks base method.
func (m *MockReader) StateUpdateByNumber(arg0 uint64) (*core.StateUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateUpdateByNumber", arg0)
	ret0, _ := ret[0].(*core.StateUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateUpdateByNumber indicates an expected call of StateUpdateByNumber.
func (mr *MockReaderMockRecorder) StateUpdateByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateUpdateByNumber", reflect.TypeOf((*MockReader)(nil).StateUpdateByNumber), arg0)
}

// TransactionByBlockNumberAndIndex mocks base method.
func (m *MockReader) TransactionByBlockNumberAndIndex(arg0, arg1 uint64) (core.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByBlockNumberAndIndex", arg0, arg1)
	ret0, _ := ret[0].(core.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByBlockNumberAndIndex indicates an expected call of TransactionByBlockNumberAndIndex.
func (mr *MockReaderMockRecorder) TransactionByBlockNumberAndIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByBlockNumberAndIndex", reflect.TypeOf((*MockReader)(nil).TransactionByBlockNumberAndIndex), arg0, arg1)
}

// TransactionByHash mocks base method.
func (m *MockReader) TransactionByHash(arg0 *felt.Felt) (core.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", arg0)
	ret0, _ := ret[0].(core.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockReaderMockRecorder) TransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockReader)(nil).TransactionByHash), arg0)
}
