package dto

import "github.com/jpmoraess/tracking/internal/domain/entity"

type CreateShippingCompanyInput struct {
	Name string `json:"name"`
}

func (input CreateShippingCompanyInput) NewShippingCompanyFromInput() *entity.ShippingCompany {
	return &entity.ShippingCompany{
		Name: input.Name,
	}
}
