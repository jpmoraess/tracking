package usecase

import (
	"context"
	"fmt"

	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/repository"
)

type CreateCustomerIntegration struct {
	customerRepository            repository.CustomerRepository
	customerSystemRepository      repository.CustomerSystemRepository
	shippingCompanyRepository     repository.ShippingCompanyRepository
	customerIntegrationRepository repository.CustomerIntegrationRepository
}

func NewCreateCustomerIntegration(
	customerRepository repository.CustomerRepository,
	customerSystemRepository repository.CustomerSystemRepository,
	shippingCompanyRepository repository.ShippingCompanyRepository,
	customerIntegrationRepository repository.CustomerIntegrationRepository,
) *CreateCustomerIntegration {
	return &CreateCustomerIntegration{
		customerRepository:            customerRepository,
		customerSystemRepository:      customerSystemRepository,
		shippingCompanyRepository:     shippingCompanyRepository,
		customerIntegrationRepository: customerIntegrationRepository,
	}
}

// TODO: create method exists by id for customer, customer system and shipping company
func (c *CreateCustomerIntegration) Execute(input dto.CreateCustomerIntegrationInput) error {
	var ctx = context.Background()
	_, err := c.customerRepository.GetByID(ctx, input.CustomerId)
	if err != nil {
		return fmt.Errorf("customer with id: %s not found", input.CustomerId)
	}
	_, err = c.customerSystemRepository.GetByID(ctx, input.CustomerSystemId)
	if err != nil {
		return fmt.Errorf("customer system with id: %s not found", input.CustomerSystemId)
	}
	_, err = c.shippingCompanyRepository.GetByID(ctx, input.ShippingCompanyId)
	if err != nil {
		return fmt.Errorf("shipping company with id: %s not found", input.ShippingCompanyId)
	}
	// TODO: call customer system and shipping company to validate api key's
	entity := input.NewCustomerIntegrationFromInput()
	saved, err := c.customerIntegrationRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	fmt.Printf("customer integration created successfully: %+v", saved)
	return nil
}
