package in

import (
	"api/app/controllers"
	"api/app/types"
	"api/pkg/ports/adapters"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileSaveRoute(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"invalid body structure\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{},
		)

		newProfile.Controller = mockProfileController.CreateNewProfile

		req := httptest.NewRequest("POST", "/profiles", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"invalid body structure\"}", string(body))
	})

	testDescribe.Run("Should return status code 500 and '{\"code\":500,\"message\":\"unable to save profile\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				SaveFunc: func(phr types.ProfileHttpRequest) (uint64, error) {
					return uint64(1), fmt.Errorf("error")
				},
			},
		)

		newProfile.Controller = mockProfileController.CreateNewProfile

		req := httptest.NewRequest("POST", "/profiles", strings.NewReader("{\"email\":\"example@example.com\",\"firstname\":\"John\",\"lastname\":\"Doe\"}"))
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 500, resp.StatusCode)
		assert.Equal(test, "{\"code\":500,\"message\":\"unable to save profile\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and '1' always", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				SaveFunc: func(phr types.ProfileHttpRequest) (uint64, error) {
					return uint64(1), nil
				},
			},
		)

		newProfile.Controller = mockProfileController.CreateNewProfile

		req := httptest.NewRequest("POST", "/profiles", strings.NewReader("{\"email\":\"example@example.com\",\"firstname\":\"John\",\"lastname\":\"Doe\"}"))
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "1", string(body))
	})
}

func TestProfileGetById(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"int id expected in path\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{},
		)

		getProfileById.Controller = mockProfileController.GetById

		req := httptest.NewRequest("GET", "/profiles/abc", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"int id expected in path\"}", string(body))
	})

	testDescribe.Run("Should return status code 404 and '{\"code\":404,\"message\":\"profile not found\"}-'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				GetByIDFunc: func(u uint64) (*types.Profile, error) {
					return &types.Profile{}, fmt.Errorf("error")
				},
			},
		)

		getProfileById.Controller = mockProfileController.GetById

		req := httptest.NewRequest("GET", "/profiles/1", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 404, resp.StatusCode)
		assert.Equal(test, "{\"code\":404,\"message\":\"profile not found\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and '{\"id\":1,\"email\":\"email\",\"firstname\":\"name\",\"lastname\":\"last_name\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				GetByIDFunc: func(u uint64) (*types.Profile, error) {
					return &types.Profile{
						ID:        uint64(1),
						Email:     "email",
						Firstname: "name",
						Lastname:  "last_name",
					}, nil
				},
			},
		)

		getProfileById.Controller = mockProfileController.GetById

		req := httptest.NewRequest("GET", "/profiles/1", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "{\"id\":1,\"email\":\"email\",\"firstname\":\"name\",\"lastname\":\"last_name\"}", string(body))
	})
}

func TestProfileGetAll(testDescribe *testing.T) {
	testDescribe.Run("Should return status code 200 and a string representation of  []types.ProfileHttpResponse{}", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				GetAllFunc: func() []types.ProfileHttpResponse {
					return []types.ProfileHttpResponse{
						{
							ID:                   1,
							Email:                "email",
							Firstname:            "name",
							Lastname:             "lastame",
							ProductRegistrations: []types.ProductRegistrationHttpRes{},
						},
					}
				},
			},
		)

		getAllProfile.Controller = mockProfileController.GetAll

		req := httptest.NewRequest("GET", "/profiles", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "[{\"id\":1,\"email\":\"email\",\"firstname\":\"name\",\"lastname\":\"lastame\",\"product_registrations\":[]}]", string(body))
	})

}

