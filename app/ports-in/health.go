package in

import (
	"api/pkg/errors"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

func init() {
	var health = types.HttpServerPort{
		SilentRoute: true,
		Name:        "Health",
		Path:        "health",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller: func(requestData types.RequestData) (interface{}, *errors.HttpError) {
			return "ok", nil
		},
		Doc: types.AddSwagger("Health Check Route",
			types.WithResponseOK("ok"),
		),
	}

	health.Start()
}
