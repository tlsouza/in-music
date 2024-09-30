package services_test

import (
	"api/app/services"
	"api/app/types"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the Save function of productService
func TestProductService_Save(t *testing.T) {
	mockRepo := &MockProductRepository{
		SaveFunc: func(product types.Product) (string, error) {
			return "saved-sku", nil
		},
	}

	productService := services.NewProductService(mockRepo)

	product := types.Product{SKU: nil}

	// Test successful save
	savedSku, err := productService.Save(product)
	assert.NoError(t, err)
	assert.Equal(t, "saved-sku", savedSku)

	// Test failure scenario
	mockRepo.SaveFunc = func(product types.Product) (string, error) {
		return "", errors.New("failed to save product")
	}

	savedSku, err = productService.Save(product)
	assert.Error(t, err)
	assert.Equal(t, "", savedSku)
	assert.Equal(t, "failed to save product", err.Error())
}

// Test the GetBySku function of productService
func TestProductService_GetBySku(t *testing.T) {
	mockRepo := &MockProductRepository{
		GetBySkuFunc: func(sku string) (*types.Product, error) {
			return &types.Product{SKU: &sku}, nil
		},
	}

	productService := services.NewProductService(mockRepo)

	// Test successful get by SKU
	sku := "test-sku"
	product, err := productService.GetBySku(sku)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "test-sku", *product.SKU)

	// Test failure scenario
	mockRepo.GetBySkuFunc = func(sku string) (*types.Product, error) {
		return nil, errors.New("product not found")
	}

	product, err = productService.GetBySku("unknown-sku")
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "product not found", err.Error())
}
