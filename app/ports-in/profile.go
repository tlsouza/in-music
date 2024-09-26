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
			repository.NewInMemoryRepository(),
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

	newProfile.Start()
	GetProfileById.Start()
	GetAllProfile.Start()
}
