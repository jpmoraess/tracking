package usecase

import (
	"context"

	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/repository"
)

type CreateShippingCompany struct {
	repository repository.ShippingCompanyRepository
}

func NewCreateShippingCompany(repository repository.ShippingCompanyRepository) *CreateShippingCompany {
	return &CreateShippingCompany{
		repository: repository,
	}
}

func (c *CreateShippingCompany) Execute(input dto.CreateShippingCompanyInput) (string, error) {
	entity := input.NewShippingCompanyFromInput()
	saved, err := c.repository.Insert(context.Background(), entity)
	if err != nil {
		return "", err
	}
	return saved.ID, nil
}
