package employee

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/config"
	mockSvc "middle-developer-test/internal/app/mocks/service"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type HandlerTestSuite struct {
	suite.Suite
	handlerMock     Handler
	employeeSvcMock *mockSvc.MockEmployeeService
	loc             *time.Location
}

func (h *HandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(h.T())
	defer mockCtrl.Finish()

	h.employeeSvcMock = mockSvc.NewMockEmployeeService(mockCtrl)

	cfg := config.AppConfig()
	app := appcontext.NewAppContext(cfg)
	validator := validator.New()
	opt := options.HandlerOption{
		AppOptions: options.AppOptions{
			AppConfig: cfg,
			AppCtx:    app,
		},
		Services: &appcontext.Service{
			Employee: h.employeeSvcMock,
		},
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	h.loc = loc

	h.handlerMock = Handler{
		HandlerOption: opt,
		validator:     *validator,
	}
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (h *HandlerTestSuite) TestGetAll() {
	employeeData := []model.Employee{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john_doe@gmail.com",
			HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, h.loc),
		},
		{
			ID:        1,
			FirstName: "Mark",
			LastName:  "Jonathan",
			Email:     "mark_jonathan@gmail.com",
			HireDate:  time.Date(2018, 03, 27, 00, 00, 00, 00, h.loc),
		},
	}

	type fields struct {
		mock func()
	}

	tests := []struct {
		name               string
		fields             fields
		expectedStatusCode int
	}{
		{
			name: "Success",
			fields: fields{
				mock: func() {
					h.employeeSvcMock.EXPECT().GetAll(gomock.Any()).Return(employeeData, nil)
				}},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Failed get all Employees data from service",
			fields: fields{
				mock: func() {
					h.employeeSvcMock.EXPECT().GetAll(gomock.Any()).Return([]model.Employee{}, errors.New("random error"))
				}},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		h.Suite.Run(test.name, func() {
			ctx := chi.NewRouteContext()
			req, _ := http.NewRequest(http.MethodGet, "/employees", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			writer := httptest.NewRecorder()
			test.fields.mock()

			handler := builder.HttpHandler{H: h.handlerMock.GetAll}
			handler.ServeHTTP(writer, req)

			assert.Equal(h.T(), test.expectedStatusCode, writer.Code)
		})
	}
}

func (h *HandlerTestSuite) TestGetByID() {
	employeeData := model.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john_doe@gmail.com",
		HireDate:  time.Date(2020, 01, 03, 00, 00, 00, 00, h.loc),
	}

	type args struct {
		ID string
	}

	type fields struct {
		mock func(ID uint64)
	}

	tests := []struct {
		name               string
		args               args
		fields             fields
		expectedStatusCode int
	}{
		{
			name: "Success",
			args: args{
				ID: "1",
			},
			fields: fields{
				mock: func(ID uint64) {
					h.employeeSvcMock.EXPECT().GetByID(gomock.Any(), ID).Return(employeeData, nil)
				}},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Failed, get Employee data by ID from service",
			args: args{
				ID: "1",
			},
			fields: fields{
				mock: func(ID uint64) {
					h.employeeSvcMock.EXPECT().GetByID(gomock.Any(), ID).Return(model.Employee{}, errors.New("random error"))
				}},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Failed, invalid employee_id",
			fields: fields{
				mock: func(ID uint64) {
					h.employeeSvcMock.EXPECT().GetByID(gomock.Any(), ID).Return(model.Employee{}, errors.New("random error"))
				}},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		h.Suite.Run(test.name, func() {
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("employee_id", test.args.ID)
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/employees/%s", test.args.ID), nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			writer := httptest.NewRecorder()
			employeeID, _ := strconv.ParseUint(test.args.ID, 10, 64)
			test.fields.mock(employeeID)

			handler := builder.HttpHandler{H: h.handlerMock.GetByID}
			handler.ServeHTTP(writer, req)

			assert.Equal(h.T(), test.expectedStatusCode, writer.Code)
		})
	}
}
