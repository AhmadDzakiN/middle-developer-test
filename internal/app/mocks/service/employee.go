// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/contract/service.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/contract/service.go -destination=internal/app/mocks/service/employee.go
//

// Package mock_contract is a generated GoMock package.
package mock_contract

import (
	context "context"
	model "middle-developer-test/internal/app/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEmployeeService is a mock of EmployeeService interface.
type MockEmployeeService struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeServiceMockRecorder
}

// MockEmployeeServiceMockRecorder is the mock recorder for MockEmployeeService.
type MockEmployeeServiceMockRecorder struct {
	mock *MockEmployeeService
}

// NewMockEmployeeService creates a new mock instance.
func NewMockEmployeeService(ctrl *gomock.Controller) *MockEmployeeService {
	mock := &MockEmployeeService{ctrl: ctrl}
	mock.recorder = &MockEmployeeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeService) EXPECT() *MockEmployeeServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmployeeService) Create(ctx context.Context, createReq model.EmployeePayload) (model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, createReq)
	ret0, _ := ret[0].(model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEmployeeServiceMockRecorder) Create(ctx, createReq any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmployeeService)(nil).Create), ctx, createReq)
}

// DeleteByID mocks base method.
func (m *MockEmployeeService) DeleteByID(ctx context.Context, ID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockEmployeeServiceMockRecorder) DeleteByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockEmployeeService)(nil).DeleteByID), ctx, ID)
}

// GetAll mocks base method.
func (m *MockEmployeeService) GetAll(ctx context.Context) ([]model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEmployeeServiceMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEmployeeService)(nil).GetAll), ctx)
}

// GetByID mocks base method.
func (m *MockEmployeeService) GetByID(ctx context.Context, ID uint64) (model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEmployeeServiceMockRecorder) GetByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEmployeeService)(nil).GetByID), ctx, ID)
}

// Update mocks base method.
func (m *MockEmployeeService) Update(ctx context.Context, updateReq model.EmployeePayload) (model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, updateReq)
	ret0, _ := ret[0].(model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeServiceMockRecorder) Update(ctx, updateReq any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeService)(nil).Update), ctx, updateReq)
}
