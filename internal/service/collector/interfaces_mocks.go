// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package service is a generated GoMock package.
package collector

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/jaroslav1991/tts/internal/model"
)

// MockDataReader is a mock of DataReader interface.
type MockDataReader struct {
	ctrl     *gomock.Controller
	recorder *MockDataReaderMockRecorder
}

// MockDataReaderMockRecorder is the mock recorder for MockDataReader.
type MockDataReaderMockRecorder struct {
	mock *MockDataReader
}

// NewMockDataReader creates a new mock instance.
func NewMockDataReader(ctrl *gomock.Controller) *MockDataReader {
	mock := &MockDataReader{ctrl: ctrl}
	mock.recorder = &MockDataReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataReader) EXPECT() *MockDataReaderMockRecorder {
	return m.recorder
}

// ReadData mocks base method.
func (m *MockDataReader) ReadData(request any) (model.PluginInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadData", request)
	ret0, _ := ret[0].(model.PluginInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadData indicates an expected call of ReadData.
func (mr *MockDataReaderMockRecorder) ReadData(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadData", reflect.TypeOf((*MockDataReader)(nil).ReadData), request)
}

// MockDataValidator is a mock of DataValidator interface.
type MockDataValidator struct {
	ctrl     *gomock.Controller
	recorder *MockDataValidatorMockRecorder
}

// MockDataValidatorMockRecorder is the mock recorder for MockDataValidator.
type MockDataValidatorMockRecorder struct {
	mock *MockDataValidator
}

// NewMockDataValidator creates a new mock instance.
func NewMockDataValidator(ctrl *gomock.Controller) *MockDataValidator {
	mock := &MockDataValidator{ctrl: ctrl}
	mock.recorder = &MockDataValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataValidator) EXPECT() *MockDataValidatorMockRecorder {
	return m.recorder
}

// ValidateData mocks base method.
func (m *MockDataValidator) ValidateData(data model.PluginInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateData", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateData indicates an expected call of ValidateData.
func (mr *MockDataValidatorMockRecorder) ValidateData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateData", reflect.TypeOf((*MockDataValidator)(nil).ValidateData), data)
}

// MockDataPreparer is a mock of DataPreparer interface.
type MockDataPreparer struct {
	ctrl     *gomock.Controller
	recorder *MockDataPreparerMockRecorder
}

// MockDataPreparerMockRecorder is the mock recorder for MockDataPreparer.
type MockDataPreparerMockRecorder struct {
	mock *MockDataPreparer
}

// NewMockDataPreparer creates a new mock instance.
func NewMockDataPreparer(ctrl *gomock.Controller) *MockDataPreparer {
	mock := &MockDataPreparer{ctrl: ctrl}
	mock.recorder = &MockDataPreparerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataPreparer) EXPECT() *MockDataPreparerMockRecorder {
	return m.recorder
}

// PrepareData mocks base method.
func (m *MockDataPreparer) PrepareData(data model.PluginInfo) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareData", data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareData indicates an expected call of PrepareData.
func (mr *MockDataPreparerMockRecorder) PrepareData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareData", reflect.TypeOf((*MockDataPreparer)(nil).PrepareData), data)
}

// MockDataSaver is a mock of DataSaver interface.
type MockDataSaver struct {
	ctrl     *gomock.Controller
	recorder *MockDataSaverMockRecorder
}

// MockDataSaverMockRecorder is the mock recorder for MockDataSaver.
type MockDataSaverMockRecorder struct {
	mock *MockDataSaver
}

// NewMockDataSaver creates a new mock instance.
func NewMockDataSaver(ctrl *gomock.Controller) *MockDataSaver {
	mock := &MockDataSaver{ctrl: ctrl}
	mock.recorder = &MockDataSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSaver) EXPECT() *MockDataSaverMockRecorder {
	return m.recorder
}

// SaveData mocks base method.
func (m *MockDataSaver) SaveData(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveData indicates an expected call of SaveData.
func (mr *MockDataSaverMockRecorder) SaveData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveData", reflect.TypeOf((*MockDataSaver)(nil).SaveData), arg0)
}
