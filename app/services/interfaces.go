package services

import "api/app/types"

type IProfileService interface {
	Save(profile types.ProfileHttpRequest) (uint64, error)
	GetByID(id uint64) (*types.Profile, error)
	GetAll() []types.ProfileHttpResponse
	AddProductRegistration(uint64, types.ProductRegistrationHttpReq) (uint64, error)
	GetProductRegistrationByProfileId(uint64) ([]types.ProductRegistrationHttpRes, error)
}

type IProductService interface {
	Save(types.Product) (string, error)
	GetBySku(sku string) (*types.Product, error)
}

type IProductRegistrationService interface {
	GetBundle(id uint64) (*types.ProductRegistrationHttpRes, error)
	GetAll() []types.ProductRegistration
}
