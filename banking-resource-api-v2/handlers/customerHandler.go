package handlers

import (
	"banking-resource-api-v2/service"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (c CustomerHandler) GetAllCustomers(ctx *fiber.Ctx) error {
	return ctx.SendStatus(400)
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{service}
}
