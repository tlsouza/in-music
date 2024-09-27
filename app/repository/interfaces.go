package repository

import "api/app/types"

type IProfileRepository interface {
	Save(profile types.Profile) (uint64, error)
	GetByID(id uint64) (*types.Profile, error)
	GetAll() []types.Profile
}

type IProductRepository interface {
	Save(product types.Product) (string, error)
	GetBySku(sku string) (*types.Product, error)
	GetAll() []types.Product
}

type IProductRegistrationRepositoryInterface interface {
	Save(registration types.ProductRegistration) (uint64, error)
	GetByID(id uint64) (*types.ProductRegistration, error)
	GetByProfileId(profileId uint64) ([]types.ProductRegistration, error)
	GetByParentId(parentId uint64) ([]types.ProductRegistration, error)
	GetAll() []types.ProductRegistration
}
