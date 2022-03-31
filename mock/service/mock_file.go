// Code generated by MockGen. DO NOT EDIT.
// Source: service/file.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileService is a mock of FileService interface.
type MockFileService struct {
	ctrl     *gomock.Controller
	recorder *MockFileServiceMockRecorder
}

// MockFileServiceMockRecorder is the mock recorder for MockFileService.
type MockFileServiceMockRecorder struct {
	mock *MockFileService
}

// NewMockFileService creates a new mock instance.
func NewMockFileService(ctrl *gomock.Controller) *MockFileService {
	mock := &MockFileService{ctrl: ctrl}
	mock.recorder = &MockFileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileService) EXPECT() *MockFileServiceMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockFileService) Generate(id, name string) (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", id, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockFileServiceMockRecorder) Generate(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockFileService)(nil).Generate), id, name)
}
