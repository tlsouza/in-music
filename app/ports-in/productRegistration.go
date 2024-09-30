package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

var getProductRegistrationById, geAlltProductRegistration types.HttpServerPort

func init() {
	productRegistrationController := controllers.NewProductRegistrationController(
		services.NewProductRegistrationService(
			repository.GetProductRegistrationRepositoryInstance(),
		),
	)

	getProductRegistrationById = types.HttpServerPort{
		SilentRoute: true,
		Name:        "productRegistration",
		Path:        "product_registration/:id",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productRegistrationController.GetById,
	}

	geAlltProductRegistration = types.HttpServerPort{
		SilentRoute: true,
		Name:        "productRegistration",
		Path:        "product_registration/",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productRegistrationController.GetAll,
	}

	getProductRegistrationById.Start()
	geAlltProductRegistration.Start()

}
