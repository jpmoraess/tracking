package usecase

import (
	"context"

	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/repository"
)

type CreateCustomer struct {
	repository repository.CustomerRepository
}

func NewCreateCustomer(repository repository.CustomerRepository) *CreateCustomer {
	return &CreateCustomer{
		repository: repository,
	}
}

func (c *CreateCustomer) Execute(input dto.CreateCustomerInput) (string, error) {
	entity := input.NewCustomerFromInput()
	saved, err := c.repository.Insert(context.Background(), entity)
	if err != nil {
		return "", err
	}
	return saved.ID, nil
}
