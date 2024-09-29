package services

import (
	"api/app/repository"
	"api/app/types"
)

type productService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *productService {
	return &productService{
		repo,
	}
}

func (ps *productService) Save(product types.Product) (string, error) {
	savedSku, err := ps.repo.Save(product)

	return savedSku, err
}

func (ps *productService) GetBySku(sku string) (*types.Product, error) {
	return ps.repo.GetBySku(sku)
}
