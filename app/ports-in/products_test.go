package in

import (
	"api/app/controllers"
	app_types "api/app/types"
	"api/pkg/ports/adapters"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRoutes(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"invalid body structure\"}' always", func(test *testing.T) {
		app := adapters.GetApp()
		mockproductController := controllers.NewProductController(
			&MockProductService{},
		)

		newProduct.Controller = mockproductController.CreateNewProduct

		req := httptest.NewRequest("POST", "/products", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"invalid body structure\"}", string(body))
	})

	testDescribe.Run("Should return status code 500 and '{\"code\":500,\"message\":\"unable to save product\"}' always", func(test *testing.T) {
		app := adapters.GetApp()

		mockproductController := controllers.NewProductController(
			&MockProductService{
				SaveFunc: func(p app_types.Product) (string, error) {
					return "", fmt.Errorf("error")
				},
			},
		)

		newProduct.Controller = mockproductController.CreateNewProduct

		req := httptest.NewRequest("POST", "/products", strings.NewReader("{\"SKU\":\"XPTO2\"}"))
		resp, err := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Nil(test, err)
		assert.Equal(test, 500, resp.StatusCode)
		assert.Equal(test, "{\"code\":500,\"message\":\"unable to save product\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and 'XPTO2' always", func(test *testing.T) {
		app := adapters.GetApp()
		mockproductController := controllers.NewProductController(
			&MockProductService{
				SaveFunc: func(p app_types.Product) (string, error) {
					return *p.SKU, nil
				},
			},
		)

		newProduct.Controller = mockproductController.CreateNewProduct

		req := httptest.NewRequest("POST", "/products", strings.NewReader("{\"SKU\":\"XPTO2\"}"))
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "\"XPTO2\"", string(body))
	})
}

func TestProductGetBySku(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 404 and '{\"code\":404,\"message\":\"product not found\"}' always", func(test *testing.T) {
		app := adapters.GetApp()
		mockproductController := controllers.NewProductController(
			&MockProductService{
				GetBySkuFunc: func(s string) (*app_types.Product, error) {
					return nil, fmt.Errorf("not found")
				},
			},
		)

		getProductBySku.Controller = mockproductController.GetBySku

		req := httptest.NewRequest("GET", "/products/XPTO", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 404, resp.StatusCode)
		assert.Equal(test, "{\"code\":404,\"message\":\"product not found\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and '{\"SKU\":\"XPTO\"}' always", func(test *testing.T) {
		app := adapters.GetApp()
		mockproductController := controllers.NewProductController(
			&MockProductService{
				GetBySkuFunc: func(s string) (*app_types.Product, error) {
					return &app_types.Product{
						SKU: StringPointer("XPTO"),
					}, nil
				},
			},
		)

		getProductBySku.Controller = mockproductController.GetBySku

		req := httptest.NewRequest("GET", "/products/XPTO", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "{\"SKU\":\"XPTO\"}", string(body))
	})
}

func StringPointer(s string) *string {
	return &s
}
