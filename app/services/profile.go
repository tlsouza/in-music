package services

import (
	"api/app/repository"
	"api/app/types"
	"api/pkg/errors"
	"fmt"
)

type profileService struct {
	repo repository.IProfileRepository
}

func NewProfileService(repo repository.IProfileRepository) *profileService {
	return &profileService{
		repo,
	}
}

func (ps *profileService) Save(profile types.ProfileHttpRequest) (uint64, error) {
	newProfile := types.Profile{
		Email:     profile.Email,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
	}
	savedProfileId, err := ps.repo.Save(newProfile)

	return savedProfileId, err
}

func (ps *profileService) GetByID(id uint64) (*types.Profile, error) {
	return ps.repo.GetByID(id)
}

func (ps *profileService) GetAll() []types.Profile {
	return ps.repo.GetAll()
}

func (ps *profileService) AddProductRegistration(profileId uint64, productRegistration types.ProductRegistrationHttpReq) (uint64, error) {
	// validate user
	if _, err := ps.repo.GetByID(profileId); err != nil {
		return 0, errors.NewHttpError(fmt.Errorf("user not found"), 404)
	}

	//validate root Product
	if err := ValidateProductRegistrationsProducts(productRegistration); err != nil {
		return 0, errors.NewHttpError(err, 404)
	}

	rootId := saveProductRegistrationsProductRoot(productRegistration, profileId)
	saveProductRegistrationsProductChildren(productRegistration.AdditionalProductRegistrations, profileId, rootId, rootId)

	return rootId, nil
}

func ValidateProductRegistrationsProducts(productRegistration types.ProductRegistrationHttpReq) error {
	productRepo := repository.GetProductRepositoryInstance()

	// Check if the product exists for the current request
	if product, err := productRepo.GetBySku(*productRegistration.Product.SKU); err != nil || product == nil {
		return errors.NewHttpError(fmt.Errorf("product %s not found", *productRegistration.Product.SKU), 404)
	}

	// Recursively validate products in the child registrations
	for _, httpReqChild := range productRegistration.AdditionalProductRegistrations {
		if err := ValidateProductRegistrationsProducts(httpReqChild); err != nil {
			return err
		}
	}

	return nil
}

func saveProductRegistrationsProductRoot(productRegistration types.ProductRegistrationHttpReq, profileId uint64) uint64 {
	prRepo := repository.GetProductRegistrationRepositoryInstance()

	newId, _ := prRepo.Save(types.ProductRegistration{
		PurchaseDate: productRegistration.PurchaseDate,
		ExpiryAt:     productRegistration.ExpiryAt,
		Product:      productRegistration.Product,
		SerialCode:   productRegistration.SerialCode,
		RootId:       nil,
		ProfileId:    &profileId,
		ParentId:     nil,
	})

	return newId

}

func saveProductRegistrationsProductChildren(productRegistration []types.ProductRegistrationHttpReq, profileId uint64, rootId uint64, parentId uint64) {
	prRepo := repository.GetProductRegistrationRepositoryInstance()

	for _, v := range productRegistration {
		newId, _ := prRepo.Save(types.ProductRegistration{
			PurchaseDate: v.PurchaseDate,
			ExpiryAt:     v.ExpiryAt,
			Product:      v.Product,
			SerialCode:   v.SerialCode,
			RootId:       &rootId,
			ProfileId:    &profileId,
			ParentId:     &parentId,
		})
		saveProductRegistrationsProductChildren(v.AdditionalProductRegistrations, profileId, rootId, newId)
	}

}
