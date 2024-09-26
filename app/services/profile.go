package services

import (
	"api/app/repository"
	"api/app/types"
)

type profileService struct {
	repo repository.IProfileRepository
}

func NewProfileService(repo repository.IProfileRepository) *profileService {
	return &profileService{
		repo,
	}
}

func (ps *profileService) Save(profile types.ProfileHttpRequest) (uint64, error) {
	newProfile := types.Profile{
		Email:     profile.Email,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
	}
	savedProfileId, err := ps.repo.Save(newProfile)

	return savedProfileId, err
}

func (ps *profileService) GetByID(id uint64) (*types.Profile, error) {
	return ps.repo.GetByID(id)
}

func (ps *profileService) GetAll() []types.Profile {
	return ps.repo.GetAll()
}
