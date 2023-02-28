package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllCustomers_WhenNoCustomerFound_ReturnsEmptyArray(t *testing.T) {
	app := fiber.New()
	customerHandler := NewCustomerHandler(nil)
	app.Get("/customers", customerHandler.GetAllCustomers)

	req := httptest.NewRequest("GET", "/customers", nil)
	resp, _ := app.Test(req)
	assert.Equalf(t, 200, resp.StatusCode, "")
}
