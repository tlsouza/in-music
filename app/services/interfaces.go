package services

import "api/app/types"

type IProfileService interface {
	Save(profile types.ProfileHttpRequest) (int, error)
	GetByID(id int) (*types.Profile, error)
	GetAll() []types.Profile
}

type IProductService interface {
	Save(types.Product) (string, error)
	GetBySku(sku string) (*types.Product, error)
	GetAll() []types.Product
}
