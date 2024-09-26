package services

import "api/app/types"

type IProfileService interface {
	Save(profile types.ProfileHttpRequest) (int, error)
	GetByID(id int) (*types.Profile, error)
	GetAll() []types.Profile
}
