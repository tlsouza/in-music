package controllers

import (
	"api/app/services"
	"api/pkg/errors"
	pgk_types "api/pkg/ports/types"
	"fmt"
	"strconv"
)

type productRegistrationController struct {
	//Inject product Service
	ps services.IProductRegistrationService
}

func NewProductRegistrationController(ps services.IProductRegistrationService) productRegistrationController {
	return productRegistrationController{
		ps,
	}
}

func (pc *productRegistrationController) GetById(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	id, ok := requestData.PathParams["id"]
	if !ok {
		return nil, errors.NewHttpError(fmt.Errorf("invalid id in path"), 400)
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("int id expected in path"), 400)
	}

	productRegistration, err := pc.ps.GetBundle(uint64(intId))

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("not found"), 404)
	}

	return productRegistration, nil

}

func (pc *productRegistrationController) GetAll(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	productRegistration := pc.ps.GetAll()
	return productRegistration, nil
}
