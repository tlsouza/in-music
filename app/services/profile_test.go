package services_test

import (
	"api/app/services"
	"api/app/types"
	"api/pkg/errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProfileService_Save_Success(t *testing.T) {
	profileToSave := types.ProfileHttpRequest{
		Email:     "expected_email",
		Firstname: "first_name",
		Lastname:  "last_name",
	}

	expectedId := uint64(1)
	profileService := services.NewProfileService(
		&MockProfileRepository{
			SaveFunc: func(profile types.Profile) (uint64, error) {
				assert.Equal(t, profile.Email, profileToSave.Email)
				assert.Equal(t, profile.Firstname, profileToSave.Firstname)
				assert.Equal(t, profile.Lastname, profileToSave.Lastname)
				return expectedId, nil
			},
		},
		&MockProductRegistrationRepository{},
		&MockProductRepository{},
	)

	newId, err := profileService.Save(profileToSave)
	assert.Nil(t, err)
	assert.Equal(t, newId, expectedId)
}

func TestProfileService_Save_Error(t *testing.T) {
	profileService := services.NewProfileService(
		&MockProfileRepository{
			SaveFunc: func(profile types.Profile) (uint64, error) {
				return 0, fmt.Errorf("error_creating")
			},
		}, &MockProductRegistrationRepository{}, &MockProductRepository{},
	)
	_, err := profileService.Save(types.ProfileHttpRequest{})
	assert.Equal(t, "error_creating", err.Error())
}

func TestProfileService_GetByID_Success(t *testing.T) {
	expectedId := uint64(1)
	expectedProfile := types.Profile{
		ID:        expectedId,
		Email:     "expected_email",
		Firstname: "first_name",
		Lastname:  "last_name",
	}

	profileService := services.NewProfileService(
		&MockProfileRepository{
			GetByIDFunc: func(id uint64) (*types.Profile, error) {
				return &expectedProfile, nil
			},
		}, &MockProductRegistrationRepository{}, &MockProductRepository{},
	)

	profile, err := profileService.GetByID(expectedId)
	assert.Nil(t, err)
	assert.Equal(t, profile.Email, expectedProfile.Email)
	assert.Equal(t, profile.Firstname, expectedProfile.Firstname)
	assert.Equal(t, profile.Lastname, expectedProfile.Lastname)
}

func TestProfileService_GetByID_Error(t *testing.T) {
	profileService := services.NewProfileService(
		&MockProfileRepository{
			GetByIDFunc: func(id uint64) (*types.Profile, error) {
				return &types.Profile{}, fmt.Errorf("not_found")
			},
		}, &MockProductRegistrationRepository{}, &MockProductRepository{},
	)
	_, err := profileService.GetByID(uint64(1))
	assert.Equal(t, "not_found", err.Error())
}

func TestProfileService_GetAll_Success(t *testing.T) {
	expectedId := uint64(1)
	expecteProfileId := uint64(1)
	expecteParentId := uint64Pointer(1)
	expectedRoot := uint64Pointer(1)
	now := time.Now()
	expectedProfile := types.Profile{
		ID:        expectedId,
		Email:     "expected_email",
		Firstname: "first_name",
		Lastname:  "last_name",
	}

	expectedProductRegistrations := []types.ProductRegistration{
		{
			Id:           1,
			PurchaseDate: now,
			ExpiryAt:     &now,
			Product:      types.Product{SKU: stringPointer("Product A")},
			SerialCode:   "ABC123",
			ProfileId:    &expecteProfileId,
			ParentId:     nil,
			RootId:       nil,
		},
		{
			Id:           2,
			PurchaseDate: now,
			ExpiryAt:     &now,
			Product:      types.Product{SKU: stringPointer("Product B")},
			SerialCode:   "XYZ789",
			ProfileId:    &expecteProfileId, // Example of a profile linked
			ParentId:     expecteParentId,   // Example of a parent registration
			RootId:       expectedRoot,
		},
	}

	expectedProducRegistrationHttpRes := []types.ProductRegistrationHttpRes{
		{
			Id:           expectedProductRegistrations[0].Id,
			PurchaseDate: expectedProductRegistrations[0].PurchaseDate,
			ExpiryAt:     expectedProductRegistrations[0].ExpiryAt,
			SerialCode:   expectedProductRegistrations[0].SerialCode,
			Product:      expectedProductRegistrations[0].Product,
			AdditionalProductRegistrations: []types.ProductRegistrationHttpResChild{{
				Id:           expectedProductRegistrations[1].Id,
				PurchaseDate: expectedProductRegistrations[1].PurchaseDate,
				ExpiryAt:     expectedProductRegistrations[1].ExpiryAt,
				SerialCode:   expectedProductRegistrations[1].SerialCode,
				Product:      expectedProductRegistrations[1].Product,
			},
			},
		},
	}

	profileService := services.NewProfileService(
		&MockProfileRepository{
			GetAllFunc: func() []types.Profile {
				return []types.Profile{
					expectedProfile,
				}
			},
			GetByIDFunc: func(id uint64) (*types.Profile, error) {
				return &expectedProfile, nil
			},
		}, &MockProductRegistrationRepository{
			GetByProfileIdFunc: func(profileId uint64) []types.ProductRegistration {
				return expectedProductRegistrations
			},
		}, &MockProductRepository{},
	)

	profiles := profileService.GetAll()
	assert.Equal(t, len(profiles), 1)
	//assertProfileFields
	assert.Equal(t, len(profiles[0].ProductRegistrations), 1)
	assert.Equal(t, profiles[0].ProductRegistrations[0], expectedProducRegistrationHttpRes[0])
}

