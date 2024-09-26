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
		Name:        "profile",
		Path:        "profile",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  profileController.CreateNewProfile,
	}

	var GetProfile = types.HttpServerPort{
		SilentRoute: true,
		Name:        "profile",
		Path:        "profile/:id",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  profileController.GetById,
	}

	newProfile.Start()
	GetProfile.Start()
}
