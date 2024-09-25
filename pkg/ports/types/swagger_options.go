package types

type swaggerOptFn func(*SwaggerRouteObject)

func AddSwagger(description string, opt ...swaggerOptFn) *SwaggerRouteObject {
	route := SwaggerRouteObject{
		Description: description,
		Headers:     make(map[string]string),
		Path:        make(map[string]string),
		Query:       make(map[string]string),
	}
	for _, o := range opt {
		o(&route)
	}
	return &route
}

func WithRequest(request any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Request = request
	}
}

func WithHeader(key string, description string) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Headers[key] = description
	}
}

func WithQuery(key string, description string) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Query[key] = description
	}
}

func WithParam(key string, description string) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Path[key] = description
	}
}

func WithResponse(status int, description string, body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      status,
			Description: description,
			Body:        body,
		})
	}
}

func WithResponseOK(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      200,
			Description: "OK",
			Body:        body,
		})
	}
}

func WithResponseCreated(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      201,
			Description: "Created",
			Body:        body,
		})
	}
}

func WithResponseBadRequest(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      400,
			Description: "Bad Request",
			Body:        body,
		})
	}
}

func WithResponseUnauthorized(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      401,
			Description: "Unauthorized",
			Body:        body,
		})
	}
}

func WithResponseForbidden(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      403,
			Description: "Forbidden",
			Body:        body,
		})
	}
}

func WithResponseNotFound(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      404,
			Description: "Not Found",
			Body:        body,
		})
	}
}

func WithResponseUnprocessableEntity(body any) swaggerOptFn {
	return func(s *SwaggerRouteObject) {
		s.Responses = append(s.Responses, DocResponse{
			Status:      422,
			Description: "Unprocessable Entity",
			Body:        body,
		})
	}
}
