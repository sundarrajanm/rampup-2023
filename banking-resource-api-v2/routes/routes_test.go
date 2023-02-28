package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_routes(t *testing.T) {
	tests := []struct {
		desc         string
		route        string
		expectedCode int
	}{
		{
			desc:         "Get All Customers",
			route:        "/customers",
			expectedCode: 200,
		},
		{
			desc:         "Home endpoint",
			route:        "/",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Get("/customers", func(c *fiber.Ctx) error {
		// return c.SendString("Hello World!")
		return c.SendStatus(404)
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.desc)
	}
}
