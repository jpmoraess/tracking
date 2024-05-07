package dto

import "github.com/jpmoraess/tracking/internal/domain/entity"

type CreateCustomerSystemInput struct {
	Name string `json:"name"`
}

func (input CreateCustomerSystemInput) NewCustomerSystemFromInput() *entity.CustomerSystem {
	return &entity.CustomerSystem{
		Name: input.Name,
	}
}