func TestAddProductRegistration(t *testing.T) {
	mockProfileRepo := &MockProfileRepository{}
	mockProductRegRepo := &MockProductRegistrationRepository{}
	mockProductRepo := &MockProductRepository{}

	profileService := services.NewProfileService(mockProfileRepo, mockProductRegRepo, mockProductRepo)

	profileId := uint64(1)
	productSku := "SKU123"
	now := time.Now()
	expiry := now.AddDate(1, 0, 0)

	// Setup ProductRegistrationHttpReq mock
	productRegistrationReq := types.ProductRegistrationHttpReq{
		PurchaseDate: now,
		ExpiryAt:     &expiry,
		Product:      types.Product{SKU: &productSku},
		SerialCode:   "SER123",
	}

	// Test Case 1: Profile not found
	mockProfileRepo.GetByIDFunc = func(id uint64) (*types.Profile, error) {
		return nil, errors.NewHttpError(nil, 404)
	}

	_, err := profileService.AddProductRegistration(profileId, productRegistrationReq)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(*errors.HttpError).StatusCode)

	// Test Case 2: Product not found
	mockProfileRepo.GetByIDFunc = func(id uint64) (*types.Profile, error) {
		return &types.Profile{ID: id, Email: "test@example.com"}, nil
	}
	mockProductRepo.GetBySkuFunc = func(sku string) (*types.Product, error) {
		return nil, errors.NewHttpError(nil, 404)
	}

	_, err = profileService.AddProductRegistration(profileId, productRegistrationReq)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(*errors.HttpError).StatusCode)

	// Test Case 3: Successful Product Registration
	mockProductRepo.GetBySkuFunc = func(sku string) (*types.Product, error) {
		return &types.Product{SKU: &sku}, nil
	}
	mockProductRegRepo.SaveFunc = func(registration types.ProductRegistration) (uint64, error) {
		return 100, nil
	}

	rootId, err := profileService.AddProductRegistration(profileId, productRegistrationReq)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), rootId)

	// Test Case 4: Additional Product Registrations
	productRegistrationReqWithChildren := types.ProductRegistrationHttpReq{
		PurchaseDate: now,
		ExpiryAt:     &expiry,
		Product:      types.Product{SKU: &productSku},
		SerialCode:   "SER456",
		AdditionalProductRegistrations: []types.ProductRegistrationHttpReq{
			{
				PurchaseDate: now,
				ExpiryAt:     &expiry,
				Product:      types.Product{SKU: &productSku},
				SerialCode:   "SER789",
			},
		},
	}

	mockProductRegRepo.SaveFunc = func(registration types.ProductRegistration) (uint64, error) {
		if registration.SerialCode == "SER456" {
			return 101, nil
		}
		return 102, nil
	}

	rootId, err = profileService.AddProductRegistration(profileId, productRegistrationReqWithChildren)
	assert.Nil(t, err)
	assert.Equal(t, uint64(101), rootId)
}
