package in

import "api/app/types"

// MockProductService is a custom mock struct for IProductService
type MockProductService struct {
	SaveFunc     func(types.Product) (string, error)
	GetBySkuFunc func(string) (*types.Product, error)
}

// Save calls the mock SaveFunc if it's set, or returns a default value
func (m *MockProductService) Save(product types.Product) (string, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(product)
	}
	return "", nil // default behavior
}

// GetBySku calls the mock GetBySkuFunc if it's set, or returns a default value
func (m *MockProductService) GetBySku(sku string) (*types.Product, error) {
	if m.GetBySkuFunc != nil {
		return m.GetBySkuFunc(sku)
	}
	return nil, nil // default behavior
}

type MockProductRegistrationService struct {
	GetBundleFunc func(id uint64) (*types.ProductRegistrationHttpRes, error)
	GetAllFunc    func() []types.ProductRegistration
}

func (m *MockProductRegistrationService) GetBundle(id uint64) (*types.ProductRegistrationHttpRes, error) {
	if m.GetBundleFunc != nil {
		return m.GetBundleFunc(id)
	}
	return nil, nil
}

func (m *MockProductRegistrationService) GetAll() []types.ProductRegistration {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil
}

type MockProfileService struct {
	SaveFunc                              func(types.ProfileHttpRequest) (uint64, error)
	GetByIDFunc                           func(uint64) (*types.Profile, error)
	GetAllFunc                            func() []types.ProfileHttpResponse
	AddProductRegistrationFunc            func(uint64, types.ProductRegistrationHttpReq) (uint64, error)
	GetProductRegistrationByProfileIdFunc func(uint64) ([]types.ProductRegistrationHttpRes, error)
}

func (m *MockProfileService) Save(profile types.ProfileHttpRequest) (uint64, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(profile)
	}
	return 0, nil
}

func (m *MockProfileService) GetByID(id uint64) (*types.Profile, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

func (m *MockProfileService) GetAll() []types.ProfileHttpResponse {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil
}

func (m *MockProfileService) AddProductRegistration(profileID uint64, req types.ProductRegistrationHttpReq) (uint64, error) {
	if m.AddProductRegistrationFunc != nil {
		return m.AddProductRegistrationFunc(profileID, req)
	}
	return 0, nil
}

func (m *MockProfileService) GetProductRegistrationByProfileId(profileID uint64) ([]types.ProductRegistrationHttpRes, error) {
	if m.GetProductRegistrationByProfileIdFunc != nil {
		return m.GetProductRegistrationByProfileIdFunc(profileID)
	}
	return nil, nil
}
