package employee

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/contract"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
)

var _ contract.EmployeeRepository = new(repository)

type repository struct {
	opt options.RepositoryOption
}

func NewRepository(opt options.RepositoryOption) *repository {
	return &repository{
		opt: opt,
	}
}

func (r *repository) GetAll(ctx context.Context) (employees []model.Employee, err error) {
	result := r.opt.DbPostgre.WithContext(ctx).Find(&employees)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected < 1 {
		err = builder.WrapError(errors.New("no Employee found"), builder.ErrEmployeeNotFound)
		return
	}

	return
}

func (r *repository) GetByID(ctx context.Context, ID uint64) (employee model.Employee, err error) {
	result := r.opt.DbPostgre.WithContext(ctx).First(&employee, "id = ?", ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = builder.WrapError(gorm.ErrRecordNotFound, builder.ErrEmployeeNotFound)
			return
		}
		err = result.Error
		return
	}

	return
}

func (r *repository) Create(ctx context.Context, employee *model.Employee) (err error) {
	result := r.opt.DbPostgre.WithContext(ctx).Create(&employee)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected < 1 {
		err = builder.WrapError(errors.New("no new Employee data is created"), builder.ErrInternalServer)
		return
	}

	return
}

func (r *repository) Update(ctx context.Context, employee *model.Employee) (err error) {
	db := r.opt.DbPostgre.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true})
	result := db.Updates(&employee)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected < 1 {
		err = builder.WrapError(errors.New("no Employee data is updated"), builder.ErrEmployeeNotFound)
		return
	}

	return
}

func (r *repository) Delete(ctx context.Context, employee *model.Employee) (err error) {
	db := r.opt.DbPostgre.WithContext(ctx).Session(&gorm.Session{PrepareStmt: true})
	result := db.Delete(&employee)
	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected < 1 {
		err = builder.WrapError(errors.New("no Employee data is deleted"), builder.ErrInternalServer)
		return
	}

	return
}
