package services

import (
	"api/app/repository"
	"api/app/types"
	"api/pkg/errors"
	"fmt"
)

type productRegistrationService struct {
	repo repository.IProductRegistrationRepository
}

func NewProductRegistrationService(repo repository.IProductRegistrationRepository) *productRegistrationService {
	return &productRegistrationService{
		repo,
	}
}

func (ps *productRegistrationService) GetBundle(id uint64) (*types.ProductRegistrationHttpRes, error) {
	root, chilren := ps.repo.GetByBundle(id)
	if root == nil {
		return nil, errors.NewHttpError(fmt.Errorf("product rgistration not found: %d", id), 404)
	}
	res := types.ProductRegistrationHttpRes{
		Id:                             root.Id,
		ExpiryAt:                       root.ExpiryAt,
		PurchaseDate:                   root.PurchaseDate,
		SerialCode:                     root.SerialCode,
		Product:                        root.Product,
		AdditionalProductRegistrations: []types.ProductRegistrationHttpResChild{},
	}

	for _, v := range chilren {
		res.AdditionalProductRegistrations = append(res.AdditionalProductRegistrations, types.ProductRegistrationHttpResChild{
			Id:           v.Id,
			ExpiryAt:     v.ExpiryAt,
			PurchaseDate: v.PurchaseDate,
			SerialCode:   v.SerialCode,
			Product:      v.Product,
		})
	}

	return &res, nil
}

func (ps *productRegistrationService) GetAll() []types.ProductRegistration {
	return ps.repo.GetAll()
}
