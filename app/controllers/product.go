package controllers

import (
	"api/app/services"
	"api/app/types"
	"api/pkg/errors"
	"api/pkg/ports/logic"
	pgk_types "api/pkg/ports/types"
	"fmt"
)

type productController struct {
	//Inject product Service
	ps services.IProductService
}

func NewProductController(ps services.IProductService) productController {
	return productController{
		ps,
	}
}

func (pc *productController) CreateNewProduct(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	product, err := logic.Unmarshal[types.Product](requestData.BodyByte, requestData.Ctx)

	if err != nil || product.SKU == nil {
		return nil, errors.NewHttpError(fmt.Errorf("invalid body structure"), 400)
	}

	newproductSku, err := pc.ps.Save(product)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("unable to save product"), 500)
	}

	return newproductSku, nil
}

func (pc *productController) GetBySku(requestData pgk_types.RequestData) (interface{}, *errors.HttpError) {
	sku, ok := requestData.PathParams["sku"]
	if !ok {
		return nil, errors.NewHttpError(fmt.Errorf("invalid id in path"), 400)
	}

	product, err := pc.ps.GetBySku(sku)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("product not found"), 404)
	}

	return product, nil

}
