package handlers

import (
	"banking-resource-api-v2/dto"
	"banking-resource-api-v2/mocks/service"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllCustomers_WhenNoCustomerFound_ReturnsEmptyArray(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)
	mockService.
		EXPECT().
		GetAllCustomers().
		Return([]dto.CustomerResponse{}, nil)

	ch := CustomerHandler{mockService}

	app := fiber.New()
	app.Get("/customers", ch.GetAllCustomers)

	req := httptest.NewRequest("GET", "/customers", nil)

	// Act
	resp, _ := app.Test(req)

	assert.Equalf(t, 200, resp.StatusCode, "Unexpected status")
}
