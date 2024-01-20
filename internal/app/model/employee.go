package model

import "time"

type Employee struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	HireDate  time.Time `json:"hire_date"`
}

type EmployeePayload struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name" validate:"required,alphanum"`
	LastName  string `json:"last_name" validate:"required,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	HireDate  string `json:"hire_date" validate:"required"`
}
