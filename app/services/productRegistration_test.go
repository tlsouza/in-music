package services_test

import (
	"api/app/services"
	"api/app/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductRegistrationService_GetBundle(t *testing.T) {
	now := time.Now()
	mockRepo := &MockProductRegistrationRepository{

		GetByBundleFunc: func(id uint64) (*types.ProductRegistration, []types.ProductRegistration) {
			rootProduct := types.Product{SKU: stringPointer("root-product")}
			childProduct := types.Product{SKU: stringPointer("child-product")}

			root := &types.ProductRegistration{
				Id:           id,
				PurchaseDate: now,
				ExpiryAt:     &now,
				Product:      rootProduct,
				SerialCode:   "root-serial",
			}
			children := []types.ProductRegistration{
				{
					Id:           2,
					PurchaseDate: now,
					ExpiryAt:     &now,
					Product:      childProduct,
					SerialCode:   "child-serial",
				},
			}
			return root, children
		},
	}

	productRegistrationService := services.NewProductRegistrationService(mockRepo)

	// Test successful GetBundle
	bundleId := uint64(1)
	res, err := productRegistrationService.GetBundle(bundleId)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, bundleId, res.Id)
	assert.Equal(t, "root-serial", res.SerialCode)

	// Assert PurchaseDate and ExpiryAt
	assert.Equal(t, now, res.PurchaseDate)
	assert.Equal(t, now, *res.ExpiryAt)

	// Assert Product details
	assert.Equal(t, "root-product", *res.Product.SKU)

	// Assert AdditionalProductRegistrations
	assert.Equal(t, 1, len(res.AdditionalProductRegistrations))
	assert.Equal(t, "child-serial", res.AdditionalProductRegistrations[0].SerialCode)

	// Assert child PurchaseDate and ExpiryAt
	assert.Equal(t, now, res.AdditionalProductRegistrations[0].PurchaseDate)
	assert.Equal(t, now, *res.AdditionalProductRegistrations[0].ExpiryAt)

	// Assert child Product details
	assert.Equal(t, "child-product", *res.AdditionalProductRegistrations[0].Product.SKU)
}

// Test the GetAll function of productRegistrationService
func TestProductRegistrationService_GetAll(t *testing.T) {
	now := time.Now()
	expectedParentId := uint64Pointer(1)
	expectedProfileId := uint64Pointer(2)
	expectedRootId := uint64Pointer(1)
	mockRepo := &MockProductRegistrationRepository{
		GetAllFunc: func() []types.ProductRegistration {
			return []types.ProductRegistration{
				{
					Id:           *expectedRootId,
					PurchaseDate: now,
					ExpiryAt:     &now,
					Product:      types.Product{SKU: stringPointer("product-1")},
					SerialCode:   "serial-1",
					RootId:       nil,
					ProfileId:    expectedProfileId,
					ParentId:     nil,
				},
				{
					Id:           2,
					PurchaseDate: now,
					ExpiryAt:     nil,
					Product:      types.Product{SKU: stringPointer("product-2")},
					SerialCode:   "serial-2",
					ParentId:     expectedParentId,
					RootId:       expectedRootId,
					ProfileId:    expectedProfileId,
				},
			}
		},
	}

	productRegistrationService := services.NewProductRegistrationService(mockRepo)

	// Test GetAll
	res := productRegistrationService.GetAll()
	assert.Equal(t, 2, len(res))

	//Assert first element
	assert.Equal(t, res[0].Id, *expectedRootId)
	assert.Equal(t, *res[0].ExpiryAt, now)
	assert.Equal(t, res[0].PurchaseDate, now)
	assert.Equal(t, res[0].SerialCode, "serial-1")
	assert.Equal(t, *res[0].Product.SKU, "product-1")
	assert.Nil(t, res[0].RootId)
	assert.Nil(t, res[0].ParentId)
	assert.Equal(t, *res[0].ProfileId, *expectedProfileId)

	//Assert second element
	assert.Equal(t, res[1].Id, uint64(2))
	assert.Nil(t, res[1].ExpiryAt)
	assert.Equal(t, res[1].PurchaseDate, now)
	assert.Equal(t, res[1].SerialCode, "serial-2")
	assert.Equal(t, *res[1].Product.SKU, "product-2")
	assert.Equal(t, *res[1].RootId, *expectedRootId)
	assert.Equal(t, *res[1].ProfileId, *expectedProfileId)
	assert.Equal(t, *res[1].ParentId, *expectedParentId)

}
