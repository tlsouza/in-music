package errors

import (
	"net/http"
	"os"
)

type HttpError struct {
	error
	StatusCode int                    `json:"code"`
	Message    string                 `json:"message"`
	Attributes map[string]interface{} `json:"-"`
}

func (httpError HttpError) Error() string {
	if httpError.Message == "" {
		return "BadRequest"
	}
	return httpError.Message
}

func NewHttpError(err error, statusCode int, attributes ...map[string]interface{}) *HttpError {
	if err == nil {
		return nil
	}
	httpError := &HttpError{Message: err.Error(), StatusCode: statusCode}
	if len(attributes) > 0 {
		httpError.Attributes = attributes[0]
	}
	return httpError
}

type HttpClientError struct {
	error
	StatusCode int
	Message    string
	Body       *[]byte
	IsTimeout  bool
}

func (httpClientError HttpClientError) Error() string {
	if httpClientError.Message == "" {
		return "http client error"
	}
	return httpClientError.Message
}

func (httpClientError *HttpClientError) Timeout() bool {
	return httpClientError.IsTimeout
}

func (httpClientError HttpClientError) IsRecoverableError() bool {
	statusCode := httpClientError.StatusCode
	return statusCode == 429 || statusCode >= 500
}

func NewHttpClientError(err error, response *http.Response) *HttpClientError {
	if err == nil {
		return nil
	}
	statusCode := 0
	if response != nil {
		statusCode = response.StatusCode
	}

	isTimeout := os.IsTimeout(err)
	return &HttpClientError{Message: err.Error(), IsTimeout: isTimeout, StatusCode: statusCode}
}
