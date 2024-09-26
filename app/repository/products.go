package repository

import (
	"api/app/types"
	"errors"
	"sync"
)

// InMemoryRepository struct implementing Repository interface
type inMemoryProductRepository struct {
	products []types.Product
	mu       sync.Mutex // to ensure thread-safe operations
}

// NewInMemoryRepository creates a new InMemoryRepository
func NewInMemoryProductRepository() *inMemoryProductRepository {
	return &inMemoryProductRepository{
		products: []types.Product{},
	}
}

// Save adds a new Product to the repository
func (r *inMemoryProductRepository) Save(Product types.Product) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Add the Product to the slice
	r.products = append(r.products, Product)
	return Product.SKU, nil
}

// GetBySku retrieves a Product by its SKU
func (r *inMemoryProductRepository) GetBySku(sku string) (*types.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, Product := range r.products {
		if Product.SKU == sku {
			return &Product, nil
		}
	}
	return nil, errors.New("product not found")
}

// GetAll retrieves all Products
func (r *inMemoryProductRepository) GetAll() []types.Product {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.products
}
