package contract

import (
	"context"
	"middle-developer-test/internal/app/model"
)

type EmployeeRepository interface {
	GetAll(ctx context.Context) (employees []model.Employee, err error)
	GetByID(ctx context.Context, ID uint64) (employee model.Employee, err error)
	Create(ctx context.Context, employee *model.Employee) (err error)
	Update(ctx context.Context, employee *model.Employee) (err error)
	Delete(ctx context.Context, employee *model.Employee) (err error)
}