func TestAddNewPRtoProfile(testDescribe *testing.T) {

	requestBody := `{
		"purchase_date": "2023-09-30T10:15:30Z",
		"expiry_at": "2024-09-30T10:15:30Z",
		"product": {
		  "id": 123,
		  "name": "Example Product"
		},
		"serial_code": "ABC123456",
		"additional_product_registrations": [
		  {
			"purchase_date": "2022-09-30T10:15:30Z",
			"expiry_at": null,
			"product": {
			  "id": 456,
			  "name": "Another Product"
			},
			"serial_code": "XYZ987654",
			"additional_product_registrations": []
		  }
		]
	  }`

	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"int id expected in path\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{},
		)

		addNewProductRegistrationsToProfile.Controller = mockProfileController.CreateNewProductRegistration

		req := httptest.NewRequest("POST", "/profiles/abc/product_registrations", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"int id expected in path\"}", string(body))
	})

	testDescribe.Run("Should return status code 400 and and a message with : 'invalid body structure'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{},
		)

		addNewProductRegistrationsToProfile.Controller = mockProfileController.CreateNewProductRegistration

		req := httptest.NewRequest("POST", "/profiles/1/product_registrations", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"invalid body structure\"}", string(body))
	})

	testDescribe.Run("Should return status code 404 and and a message with error'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				AddProductRegistrationFunc: func(u uint64, prhr types.ProductRegistrationHttpReq) (uint64, error) {
					return 0, fmt.Errorf("error")
				},
			},
		)

		addNewProductRegistrationsToProfile.Controller = mockProfileController.CreateNewProductRegistration

		req := httptest.NewRequest("POST", "/profiles/1/product_registrations", strings.NewReader(requestBody))
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 404, resp.StatusCode)
		assert.Equal(test, "{\"code\":404,\"message\":\"error\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and and a the root product registration id'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				AddProductRegistrationFunc: func(u uint64, prhr types.ProductRegistrationHttpReq) (uint64, error) {
					return 1, nil
				},
			},
		)

		addNewProductRegistrationsToProfile.Controller = mockProfileController.CreateNewProductRegistration

		req := httptest.NewRequest("POST", "/profiles/1/product_registrations", strings.NewReader(requestBody))
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "1", string(body))
	})

}

func TestGetProductRegistrationByProfileId(testDescribe *testing.T) {

	testDescribe.Run("Should return status code 400 and '{\"code\":400,\"message\":\"int id expected in path\"}'", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{},
		)

		getProductRegistrationsForProfile.Controller = mockProfileController.GetProductRegistrationByProfileId

		req := httptest.NewRequest("GET", "/profiles/abc/product_registrations", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 400, resp.StatusCode)
		assert.Equal(test, "{\"code\":400,\"message\":\"int id expected in path\"}", string(body))
	})

	testDescribe.Run("Should return status code 200 and message with an array of ProductRegistrationHttpRes ", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				GetProductRegistrationByProfileIdFunc: func(u uint64) ([]types.ProductRegistrationHttpRes, error) {
					return []types.ProductRegistrationHttpRes{
						{},
					}, nil
				},
			},
		)

		getProductRegistrationsForProfile.Controller = mockProfileController.GetProductRegistrationByProfileId

		req := httptest.NewRequest("GET", "/profiles/1/product_registrations", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 200, resp.StatusCode)
		assert.Equal(test, "[{\"id\":0,\"purchase_date\":\"0001-01-01T00:00:00Z\",\"expiry_at\":null,\"product\":{\"SKU\":null},\"serial_code\":\"\",\"additional_product_registrations\":null}]", string(body))
	})

	testDescribe.Run("Should return status code 404 and message with error", func(test *testing.T) {
		app := adapters.GetApp()
		mockProfileController := controllers.NewProfileController(
			&MockProfileService{
				GetProductRegistrationByProfileIdFunc: func(u uint64) ([]types.ProductRegistrationHttpRes, error) {
					return nil, fmt.Errorf("error")
				},
			},
		)

		getProductRegistrationsForProfile.Controller = mockProfileController.GetProductRegistrationByProfileId

		req := httptest.NewRequest("GET", "/profiles/1/product_registrations", nil)
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Equal(test, 404, resp.StatusCode)
		assert.Equal(test, "{\"code\":404,\"message\":\"error\"}", string(body))
	})
}
