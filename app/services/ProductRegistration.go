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

func (ps *productRegistrationService) GetByID(id uint64) (*types.ProductRegistrationHttpRes, error) {
	parent, err := ps.repo.GetByID(id)
	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("product not found: %d", id), 404)
	}

	child := ps.repo.GetByParentId(id)

	return mapProductRegistrationToHttpResponse(*parent, child), nil
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

func mapProductRegistrationToHttpResponse(parent types.ProductRegistration, children []types.ProductRegistration) *types.ProductRegistrationHttpRes {
	res := types.ProductRegistrationHttpRes{
		Id:                             parent.Id,
		PurchaseDate:                   parent.PurchaseDate,
		ExpiryAt:                       parent.ExpiryAt,
		Product:                        parent.Product,
		SerialCode:                     parent.SerialCode,
		AdditionalProductRegistrations: []types.ProductRegistrationHttpResChild{},
	}
	for _, v := range children {
		res.AdditionalProductRegistrations = append(res.AdditionalProductRegistrations,
			types.ProductRegistrationHttpResChild{
				Id:           v.Id,
				PurchaseDate: v.PurchaseDate,
				ExpiryAt:     v.ExpiryAt,
				Product:      v.Product,
				SerialCode:   v.SerialCode,
			})
	}

	return &res
}
