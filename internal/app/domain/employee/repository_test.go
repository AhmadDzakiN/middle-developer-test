package employee

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/config"
	"middle-developer-test/internal/app/contract"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
	"testing"
	"time"
)

type RepositoryTestSuite struct {
	suite.Suite
	sqlMock        sqlmock.Sqlmock
	repositoryMock contract.EmployeeRepository
	loc            *time.Location
	ctx            context.Context
}

func (r *RepositoryTestSuite) SetupTest() {
	sqlDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, DriverName: "postgres", WithoutQuotingCheck: true}), &gorm.Config{})

	cfg := config.AppConfig()
	app := appcontext.NewAppContext(cfg)
	opt := options.RepositoryOption{
		AppOptions: options.AppOptions{
			AppConfig: cfg,
			AppCtx:    app,
			DbPostgre: gormdb,
		}}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	r.loc = loc

	r.ctx = context.Background()
	r.repositoryMock = NewRepository(opt)
	r.sqlMock = mock
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (r *RepositoryTestSuite) TestGetAll() {
	type fields struct {
		mock func()
	}

	type args struct {
		ctx context.Context
	}

	query := "SELECT * FROM employees"

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
				ctx: r.ctx,
			},
			fields: fields{
				mock: func() {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).
						WillReturnRows(r.sqlMock.NewRows([]string{"id", "first_name", "last_name", "email", "hire_date"}).
							AddRow(1, "John", "Doe", "john_doe@example.com", time.Date(2020, 01, 03, 00, 00, 00, 00, r.loc)).
							AddRow(2, "Mark", "Jonathan", "mark_jonathan@gmail.com", time.Date(2018, 03, 27, 00, 00, 00, 00, r.loc)))
				}},
			expectedResult: []model.Employee{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john_doe@example.com",
					HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, r.loc),
				},
				{
					ID:        2,
					FirstName: "Mark",
					LastName:  "Jonathan",
					Email:     "mark_jonathan@gmail.com",
					HireDate:  time.Date(2018, 03, 27, 00, 00, 00, 00, r.loc),
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: r.ctx,
			},
			fields: fields{
				mock: func() {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: []model.Employee(nil),
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
		{
			name: "Failed, no Employee data returned",
			args: args{
				ctx: r.ctx,
			},
			fields: fields{
				mock: func() {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).WillReturnRows(r.sqlMock.NewRows([]string{"id", "first_name", "last_name", "email", "hire_date"}))
				}},
			expectedResult: []model.Employee{},
			expectedErr:    builder.WrapError(errors.New("no Employee found"), builder.ErrEmployeeNotFound),
		},
	}

	for _, test := range tests {
		r.Suite.Run(test.name, func() {
			test.fields.mock()

			actualResult, actualErr := r.repositoryMock.GetAll(test.args.ctx)

			assert.Equal(r.T(), test.expectedErr, actualErr, test.name)
			assert.Equal(r.T(), test.expectedResult, actualResult, test.name)
		})
	}

}

func (r *RepositoryTestSuite) TestGetByID() {
	type fields struct {
		mock func(employeeID uint64)
	}

	type args struct {
		ctx context.Context
		ID  uint64
	}

	query := `SELECT * FROM employees WHERE id = $1 ORDER BY employees.id LIMIT 1`

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
				ctx: r.ctx,
				ID:  uint64(1),
			},
			fields: fields{
				mock: func(ID uint64) {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).WithArgs(ID).
						WillReturnRows(r.sqlMock.NewRows([]string{"id", "first_name", "last_name", "email", "hire_date"}).
							AddRow(1, "John", "Doe", "john_doe@example.com", time.Date(2020, 01, 03, 00, 00, 00, 00, r.loc)))
				}},
			expectedResult: model.Employee{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john_doe@example.com",
				HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, r.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, Employee data with associated ID is not found",
			args: args{
				ctx: r.ctx,
				ID:  uint64(1),
			},
			fields: fields{
				mock: func(ID uint64) {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).WithArgs(ID).WillReturnError(gorm.ErrRecordNotFound)
				}},
			expectedResult: model.Employee{},
			expectedErr:    builder.WrapError(gorm.ErrRecordNotFound, builder.ErrEmployeeNotFound),
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: r.ctx,
				ID:  uint64(1),
			},
			fields: fields{
				mock: func(ID uint64) {
					r.sqlMock.MatchExpectationsInOrder(false)
					r.sqlMock.ExpectQuery(query).WithArgs(ID).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: model.Employee{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		r.Suite.Run(test.name, func() {
			test.fields.mock(test.args.ID)

			actualResult, actualErr := r.repositoryMock.GetByID(test.args.ctx, test.args.ID)

			assert.Equal(r.T(), test.expectedErr, actualErr, test.name)
			assert.Equal(r.T(), test.expectedResult, actualResult, test.name)
		})
	}

}
