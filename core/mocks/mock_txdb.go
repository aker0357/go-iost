// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/Go-IOS-Protocol/core/tx (interfaces: TxDB)

// Package core_mock is a generated GoMock package.
package core_mock

import (
	gomock "github.com/golang/mock/gomock"
	tx "github.com/iost-official/Go-IOS-Protocol/core/tx"
	reflect "reflect"
)

// MockTxDB is a mock of TxDB interface
type MockTxDB struct {
	ctrl     *gomock.Controller
	recorder *MockTxDBMockRecorder
}

// MockTxDBMockRecorder is the mock recorder for MockTxDB
type MockTxDBMockRecorder struct {
	mock *MockTxDB
}

// NewMockTxDB creates a new mock instance
func NewMockTxDB(ctrl *gomock.Controller) *MockTxDB {
	mock := &MockTxDB{ctrl: ctrl}
	mock.recorder = &MockTxDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxDB) EXPECT() *MockTxDBMockRecorder {
	return m.recorder
}

// GetReceipt mocks base method
func (m *MockTxDB) GetReceipt(arg0 []byte) (*tx.TxReceipt, error) {
	ret := m.ctrl.Call(m, "GetReceipt", arg0)
	ret0, _ := ret[0].(*tx.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt
func (mr *MockTxDBMockRecorder) GetReceipt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockTxDB)(nil).GetReceipt), arg0)
}

// GetReceiptByTxHash mocks base method
func (m *MockTxDB) GetReceiptByTxHash(arg0 []byte) (*tx.TxReceipt, error) {
	ret := m.ctrl.Call(m, "GetReceiptByTxHash", arg0)
	ret0, _ := ret[0].(*tx.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptByTxHash indicates an expected call of GetReceiptByTxHash
func (mr *MockTxDBMockRecorder) GetReceiptByTxHash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptByTxHash", reflect.TypeOf((*MockTxDB)(nil).GetReceiptByTxHash), arg0)
}

// GetTx mocks base method
func (m *MockTxDB) GetTx(arg0 []byte) (*tx.Tx, error) {
	ret := m.ctrl.Call(m, "GetTx", arg0)
	ret0, _ := ret[0].(*tx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx
func (mr *MockTxDBMockRecorder) GetTx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockTxDB)(nil).GetTx), arg0)
}

// HasReceipt mocks base method
func (m *MockTxDB) HasReceipt(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "HasReceipt", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasReceipt indicates an expected call of HasReceipt
func (mr *MockTxDBMockRecorder) HasReceipt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasReceipt", reflect.TypeOf((*MockTxDB)(nil).HasReceipt), arg0)
}

// HasTx mocks base method
func (m *MockTxDB) HasTx(arg0 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "HasTx", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasTx indicates an expected call of HasTx
func (mr *MockTxDBMockRecorder) HasTx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTx", reflect.TypeOf((*MockTxDB)(nil).HasTx), arg0)
}

// Push mocks base method
func (m *MockTxDB) Push(arg0 []*tx.Tx, arg1 []*tx.TxReceipt) error {
	ret := m.ctrl.Call(m, "Push", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockTxDBMockRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockTxDB)(nil).Push), arg0, arg1)
}
