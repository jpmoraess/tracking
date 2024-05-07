package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/usecase"
)

type CustomerHandler struct {
	createCustomer usecase.CreateCustomer
}

// TODO: 1 handler por usecase? rever injeção de dependências..
func NewCustomerHandler(createCustomer usecase.CreateCustomer) *CustomerHandler {
	return &CustomerHandler{
		createCustomer: createCustomer,
	}
}

func (h *CustomerHandler) HandlePostCustomer(c *fiber.Ctx) error {
	var input dto.CreateCustomerInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	id, err := h.createCustomer.Execute(input)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"id": id})
}
