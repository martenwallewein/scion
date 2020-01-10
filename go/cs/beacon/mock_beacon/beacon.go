// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/cs/beacon (interfaces: DB,Transaction)

// Package mock_beacon is a generated GoMock package.
package mock_beacon

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	beacon "github.com/scionproto/scion/go/cs/beacon"
	addr "github.com/scionproto/scion/go/lib/addr"
	common "github.com/scionproto/scion/go/lib/common"
	path_mgmt "github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
)

// MockDB is a mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// AllRevocations mocks base method
func (m *MockDB) AllRevocations(arg0 context.Context) (<-chan beacon.RevocationOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRevocations", arg0)
	ret0, _ := ret[0].(<-chan beacon.RevocationOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRevocations indicates an expected call of AllRevocations
func (mr *MockDBMockRecorder) AllRevocations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRevocations", reflect.TypeOf((*MockDB)(nil).AllRevocations), arg0)
}

// BeaconSources mocks base method
func (m *MockDB) BeaconSources(arg0 context.Context) ([]addr.IA, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeaconSources", arg0)
	ret0, _ := ret[0].([]addr.IA)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeaconSources indicates an expected call of BeaconSources
func (mr *MockDBMockRecorder) BeaconSources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeaconSources", reflect.TypeOf((*MockDB)(nil).BeaconSources), arg0)
}

// BeginTransaction mocks base method
func (m *MockDB) BeginTransaction(arg0 context.Context, arg1 *sql.TxOptions) (beacon.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", arg0, arg1)
	ret0, _ := ret[0].(beacon.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction
func (mr *MockDBMockRecorder) BeginTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockDB)(nil).BeginTransaction), arg0, arg1)
}

