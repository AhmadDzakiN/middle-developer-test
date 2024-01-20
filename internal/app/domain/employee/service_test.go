package employee

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/config"
	"middle-developer-test/internal/app/contract"
	mockRepo "middle-developer-test/internal/app/mocks/repository"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
	serviceMock      contract.EmployeeService
	employeeRepoMock *mockRepo.MockEmployeeRepository
	loc              *time.Location
	ctx              context.Context
}

func (s *ServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	s.employeeRepoMock = mockRepo.NewMockEmployeeRepository(mockCtrl)

	cfg := config.AppConfig()
	app := appcontext.NewAppContext(cfg)
	opt := options.ServiceOption{
		AppOptions: options.AppOptions{
			AppConfig: cfg,
			AppCtx:    app,
		},
		Repository: &appcontext.Repository{
			Employee: s.employeeRepoMock,
		},
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	s.loc = loc

	s.ctx = context.Background()
	s.serviceMock = NewService(opt)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestGetAll() {
	employeeData := []model.Employee{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john_doe@gmail.com",
			HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
		},
		{
			ID:        2,
			FirstName: "Mark",
			LastName:  "Jonathan",
			Email:     "mark_jonathan@gmail.com",
			HireDate:  time.Date(2018, 03, 27, 00, 00, 00, 00, s.loc),
		},
	}

	type fields struct {
		mock func(ctx context.Context)
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult []model.Employee
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: s.ctx,
			},
			fields: fields{
				mock: func(ctx context.Context) {
					s.employeeRepoMock.EXPECT().GetAll(ctx).Return(employeeData, nil)
				},
			},
			expectedResult: employeeData,
			expectedErr:    nil,
		},
		{
			name: "Failed get all Employees data from repo",
			args: args{
				ctx: s.ctx,
			},
			fields: fields{
				mock: func(ctx context.Context) {
					s.employeeRepoMock.EXPECT().GetAll(ctx).Return([]model.Employee{}, errors.New("random error"))
				},
			},
			expectedResult: []model.Employee{},
			expectedErr:    errors.New("random error"),
		},
	}

	for _, test := range tests {
		s.Suite.Run(test.name, func() {
			test.fields.mock(test.args.ctx)

			actualResult, actualErr := s.serviceMock.GetAll(test.args.ctx)

			assert.Equal(s.T(), test.expectedErr, actualErr, test.name)
			assert.Equal(s.T(), test.expectedResult, actualResult, test.name)
		})
	}
}

func (s *ServiceTestSuite) TestGetByID() {
	employeeData := model.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john_doe@gmail.com",
		HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
	}

	type fields struct {
		mock func(ctx context.Context, ID uint64)
	}

	type args struct {
		ctx context.Context
		ID  uint64
	}

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult model.Employee
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: s.ctx,
				ID:  uint64(1),
			},
			fields: fields{
				mock: func(ctx context.Context, ID uint64) {
					s.employeeRepoMock.EXPECT().GetByID(ctx, ID).Return(employeeData, nil)
				},
			},
			expectedResult: employeeData,
			expectedErr:    nil,
		},
		{
			name: "Failed get Employee data by ID from repo",
			args: args{
				ctx: s.ctx,
			},
			fields: fields{
				mock: func(ctx context.Context, ID uint64) {
					s.employeeRepoMock.EXPECT().GetByID(ctx, ID).Return(model.Employee{}, errors.New("random error"))
				},
			},
			expectedResult: model.Employee{},
			expectedErr:    errors.New("random error"),
		},
	}

	for _, test := range tests {
		s.Suite.Run(test.name, func() {
			test.fields.mock(test.args.ctx, test.args.ID)

			actualResult, actualErr := s.serviceMock.GetByID(test.args.ctx, test.args.ID)

			assert.Equal(s.T(), test.expectedErr, actualErr, test.name)
			assert.Equal(s.T(), test.expectedResult, actualResult, test.name)
		})
	}
}
