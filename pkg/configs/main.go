package configs

import (
	_ "api/pkg/dotenv"
	"os"
	"strconv"
)

var APP_NAME = Get("APP_NAME", "api")
var PORT, _ = strconv.Atoi(Get("PORT", "8000"))
var HOST_NAME = Get("HOST_NAME", "0.0.0.0")
var LOG_LEVEL = Get("LOG_LEVEL", "info")
var APP_VERSION = Get("APP_VERSION", "0.0.1")
var REQUEST_TIMEOUT, _ = strconv.Atoi(Get("REQUEST_TIMEOUT", "5"))
var ENV = Get("ENV", "dev")

func Get(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
