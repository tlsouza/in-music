package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

var newProduct, getProductBySku types.HttpServerPort

func init() {
	productController := controllers.NewProductController(
		services.NewProductService(
			repository.GetProductRepositoryInstance(),
		),
	)

	newProduct = types.HttpServerPort{
		SilentRoute: true,
		Name:        "products",
		Path:        "products",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  productController.CreateNewProduct,
	}

	getProductBySku = types.HttpServerPort{
		SilentRoute: true,
		Name:        "products",
		Path:        "products/:sku",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productController.GetBySku,
	}

	newProduct.Start()
	getProductBySku.Start()
}
