package repository

import (
	"api/app/types"
	"errors"
	"sync"
)

var lockProductRepository = &sync.Mutex{}
var singleProductRepositoryInstance *ProductRepository

// InMemoryRepository struct implementing Repository interface
type ProductRepository struct {
	products []types.Product
	mu       sync.Mutex // to ensure thread-safe operations
}

// NewInMemoryRepository creates a new InMemoryRepository
func GetProductRepositoryInstance() *ProductRepository {
	if singleProductRepositoryInstance == nil {
		lockProductRepository.Lock()
		defer lockProductRepository.Unlock()
		if singleProductRepositoryInstance == nil {
			singleProductRepositoryInstance = &ProductRepository{
				products: []types.Product{},
			}
		}
	}
	return singleProductRepositoryInstance
}

// Save adds a new Product to the repository
func (r *ProductRepository) Save(Product types.Product) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Add the Product to the slice
	r.products = append(r.products, Product)
	return *Product.SKU, nil
}

// GetBySku retrieves a Product by its SKU
func (r *ProductRepository) GetBySku(sku string) (*types.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, Product := range r.products {
		if *Product.SKU == sku {
			return &Product, nil
		}
	}
	return nil, errors.New("product not found")
}
