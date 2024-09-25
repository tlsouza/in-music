package types

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func (s *SwaggerRouteObject) SwaggerDescription() *SwaggerDescription {
	paramLength := len(s.Headers) + len(s.Path) + len(s.Query)
	parameters := make([]SwaggerParameters, 0, paramLength)
	response := s.mapResponses()

	request, err := s.mapRequest()
	if err == nil {
		parameters = append(parameters, request)
	}

	parameters = append(parameters, s.mapParameters(s.Headers, "header")...)
	parameters = append(parameters, s.mapParameters(s.Path, "path")...)
	parameters = append(parameters, s.mapParameters(s.Query, "query")...)

	return &SwaggerDescription{
		Description: s.Description,
		Parameters:  &parameters,
		Responses:   response,
	}
}

func (s *SwaggerRouteObject) mapRequest() (SwaggerParameters, error) {
	structType := reflect.TypeOf(s.Request)
	properties, err := getProperties(structType)

	if err == nil {
		return SwaggerParameters{
			In:       "body",
			Name:     "body",
			Required: true,
			Schema: &SwaggerSchema{
				Type:       "object",
				Properties: &properties,
			},
		}, nil
	}
	return SwaggerParameters{}, errors.New("no request found")
}

func (s *SwaggerRouteObject) mapResponses() *map[string]SwaggerResponse {
	responses := make(map[string]SwaggerResponse)

	for _, response := range s.Responses {

		structType := reflect.TypeOf(response.Body)

		if structType.Kind() == reflect.Struct {
			properties, err := getProperties(structType)
			if err == nil {
				responses[strconv.Itoa(response.Status)] = SwaggerResponse{
					Description: response.Description,
					Schema: &SwaggerSchema{
						Type:       "object",
						Properties: &properties,
					},
				}
			}
		} else {
			responses[strconv.Itoa(response.Status)] = SwaggerResponse{
				Description: response.Description,
				Schema: &SwaggerSchema{
					Type: "string",
					Item: response.Body,
				},
			}
		}

	}
	return &responses
}

func (s *SwaggerRouteObject) mapParameters(parameters map[string]string, in string) []SwaggerParameters {
	if parameters == nil {
		return []SwaggerParameters{}
	}
	swaggerParameters := make([]SwaggerParameters, 0, len(parameters))

	for key, value := range parameters {
		swaggerParameters = append(swaggerParameters, SwaggerParameters{
			In:          in,
			Name:        key,
			Description: value,
			Required:    false,
			Type:        "string",
		})
	}
	return swaggerParameters
}

func getProperties(structType reflect.Type) (map[string]any, error) {
	if structType == nil {
		return nil, errors.New("properties not found")
	}

	properties := make(map[string]any)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName, _, _ := strings.Cut(field.Tag.Get("json"), ",")

		switch field.Type.Kind() {
		case reflect.Struct:
			if checkInternalStruct(field.Type) {
				properties[fieldName] = map[string]any{
					"type": typeOf(field.Type.Kind()),
				}
			} else {
				subProperties, err := getProperties(field.Type)
				if err == nil {
					properties[fieldName] = map[string]any{
						"type":       "object",
						"properties": subProperties,
					}
				}
			}
		case reflect.Ptr:
			if checkInternalStruct(field.Type.Elem()) {
				properties[fieldName] = map[string]any{
					"type": typeOf(field.Type.Kind()),
				}
			} else {
				subProperties, err := getProperties(field.Type.Elem())
				if err == nil {
					properties[fieldName] = map[string]any{
						"type":       "object",
						"properties": subProperties,
					}
				}
			}
		case reflect.Array, reflect.Slice:
			subType := field.Type.Elem()
			if subType.Kind() < 17 { //Primitive types
				properties[fieldName] = map[string]any{
					"type": typeOf(field.Type.Kind()),
				}
			} else {
				subProperties, err := getProperties(subType)
				if err == nil {
					properties[fieldName] = map[string]any{
						"type": "array",
						"items": map[string]any{
							"type":       "object",
							"properties": subProperties,
						},
					}
				}
			}

		default:
			properties[fieldName] = map[string]any{
				"type": typeOf(field.Type.Kind()),
			}
		}
	}

	if len(properties) > 0 {
		return properties, nil
	}
	return nil, errors.New("properties not found")
}

func typeOf(kind reflect.Kind) string {
	switch kind {
	case reflect.String:
		return "string"
	case reflect.Int:
		return "integer"
	case reflect.Bool:
		return "boolean"
	case reflect.Float32, reflect.Float64:
		return "number"
	default:
		return "string"
	}
}

func checkInternalStruct(structType reflect.Type) bool {
	if structType.NumField() > 0 && len(structType.Field(0).PkgPath) > 0 {
		return true
	}
	return false
}
