package services_test

import "api/app/types"

// Mock implementation of the IProductRepository
type MockProductRepository struct {
	SaveFunc     func(product types.Product) (string, error)
	GetBySkuFunc func(sku string) (*types.Product, error)
}

func (m *MockProductRepository) Save(product types.Product) (string, error) {
	return m.SaveFunc(product)
}

func (m *MockProductRepository) GetBySku(sku string) (*types.Product, error) {
	return m.GetBySkuFunc(sku)
}

// Mock implementation of the IProductRegistrationRepository
type MockProductRegistrationRepository struct {
	GetByBundleFunc    func(id uint64) (*types.ProductRegistration, []types.ProductRegistration)
	GetAllFunc         func() []types.ProductRegistration
	SaveFunc           func(registration types.ProductRegistration) (uint64, error)
	GetByIDFunc        func(id uint64) (*types.ProductRegistration, error)
	GetByProfileIdFunc func(profileId uint64) []types.ProductRegistration
	GetByParentIdFunc  func(parentId uint64) []types.ProductRegistration
}

func (m *MockProductRegistrationRepository) GetByBundle(id uint64) (*types.ProductRegistration, []types.ProductRegistration) {
	return m.GetByBundleFunc(id)
}

func (m *MockProductRegistrationRepository) GetAll() []types.ProductRegistration {
	return m.GetAllFunc()
}

func (m *MockProductRegistrationRepository) Save(registration types.ProductRegistration) (uint64, error) {
	return m.SaveFunc(registration)
}

func (m *MockProductRegistrationRepository) GetByID(id uint64) (*types.ProductRegistration, error) {
	return m.GetByIDFunc(id)
}

func (m *MockProductRegistrationRepository) GetByProfileId(profileId uint64) []types.ProductRegistration {
	return m.GetByProfileIdFunc(profileId)
}
func (m *MockProductRegistrationRepository) GetByParentId(parentId uint64) []types.ProductRegistration {
	return m.GetByParentIdFunc(parentId)
}

type MockProfileRepository struct {
	SaveFunc    func(profile types.Profile) (uint64, error)
	GetByIDFunc func(id uint64) (*types.Profile, error)
	GetAllFunc  func() []types.Profile
}

func (m *MockProfileRepository) Save(profile types.Profile) (uint64, error) {
	return m.SaveFunc(profile)
}

func (m *MockProfileRepository) GetByID(id uint64) (*types.Profile, error) {
	return m.GetByIDFunc(id)
}

func (m *MockProfileRepository) GetAll() []types.Profile {
	return m.GetAllFunc()
}

// Helper function to get a pointer to a string
func stringPointer(s string) *string {
	return &s
}

// Helper function to get a pointer to a string
func uint64Pointer(id uint64) *uint64 {
	return &id
}
