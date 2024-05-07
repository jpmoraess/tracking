package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/tracking/internal/app/dto"
	"github.com/jpmoraess/tracking/internal/app/usecase"
)

type ShippingCompanyHandler struct {
	createShippingCompany usecase.CreateShippingCompany
}

func NewShippingCompanyHandler(createShippingCompany usecase.CreateShippingCompany) *ShippingCompanyHandler {
	return &ShippingCompanyHandler{
		createShippingCompany: createShippingCompany,
	}
}

func (h *ShippingCompanyHandler) HandlePostShippingCompany(c *fiber.Ctx) error {
	var input dto.CreateShippingCompanyInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	id, err := h.createShippingCompany.Execute(input)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"id": id})
}
