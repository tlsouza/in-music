package types

import (
	"api/pkg/configs"
	"api/pkg/errors"
	"api/pkg/log"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	GET     = "Get"
	POST    = "Post"
	PUT     = "Put"
	PATCH   = "Patch"
	DELETE  = "Delete"
	OPTIONS = "Options"
	HEAD    = "Head"
)

type RequestData struct {
	Ctx  context.Context
	Body map[string]interface{}
	// BodyByte armazena o body em bytes.
	// Utilizar quando for necessário fazer o Unmarshal para uma struct.
	// No futuro pode substituir o Body, já que é possível fazer o Unmarshal de []byte para map[string]any.
	BodyByte   []byte
	PathParams map[string]string
	Query      map[string]string
	Header     http.Header
}

// Headers retorna os headers da request.
// É recomendado utilizar essa função ao invés de acessar o campo Header diretamente,
// pois ela garante que o Header seja inicializado.
func (rd *RequestData) Headers() http.Header {
	if rd.Header == nil {
		rd.Header = make(http.Header)
	}

	return rd.Header
}

type HttpServerPort struct {
	Path                   string
	Verb                   string
	Validator              func(RequestData) *errors.HttpError
	Controller             func(RequestData) (interface{}, *errors.HttpError)
	Adapter                func(HttpServerPort)
	Name                   string
	StatusCodeSuccessfully int
	LogDisabled            bool
	Version                string
	SilentRoute            bool
	Doc                    *SwaggerRouteObject
}

func (port HttpServerPort) Start() {
	logger := log.New(context.Background())

	port.fill_defaults()
	logger.Info(fmt.Sprintf("start http port %s - %s %s ", port.Name, port.Verb, port.Path))
	port.Adapter(port)

	if port.Doc != nil {
		match, _ := regexp.Compile(`(:([a-zA-Z\d\_\-]+))`)
		path := match.ReplaceAllString(fmt.Sprintf("%s/%s", port.Version, port.Path), `{$2}`)

		logger.Info(fmt.Sprintf("Adding swagger doc for %s - %s %s ", port.Name, port.Verb, path))

		_, ok := configs.SwaggerConfig.Paths[path]
		if !ok {
			configs.SwaggerConfig.Paths[path] = make(map[string]any)
		}

		lowerCaseVerb := strings.ToLower(port.Verb)
		configs.SwaggerConfig.Paths[path].(map[string]any)[lowerCaseVerb] = port.Doc.SwaggerDescription()
	}

}

// GetFullPath retorna o caminho completo. Deve ser usando quando for necessário
// utilizar o Path completo da porta.
func (port HttpServerPort) GetFullPath() string {
	return fmt.Sprintf("%s/%s", port.Version, port.Path)
}

func validatorDefault(requestData RequestData) *errors.HttpError {
	return nil
}

func (port *HttpServerPort) fill_defaults() {
	if port.Validator == nil {
		port.Validator = validatorDefault
	}

	if port.StatusCodeSuccessfully == 0 {
		port.StatusCodeSuccessfully = 200
	}
}

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type HttpClientPort struct {
	Path    string
	Method  string
	Adapter func(*HttpClientPort)
	Send    func(RequestData) (*ResponseHttp, error)
	Name    string
	Ctx     context.Context

	url *url.URL
}

func (port *HttpClientPort) Start() {
	port.Ctx = context.WithValue(context.Background(), "portOut", port.Name)
	logger := log.New(port.Ctx)
	logger.Info(fmt.Sprintf("start http-out port %s - %s %s ", port.Name, port.Method, port.Path))

	parsedURL, err := url.Parse(port.Path)
	if err != nil {
		panic(fmt.Sprintf("cannot parse port path: %s", err.Error()))
	}

	port.url = parsedURL
	port.Adapter(port)
}

func (port *HttpClientPort) GetURL() *url.URL {
	return port.url
}

type ResponseHttp struct {
	Body       []byte
	StatusCode int
	Header     map[string][]string
}