// CandidateBeacons mocks base method
func (m *MockDB) CandidateBeacons(arg0 context.Context, arg1 int, arg2 beacon.Usage, arg3 addr.IA) (<-chan beacon.BeaconOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CandidateBeacons", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan beacon.BeaconOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CandidateBeacons indicates an expected call of CandidateBeacons
func (mr *MockDBMockRecorder) CandidateBeacons(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CandidateBeacons", reflect.TypeOf((*MockDB)(nil).CandidateBeacons), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (m *MockDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// DeleteExpiredBeacons mocks base method
func (m *MockDB) DeleteExpiredBeacons(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredBeacons", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredBeacons indicates an expected call of DeleteExpiredBeacons
func (mr *MockDBMockRecorder) DeleteExpiredBeacons(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredBeacons", reflect.TypeOf((*MockDB)(nil).DeleteExpiredBeacons), arg0, arg1)
}

// DeleteExpiredRevocations mocks base method
func (m *MockDB) DeleteExpiredRevocations(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredRevocations", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredRevocations indicates an expected call of DeleteExpiredRevocations
func (mr *MockDBMockRecorder) DeleteExpiredRevocations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredRevocations", reflect.TypeOf((*MockDB)(nil).DeleteExpiredRevocations), arg0, arg1)
}

// DeleteRevocation mocks base method
func (m *MockDB) DeleteRevocation(arg0 context.Context, arg1 addr.IA, arg2 common.IFIDType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRevocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRevocation indicates an expected call of DeleteRevocation
func (mr *MockDBMockRecorder) DeleteRevocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRevocation", reflect.TypeOf((*MockDB)(nil).DeleteRevocation), arg0, arg1, arg2)
}

// DeleteRevokedBeacons mocks base method
func (m *MockDB) DeleteRevokedBeacons(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRevokedBeacons", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRevokedBeacons indicates an expected call of DeleteRevokedBeacons
func (mr *MockDBMockRecorder) DeleteRevokedBeacons(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRevokedBeacons", reflect.TypeOf((*MockDB)(nil).DeleteRevokedBeacons), arg0, arg1)
}

// InsertBeacon mocks base method
func (m *MockDB) InsertBeacon(arg0 context.Context, arg1 beacon.Beacon, arg2 beacon.Usage) (beacon.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBeacon", arg0, arg1, arg2)
	ret0, _ := ret[0].(beacon.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBeacon indicates an expected call of InsertBeacon
func (mr *MockDBMockRecorder) InsertBeacon(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBeacon", reflect.TypeOf((*MockDB)(nil).InsertBeacon), arg0, arg1, arg2)
}

// InsertRevocation mocks base method
func (m *MockDB) InsertRevocation(arg0 context.Context, arg1 *path_mgmt.SignedRevInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRevocation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRevocation indicates an expected call of InsertRevocation
func (mr *MockDBMockRecorder) InsertRevocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRevocation", reflect.TypeOf((*MockDB)(nil).InsertRevocation), arg0, arg1)
}

// SetMaxIdleConns mocks base method
func (m *MockDB) SetMaxIdleConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxIdleConns", arg0)
}

// SetMaxIdleConns indicates an expected call of SetMaxIdleConns
func (mr *MockDBMockRecorder) SetMaxIdleConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxIdleConns", reflect.TypeOf((*MockDB)(nil).SetMaxIdleConns), arg0)
}

// SetMaxOpenConns mocks base method
func (m *MockDB) SetMaxOpenConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxOpenConns", arg0)
}

// SetMaxOpenConns indicates an expected call of SetMaxOpenConns
func (mr *MockDBMockRecorder) SetMaxOpenConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxOpenConns", reflect.TypeOf((*MockDB)(nil).SetMaxOpenConns), arg0)
}

// MockTransaction is a mock of Transaction interface
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// AllRevocations mocks base method
func (m *MockTransaction) AllRevocations(arg0 context.Context) (<-chan beacon.RevocationOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRevocations", arg0)
	ret0, _ := ret[0].(<-chan beacon.RevocationOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRevocations indicates an expected call of AllRevocations
func (mr *MockTransactionMockRecorder) AllRevocations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRevocations", reflect.TypeOf((*MockTransaction)(nil).AllRevocations), arg0)
}

// BeaconSources mocks base method
func (m *MockTransaction) BeaconSources(arg0 context.Context) ([]addr.IA, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeaconSources", arg0)
	ret0, _ := ret[0].([]addr.IA)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeaconSources indicates an expected call of BeaconSources
func (mr *MockTransactionMockRecorder) BeaconSources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeaconSources", reflect.TypeOf((*MockTransaction)(nil).BeaconSources), arg0)
}

// CandidateBeacons mocks base method
func (m *MockTransaction) CandidateBeacons(arg0 context.Context, arg1 int, arg2 beacon.Usage, arg3 addr.IA) (<-chan beacon.BeaconOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CandidateBeacons", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan beacon.BeaconOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CandidateBeacons indicates an expected call of CandidateBeacons
func (mr *MockTransactionMockRecorder) CandidateBeacons(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CandidateBeacons", reflect.TypeOf((*MockTransaction)(nil).CandidateBeacons), arg0, arg1, arg2, arg3)
}

// Commit mocks base method
func (m *MockTransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTransactionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit))
}

// DeleteExpiredBeacons mocks base method
func (m *MockTransaction) DeleteExpiredBeacons(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredBeacons", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredBeacons indicates an expected call of DeleteExpiredBeacons
func (mr *MockTransactionMockRecorder) DeleteExpiredBeacons(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredBeacons", reflect.TypeOf((*MockTransaction)(nil).DeleteExpiredBeacons), arg0, arg1)
}

// DeleteExpiredRevocations mocks base method
func (m *MockTransaction) DeleteExpiredRevocations(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredRevocations", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredRevocations indicates an expected call of DeleteExpiredRevocations
func (mr *MockTransactionMockRecorder) DeleteExpiredRevocations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredRevocations", reflect.TypeOf((*MockTransaction)(nil).DeleteExpiredRevocations), arg0, arg1)
}

// DeleteRevocation mocks base method
func (m *MockTransaction) DeleteRevocation(arg0 context.Context, arg1 addr.IA, arg2 common.IFIDType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRevocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRevocation indicates an expected call of DeleteRevocation
func (mr *MockTransactionMockRecorder) DeleteRevocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRevocation", reflect.TypeOf((*MockTransaction)(nil).DeleteRevocation), arg0, arg1, arg2)
}

// DeleteRevokedBeacons mocks base method
func (m *MockTransaction) DeleteRevokedBeacons(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRevokedBeacons", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRevokedBeacons indicates an expected call of DeleteRevokedBeacons
func (mr *MockTransactionMockRecorder) DeleteRevokedBeacons(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRevokedBeacons", reflect.TypeOf((*MockTransaction)(nil).DeleteRevokedBeacons), arg0, arg1)
}

// InsertBeacon mocks base method
func (m *MockTransaction) InsertBeacon(arg0 context.Context, arg1 beacon.Beacon, arg2 beacon.Usage) (beacon.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBeacon", arg0, arg1, arg2)
	ret0, _ := ret[0].(beacon.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBeacon indicates an expected call of InsertBeacon
func (mr *MockTransactionMockRecorder) InsertBeacon(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBeacon", reflect.TypeOf((*MockTransaction)(nil).InsertBeacon), arg0, arg1, arg2)
}

// InsertRevocation mocks base method
func (m *MockTransaction) InsertRevocation(arg0 context.Context, arg1 *path_mgmt.SignedRevInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRevocation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRevocation indicates an expected call of InsertRevocation
func (mr *MockTransactionMockRecorder) InsertRevocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRevocation", reflect.TypeOf((*MockTransaction)(nil).InsertRevocation), arg0, arg1)
}

// Rollback mocks base method
func (m *MockTransaction) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockTransactionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback))
}