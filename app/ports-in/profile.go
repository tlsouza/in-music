package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

func init() {
	profileController := controllers.NewProfileController(
		services.NewProfileService(
			repository.GetProfileRepositoryInstance(),
			repository.GetProductRegistrationRepositoryInstance(),
			repository.GetProductRepositoryInstance(),
		),
	)
	var newProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  profileController.CreateNewProfile,
	}

	var GetProfileById = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles/:id",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetById,
	}

	var GetAllProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetAll,
	}

	var AddNewProductRegistrationsToProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "product_registrations",
		Path:        "profiles/:profile/product_registrations",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  profileController.CreateNewProductRegistration,
	}

	var GetProductRegistrationsForProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "product_registrations",
		Path:        "profiles/:profile/product_registrations",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetProductRegistrationByProfileId,
	}

	newProfile.Start()
	GetProfileById.Start()
	GetAllProfile.Start()
	AddNewProductRegistrationsToProfile.Start()
	GetProductRegistrationsForProfile.Start()
}
