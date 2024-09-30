package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

var newProfile, getProfileById, getAllProfile, addNewProductRegistrationsToProfile, getProductRegistrationsForProfile types.HttpServerPort

func init() {
	profileController := controllers.NewProfileController(
		services.NewProfileService(
			repository.GetProfileRepositoryInstance(),
			repository.GetProductRegistrationRepositoryInstance(),
			repository.GetProductRepositoryInstance(),
		),
	)
	newProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  profileController.CreateNewProfile,
	}

	getProfileById = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles/:id",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetById,
	}

	getAllProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profiles",
		Path:        "profiles",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetAll,
	}

	addNewProductRegistrationsToProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "product_registrations",
		Path:        "profiles/:profile/product_registrations",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  profileController.CreateNewProductRegistration,
	}

	getProductRegistrationsForProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "product_registrations",
		Path:        "profiles/:profile/product_registrations",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetProductRegistrationByProfileId,
	}

	newProfile.Start()
	getProfileById.Start()
	getAllProfile.Start()
	addNewProductRegistrationsToProfile.Start()
	getProductRegistrationsForProfile.Start()
}
