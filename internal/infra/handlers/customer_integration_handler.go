package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/usecase"
)

type CustomerIntegrationHandler struct {
	createCustomerIntegration usecase.CreateCustomerIntegration
}

func NewCustomerIntegrationHandler(createCustomerIntegration usecase.CreateCustomerIntegration) *CustomerIntegrationHandler {
	return &CustomerIntegrationHandler{
		createCustomerIntegration: createCustomerIntegration,
	}
}

func (h *CustomerIntegrationHandler) HandlePostCustomerIntegration(c *fiber.Ctx) error {
	var input dto.CreateCustomerIntegrationInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.createCustomerIntegration.Execute(input); err != nil {
		return err
	}
	return nil
}
