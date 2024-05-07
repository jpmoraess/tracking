package dto

import "github.com/jpmoraess/tracking/internal/domain/entity"

type CreateCustomerIntegrationInput struct {
	CustomerId            string `json:"customerId"`
	CustomerSystemId      string `json:"customerSystemId"`
	ShippingCompanyId     string `json:"shippingCompanyId"`
	CustomerSystemApiKey  string `json:"customerSystemApiKey"`
	ShippingCompanyApiKey string `json:"shippingCompanyApiKey"`
}

func (input CreateCustomerIntegrationInput) NewCustomerIntegrationFromInput() *entity.CustomerIntegration {
	return &entity.CustomerIntegration{
		CustomerId:            input.CustomerId,
		CustomerSystemId:      input.CustomerSystemId,
		ShippingCompanyId:     input.ShippingCompanyId,
		CustomerSystemApiKey:  input.CustomerSystemApiKey,
		ShippingCompanyApiKey: input.ShippingCompanyApiKey,
		Valid:                 false,
	}
}
