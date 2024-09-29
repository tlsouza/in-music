package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

func init() {
	productRegistrationController := controllers.NewProductRegistrationController(
		services.NewProductRegistrationService(
			repository.GetProductRegistrationRepositoryInstance(),
		),
	)

	var GetProductRegistrationById = types.HttpServerPort{
		SilentRoute: true,
		Name:        "productRegistration",
		Path:        "product_registration/:id",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productRegistrationController.GetById,
	}

	var GeAlltProductRegistration = types.HttpServerPort{
		SilentRoute: true,
		Name:        "productRegistration",
		Path:        "product_registration/",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productRegistrationController.GetAll,
	}

	GetProductRegistrationById.Start()
	GeAlltProductRegistration.Start()

}
