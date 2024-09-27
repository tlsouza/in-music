package services

import "api/app/types"

type IProfileService interface {
	Save(profile types.ProfileHttpRequest) (uint64, error)
	GetByID(id uint64) (*types.Profile, error)
	GetAll() []types.Profile
	AddProductRegistration(uint64, types.ProductRegistrationHttpRequest) ([]uint64, error)
}

type IProductService interface {
	Save(types.Product) (string, error)
	GetBySku(sku string) (*types.Product, error)
	GetAll() []types.Product
}
