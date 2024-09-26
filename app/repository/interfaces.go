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
