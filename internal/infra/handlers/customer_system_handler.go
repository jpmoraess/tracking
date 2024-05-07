package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/usecase"
)

type CustomerSystemHandler struct {
	createCustomerSystem usecase.CreateCustomerSystem
}

func NewCustomerSystemHandler(createCustomerSystem usecase.CreateCustomerSystem) *CustomerSystemHandler {
	return &CustomerSystemHandler{
		createCustomerSystem: createCustomerSystem,
	}
}

func (h *CustomerSystemHandler) HandlePostCustomerSystem(c *fiber.Ctx) error {
	var input dto.CreateCustomerSystemInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	id, err := h.createCustomerSystem.Execute(input)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"id": id})
}
