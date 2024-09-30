package in

import (
	"api/app/controllers"
	"api/app/types"
	"api/pkg/ports/adapters"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRegistrationService(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"int id expected in path\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockPRControler := controllers.NewProductRegistrationController(
			&MockProductRegistrationService{},
		)

		getProductRegistrationById.Controller = mockPRControler.GetById

		req := httptest.NewRequest("GET", "/product_registration/acb", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"int id expected in path\"}", string(body))
	})

	testDescribe.Run("Should return status code 404 and '{\"code\":404,\"message\":\"not found\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockPRControler := controllers.NewProductRegistrationController(
			&MockProductRegistrationService{
				GetBundleFunc: func(id uint64) (*types.ProductRegistrationHttpRes, error) {
					return nil, fmt.Errorf("error")
				},
			},
		)

		getProductRegistrationById.Controller = mockPRControler.GetById

		req := httptest.NewRequest("GET", "/product_registration/1", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 404, resp.StatusCode)
		assert.Equal(test, "{\"code\":404,\"message\":\"not found\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and []", func(test *testing.T) {
		app := adapters.GetApp()
		mockPRControler := controllers.NewProductRegistrationController(
			&MockProductRegistrationService{
				GetBundleFunc: func(id uint64) (*types.ProductRegistrationHttpRes, error) {
					return &types.ProductRegistrationHttpRes{}, nil
				},
			},
		)

		getProductRegistrationById.Controller = mockPRControler.GetById

		req := httptest.NewRequest("GET", "/product_registration/1", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, `{"id":0,"purchase_date":"0001-01-01T00:00:00Z","expiry_at":null,"product":{"SKU":null},"serial_code":"","additional_product_registrations":null}`, string(body))
	})
}
