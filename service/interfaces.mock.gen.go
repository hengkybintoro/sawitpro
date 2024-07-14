// Code generated by MockGen. DO NOT EDIT.
// Source: service/interfaces.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/SawitProRecruitment/UserService/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// AddEstate mocks base method.
func (m *MockServiceInterface) AddEstate(width, length int) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEstate", width, length)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEstate indicates an expected call of AddEstate.
func (mr *MockServiceInterfaceMockRecorder) AddEstate(width, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEstate", reflect.TypeOf((*MockServiceInterface)(nil).AddEstate), width, length)
}

// AddTree mocks base method.
func (m *MockServiceInterface) AddTree(estateID uuid.UUID, x, y, height int) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTree", estateID, x, y, height)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTree indicates an expected call of AddTree.
func (mr *MockServiceInterfaceMockRecorder) AddTree(estateID, x, y, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTree", reflect.TypeOf((*MockServiceInterface)(nil).AddTree), estateID, x, y, height)
}

// GetDronePlanDistance mocks base method.
func (m *MockServiceInterface) GetDronePlanDistance(estateID uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDronePlanDistance", estateID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDronePlanDistance indicates an expected call of GetDronePlanDistance.
func (mr *MockServiceInterfaceMockRecorder) GetDronePlanDistance(estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDronePlanDistance", reflect.TypeOf((*MockServiceInterface)(nil).GetDronePlanDistance), estateID)
}

// GetDronePlanMaxDistance mocks base method.
func (m *MockServiceInterface) GetDronePlanMaxDistance(estateID uuid.UUID, maxDistance int) (int, model.Coordinate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDronePlanMaxDistance", estateID, maxDistance)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(model.Coordinate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDronePlanMaxDistance indicates an expected call of GetDronePlanMaxDistance.
func (mr *MockServiceInterfaceMockRecorder) GetDronePlanMaxDistance(estateID, maxDistance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDronePlanMaxDistance", reflect.TypeOf((*MockServiceInterface)(nil).GetDronePlanMaxDistance), estateID, maxDistance)
}

// GetEstateStats mocks base method.
func (m *MockServiceInterface) GetEstateStats(estateID uuid.UUID) (int, int, int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateStats", estateID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(int)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetEstateStats indicates an expected call of GetEstateStats.
func (mr *MockServiceInterfaceMockRecorder) GetEstateStats(estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateStats", reflect.TypeOf((*MockServiceInterface)(nil).GetEstateStats), estateID)
}
