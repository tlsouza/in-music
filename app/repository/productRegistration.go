package repository

import (
	"api/app/types"
	"errors"
	"sync"
)

type ProductRegistrationRepository struct {
	registrations []types.ProductRegistration
	nextID        uint64
	mu            sync.Mutex
}

var lockProductRegisRepo = &sync.Mutex{}
var singleProductRegisRepoInstance *ProductRegistrationRepository

func GetProductRegistrationRepositoryInstance() *ProductRegistrationRepository {
	if singleProductRegisRepoInstance == nil {
		lockProductRegisRepo.Lock()
		defer lockProductRegisRepo.Unlock()
		if singleProductRegisRepoInstance == nil {
			singleProductRegisRepoInstance = &ProductRegistrationRepository{
				registrations: []types.ProductRegistration{},
				nextID:        1,
			}
		}
	}
	return singleProductRegisRepoInstance
}

func (r *ProductRegistrationRepository) Save(registration types.ProductRegistration) (uint64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	registration.Id = r.nextID
	r.nextID++

	r.registrations = append(r.registrations, registration)
	return registration.Id, nil
}

func (r *ProductRegistrationRepository) GetByID(id uint64) (*types.ProductRegistration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, registration := range r.registrations {
		if registration.Id == id {
			return &registration, nil
		}
	}
	return nil, errors.New("product registration not found")
}

func (r *ProductRegistrationRepository) GetByParentId(parentId uint64) []types.ProductRegistration {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []types.ProductRegistration
	for _, registration := range r.registrations {
		if registration.ParentId != nil && *registration.ParentId == parentId && registration.ProfileId == nil {
			result = append(result, registration)
		}
	}

	return result
}

func (r *ProductRegistrationRepository) GetAll() []types.ProductRegistration {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.registrations
}

func (r *ProductRegistrationRepository) GetByBundle(id uint64) (*types.ProductRegistration, []types.ProductRegistration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var root *types.ProductRegistration = nil
	children := []types.ProductRegistration{}

	for _, registration := range r.registrations {
		if registration.Id == id {
			root = &registration
		}
		if registration.RootId != nil {
			if *registration.RootId == id {
				children = append(children, registration)
			}
		}
	}
	return root, children

}

func (r *ProductRegistrationRepository) GetByProfileId(id uint64) []types.ProductRegistration {
	r.mu.Lock()
	defer r.mu.Unlock()
	pr := []types.ProductRegistration{}

	for _, registration := range r.registrations {
		if registration.ProfileId != nil && *registration.ProfileId == id {
			pr = append(pr, registration)
		}
	}
	return pr
}
