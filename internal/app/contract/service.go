package contract

import (
	"context"
	"middle-developer-test/internal/app/model"
)

type EmployeeService interface {
	GetAll(ctx context.Context) (employees []model.Employee, err error)
	GetByID(ctx context.Context, ID uint64) (employee model.Employee, err error)
	Create(ctx context.Context, createReq model.EmployeePayload) (createdEmployee model.Employee, err error)
	Update(ctx context.Context, updateReq model.EmployeePayload) (updatedEmployee model.Employee, err error)
	DeleteByID(ctx context.Context, ID uint64) (err error)
}
