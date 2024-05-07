package usecase

import (
	"context"

	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/repository"
)

type CreateCustomerSystem struct {
	repository repository.CustomerSystemRepository
}

func NewCreateCustomerSystem(repository repository.CustomerSystemRepository) *CreateCustomerSystem {
	return &CreateCustomerSystem{
		repository: repository,
	}
}

func (c *CreateCustomerSystem) Execute(input dto.CreateCustomerSystemInput) (string, error) {
	entity := input.NewCustomerSystemFromInput()
	saved, err := c.repository.Insert(context.Background(), entity)
	if err != nil {
		return "", err
	}
	return saved.ID, nil
}
