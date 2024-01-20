package employee

import (
	"context"
	"github.com/rs/zerolog/log"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/contract"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
	"time"
)

var _ contract.EmployeeService = new(service)

type service struct {
	opt options.ServiceOption
}

func NewService(opt options.ServiceOption) *service {
	return &service{
		opt: opt,
	}
}

func (s *service) GetAll(ctx context.Context) (employees []model.Employee, err error) {
	employees, err = s.opt.Repository.Employee.GetAll(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to get all Employees data")
		return
	}

	return
}

func (s *service) GetByID(ctx context.Context, ID uint64) (employee model.Employee, err error) {
	employee, err = s.opt.Repository.Employee.GetByID(ctx, ID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Employee data by ID: %d", ID)
		return
	}

	return
}

func (s *service) Create(ctx context.Context, createReq model.EmployeePayload) (createdEmployee model.Employee, err error) {
	hireDate, err := time.Parse(time.DateOnly, createReq.HireDate)
	if err != nil {
		log.Err(err).Msg("Failed to parse hire_date into time")
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	newEmployee := model.Employee{
		FirstName: createReq.FirstName,
		LastName:  createReq.LastName,
		Email:     createReq.Email,
		HireDate:  hireDate,
	}

	err = s.opt.Repository.Employee.Create(ctx, &newEmployee)
	if err != nil {
		log.Err(err).Msgf("Failed to create new Employee data for email: %s", createReq.Email)
		return
	}

	createdEmployee = newEmployee

	return
}

func (s *service) Update(ctx context.Context, updateReq model.EmployeePayload) (updatedEmployee model.Employee, err error) {
	employee, err := s.opt.Repository.Employee.GetByID(ctx, updateReq.ID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Employee data by ID: %d", updateReq.ID)
		return
	}

	hireDate, err := time.Parse(time.DateOnly, updateReq.HireDate)
	if err != nil {
		log.Err(err).Msg("Failed to parse hire_date into time")
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	updateEmployee := model.Employee{
		ID:        employee.ID,
		FirstName: updateReq.FirstName,
		LastName:  updateReq.LastName,
		Email:     updateReq.Email,
		HireDate:  hireDate,
	}

	err = s.opt.Repository.Employee.Update(ctx, &updateEmployee)
	if err != nil {
		log.Err(err).Msgf("Failed to update Employee data by ID: %d", updateReq.ID)
		return
	}

	updatedEmployee = updateEmployee

	return
}

func (s *service) DeleteByID(ctx context.Context, ID uint64) (err error) {
	employee, err := s.opt.Repository.Employee.GetByID(ctx, ID)
	if err != nil {
		log.Err(err).Msgf("Failed to get Employee data by ID: %d", ID)
		return
	}

	err = s.opt.Repository.Employee.Delete(ctx, &employee)
	if err != nil {
		log.Err(err).Msgf("Failed to delete Employee data by ID: %d", ID)
		return
	}

	return
}
