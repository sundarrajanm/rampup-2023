package utils

import (
	"banking-resource-api/logger"
	"fmt"
	"os"
)

func CheckMandatoryEnvVar(v string) string {
	value := os.Getenv(v)
	logger.Info(fmt.Sprintf("Checking env var: '%s', isEmpty: '%v'", v, value == ""))
	if value == "" {
		errMsg := "Env variable " + v + " not found"
		logger.Error(errMsg)
		panic(errMsg)
	}
	return value
}
