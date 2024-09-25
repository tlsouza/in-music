package configs

import (
	"encoding/json"

	"github.com/swaggo/swag"
)

type (
	swaggerConfig struct {
		Schemes  []string       `json:"schemes"`
		Swagger  string         `json:"swagger"`
		Info     swaggerInfo    `json:"info"`
		Host     string         `json:"host"`
		BasePath string         `json:"basePath"`
		Paths    map[string]any `json:"paths"`
	}

	swaggerInfo struct {
		Version     string `json:"version"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	s struct{}
)

var SwaggerConfig = swaggerConfig{
	Schemes:  []string{"HTTPS"},
	Swagger:  "2.0",
	Host:     "",
	BasePath: "/",
	Info: swaggerInfo{
		Version:     APP_VERSION,
		Title:       APP_NAME,
		Description: "Swagger Docs",
	},
	Paths: map[string]any{},
}

func init() {
	if ENV == "dev" {
		SwaggerConfig.Schemes = []string{"HTTP"}
	}
	swag.Register(swag.Name, &s{})
}

func (s *s) ReadDoc() string {
	docJson, _ := json.Marshal(SwaggerConfig)
	docString := string(docJson)
	return docString
}
