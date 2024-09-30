package adapters

import (
	"api/pkg/configs"
	"api/pkg/log"
	"api/pkg/ports/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"sync"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	uuid "github.com/gofrs/uuid"
)

var (
	app  *fiber.App
	once sync.Once
)

const CORRELATION_ID_HEADER = "X-Correlation-ID"

func init() {
	once.Do(loadApp)
}

func loadApp() {
	if app != nil {
		return
	}

	appConfig := fiber.Config{
		AppName:        configs.APP_NAME,
		ReadBufferSize: 16384,
	}

	app = fiber.New(appConfig)

	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			logger := log.New(c.UserContext())
			logger.Error(fmt.Errorf("panic recover: %v", e), fmt.Sprintf("panic recover in fiber - Error : %v\n%s", e, debug.Stack()))
		},
	}))

	app.Get("/docs/*", swagger.HandlerDefault)
}

func GetCorrelationId(ctx *fiber.Ctx) (correlationId string) {
	correlationId = ctx.Get(fiber.HeaderXRequestID, "")
	if correlationId != "" {
		return
	}
	correlationId = ctx.Get(CORRELATION_ID_HEADER, "")
	if correlationId != "" {
		return
	}
	newUuid, _ := uuid.NewV4()
	correlationId = newUuid.String()
	return
}

func convertToHttpRequest(c *fiber.Ctx) *http.Request {
	headers := make(http.Header)

	c.Request().Header.VisitAll(func(key, value []byte) {
		headers.Set(string(key), string(value))
	})

	reqHost := headers.Get("Host")
	if reqHost == "" {
		reqHost = string(c.Request().URI().Host())
	}

	return &http.Request{
		Method: c.Method(),
		URL: &url.URL{
			Scheme:   string(c.Request().URI().Scheme()),
			Host:     reqHost,
			Path:     string(c.Request().URI().Path()),
			RawQuery: string(c.Request().URI().QueryString()),
		},
		Header: headers,
		Host:   string(c.Request().URI().Host()),
	}
}

func createRequestData(ctx context.Context, c *fiber.Ctx) types.RequestData {

	var bodyData map[string]interface{}

	json.Unmarshal(c.Body(), &bodyData)

	requestData := types.RequestData{
		Ctx:        ctx,
		Body:       bodyData,
		BodyByte:   c.Body(),
		PathParams: c.AllParams(),
		Query:      make(map[string]string, c.Context().QueryArgs().Len()),
	}

	c.Request().Header.VisitAll(func(key, value []byte) {
		requestData.Headers().Add(string(key), string(value))
	})

	c.Context().QueryArgs().VisitAll(
		func(key, value []byte) {
			requestData.Query[string(key)] = string(value)
		},
	)

	return requestData
}

var Fiber = func(port *types.HttpServerPort) {
	var proxyController = func(c *fiber.Ctx) error {
		newUuid, _ := uuid.NewV4()
		requestId := newUuid.String()

		ctx := context.WithValue(c.UserContext(), "traceId", GetCorrelationId(c))
		ctx = context.WithValue(ctx, "requestId", requestId)
		defer ctx.Done()

		c.Set(fiber.HeaderXRequestID, requestId)

		logger := log.New(ctx)
		enableLogRoute := !port.SilentRoute

		requestData := createRequestData(ctx, c)

		url := c.Request().URI()

		if enableLogRoute {
			logger.Info(fmt.Sprintf("Sending to Controller - url: %s : %v", url.Path(), requestData))
		}

		if err := port.Validator(requestData); err != nil {
			logger.Error(err, fmt.Sprintf("Validator Error - url: %s", url.Path()))
			return c.Status(err.StatusCode).JSON(err)
		}

		var response, err = port.Controller(requestData)
		if err != nil {
			logger.Error(err, fmt.Sprintf("Controller Error - url: %s - response: %s", url.Path(), response))
			if response != nil {
				return c.Status(err.StatusCode).JSON(response)
			}
			return c.Status(err.StatusCode).JSON(err)
		}

		if enableLogRoute {
			logger.Info(fmt.Sprintf("Controller Response - url: %s, status: %d, response: %v", url.Path(), port.StatusCodeSuccessfully, response))
		}

		return c.Status(port.StatusCodeSuccessfully).JSON(response)
	}

	handler := timeout.NewWithContext(proxyController, time.Duration(configs.REQUEST_TIMEOUT)*time.Second)

	app.Add(port.Verb, port.GetFullPath(), handler)
}

func GetApp() *fiber.App {
	return app
}

var FiberListen = func() {
	logger := log.New(context.Background())
	logger.Info(fmt.Sprintf("App listen %d", configs.PORT))
	app.Listen(fmt.Sprintf(":%d", configs.PORT))
}
