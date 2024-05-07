package dto

import "github.com/jpmoraess/tracking/internal/domain/entity"

type CreateCustomerInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (input CreateCustomerInput) NewCustomerFromInput() *entity.Customer {
	return &entity.Customer{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
}
