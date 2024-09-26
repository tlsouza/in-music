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
		return nil, errors.NewHttpError(fmt.Errorf("unable to save Profile"), 500)
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

	profile, err := pc.ps.GetByID(intId)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("profile not fo0und"), 404)
	}

	return profile, nil

}

func (pc *profileController) GetAll(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	profiles := pc.ps.GetAll()
	return profiles, nil
}
