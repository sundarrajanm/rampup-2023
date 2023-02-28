package handlers

import (
	"banking-resource-api-v2/service"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) GetAllCustomers(ctx *fiber.Ctx) error {
	customers, _ := ch.service.GetAllCustomers()
	return ctx.JSON(customers)
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{service}
}
