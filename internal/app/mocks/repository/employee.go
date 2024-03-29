// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/contract/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/contract/repository.go -destination=internal/app/mocks/repository/employee.go
//

// Package mock_contract is a generated GoMock package.
package mock_contract

import (
	context "context"
	model "middle-developer-test/internal/app/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEmployeeRepository is a mock of EmployeeRepository interface.
type MockEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryMockRecorder
}

// MockEmployeeRepositoryMockRecorder is the mock recorder for MockEmployeeRepository.
type MockEmployeeRepositoryMockRecorder struct {
	mock *MockEmployeeRepository
}

// NewMockEmployeeRepository creates a new mock instance.
func NewMockEmployeeRepository(ctrl *gomock.Controller) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepository) EXPECT() *MockEmployeeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmployeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEmployeeRepositoryMockRecorder) Create(ctx, employee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmployeeRepository)(nil).Create), ctx, employee)
}

// Delete mocks base method.
func (m *MockEmployeeRepository) Delete(ctx context.Context, employee *model.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeRepositoryMockRecorder) Delete(ctx, employee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeRepository)(nil).Delete), ctx, employee)
}

// GetAll mocks base method.
func (m *MockEmployeeRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEmployeeRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEmployeeRepository)(nil).GetAll), ctx)
}

// GetByID mocks base method.
func (m *MockEmployeeRepository) GetByID(ctx context.Context, ID uint64) (model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockEmployeeRepositoryMockRecorder) GetByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEmployeeRepository)(nil).GetByID), ctx, ID)
}

// Update mocks base method.
func (m *MockEmployeeRepository) Update(ctx context.Context, employee *model.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeRepositoryMockRecorder) Update(ctx, employee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeRepository)(nil).Update), ctx, employee)
}
