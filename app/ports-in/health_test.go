package in

import (
	"api/pkg/ports/adapters"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealhCheck(testDescribe *testing.T) {
	testDescribe.Run("Should return  status code 200 and 'ok' always", func(test *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		resp, _ := adapters.GetApp().Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "\"ok\"", string(body))
	})
}
