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

func (ps *profileService) AddProductRegistration(profileId uint64, productRegistration types.ProductRegistrationHttpRequest) ([]uint64, error) {
	// validate user
	if _, err := ps.repo.GetByID(profileId); err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("user not found"), 404)
	}

	//validate root Product
	productRepo := repository.GetProductRepositoryInstance()
	if product, err := productRepo.GetBySku(*productRegistration.Product.SKU); err != nil || product == nil {
		return nil, errors.NewHttpError(fmt.Errorf("product %s not found", *productRegistration.Product.SKU), 404)
	}

	//build Root Product Registration to save after children products validation
	rootProductRegistration := types.ProductRegistration{
		PurchaseDate: productRegistration.PurchaseDate,
		ExpiryAt:     productRegistration.ExpiryAt,
		Product:      productRegistration.Product,
		SerialCode:   productRegistration.SerialCode,
		ProfileId:    &profileId,
		ParentId:     nil,
	}

	childrenProductRegistration := []types.ProductRegistration{}

	//validate children products
	for _, v := range productRegistration.AdditionalProductRegistrations {
		if product, err := productRepo.GetBySku(*v.Product.SKU); err != nil || product == nil {
			return nil, errors.NewHttpError(fmt.Errorf("product %s not found", *v.Product.SKU), 404)
		}

		//build child Product Registration to save after other children products validation
		childrenProductRegistration = append(childrenProductRegistration,
			types.ProductRegistration{
				PurchaseDate: v.PurchaseDate,
				ExpiryAt:     v.ExpiryAt,
				Product:      v.Product,
				SerialCode:   v.SerialCode,
				ProfileId:    nil,
			})
	}
	productRegistrationRpo := repository.GetProductRegistrationRepositoryInstance()
	id, _ := productRegistrationRpo.Save(rootProductRegistration)
	childrenIds, err := productRegistrationRpo.SaveChildren(id, childrenProductRegistration)
	return append(childrenIds, id), err
}
