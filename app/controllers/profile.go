package controllers

import (
	"api/app/services"
	"api/app/types"
	"api/pkg/errors"
	"api/pkg/ports/logic"
	pgk_types "api/pkg/ports/types"
	"fmt"
	"strconv"
)

type profileController struct {
	//Inject Profile Service
	ps services.IProfileService
}

func NewProfileController(ps services.IProfileService) profileController {
	return profileController{
		ps,
	}
}

func (pc *profileController) CreateNewProfile(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	profileHttpRequest, err := logic.Unmarshal[types.ProfileHttpRequest](requestData.BodyByte, requestData.Ctx)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("invalid body structure"), 400)
	}

	newProfileId, err := pc.ps.Save(profileHttpRequest)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("unable to save profile"), 500)
	}

	return newProfileId, nil
}

func (pc *profileController) GetById(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	id, ok := requestData.PathParams["id"]
	if !ok {
		return nil, errors.NewHttpError(fmt.Errorf("invalid id in path"), 400)
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("int id expected in path"), 400)
	}

	profile, err := pc.ps.GetByID(uint64(intId))

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("profile not found"), 404)
	}

	return profile, nil

}

func (pc *profileController) GetAll(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	profiles := pc.ps.GetAll()
	return profiles, nil
}

func (pc *profileController) CreateNewProductRegistration(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	id, ok := requestData.PathParams["profile"]
	if !ok {
		return nil, errors.NewHttpError(fmt.Errorf("invalid id in path"), 400)
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("int id expected in path"), 400)
	}

	productRegistrationRequest, err := logic.Unmarshal[types.ProductRegistrationHttpReq](requestData.BodyByte, requestData.Ctx)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("invalid body structure"), 400)
	}

	ids, err := pc.ps.AddProductRegistration((uint64(intId)), productRegistrationRequest)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf(err.Error()), 404)
	}

	return ids, nil
}

func (pc *profileController) GetProductRegistrationByProfileId(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	id, ok := requestData.PathParams["profile"]
	if !ok {
		return nil, errors.NewHttpError(fmt.Errorf("invalid id in path"), 400)
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("int id expected in path"), 400)
	}

	prs, err := pc.ps.GetProductRegistrationByProfileId(uint64(intId))

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf(err.Error()), 404)
	}

	return prs, nil
}
