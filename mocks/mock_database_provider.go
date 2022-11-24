package mocks

import (
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
        patient "github.com/tranthienphucbreak/parsleytest/internal/patient"
)

// MockDatabaseProvider is a mock of DatabaseProvider interface.
type MockDatabaseProvider struct {
        ctrl     *gomock.Controller
        recorder *MockDatabaseProviderMockRecorder
}

// MockDatabaseProviderMockRecorder is the mock recorder for MockDatabaseProvider.
type MockDatabaseProviderMockRecorder struct {
        mock *MockDatabaseProvider
}

// NewMockDatabaseProvider creates a new mock instance.
func NewMockDatabaseProvider(ctrl *gomock.Controller) *MockDatabaseProvider {
        mock := &MockDatabaseProvider{ctrl: ctrl}
        mock.recorder = &MockDatabaseProviderMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseProvider) EXPECT() *MockDatabaseProviderMockRecorder {
        return m.recorder
}

// Exec mocks base method.
func (m *MockDatabaseProvider) Exec(query string, params []interface{}) (int64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Exec", query, params)
        ret0, _ := ret[0].(int64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDatabaseProviderMockRecorder) Exec(query, params interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDatabaseProvider)(nil).Exec), query, params)
}

// Query mocks base method.
func (m *MockDatabaseProvider) Query(query string, params []interface{}) ([]patient.Person, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Query", query, params)
        ret0, _ := ret[0].([]patient.Person)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDatabaseProviderMockRecorder) Query(query, params interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDatabaseProvider)(nil).Query), query, params)
}

// QueryRow mocks base method.
func (m *MockDatabaseProvider) QueryRow(query string, params []interface{}) (*patient.Person, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "QueryRow", query, params)
        ret0, _ := ret[0].(*patient.Person)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockDatabaseProviderMockRecorder) QueryRow(query, params interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockDatabaseProvider)(nil).QueryRow), query, params)
}