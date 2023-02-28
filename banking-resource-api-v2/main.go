package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CheckMandatoryEnvVar(v string) string {
	value := os.Getenv(v)
	fmt.Printf("Checking env var: '%s', isEmpty: '%v'\n", v, value == "")
	if value == "" {
		panic("Env variable " + v + " not found")
	}
	return value
}

func main() {
	fiberApp := fiber.New()
	host := CheckMandatoryEnvVar("API_HOST")
	port := CheckMandatoryEnvVar("API_PORT")
	fiberApp.Listen(fmt.Sprintf("%s:%s", host, port))
}
