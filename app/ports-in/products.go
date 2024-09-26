package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
)

func init() {
	productController := controllers.NewProductController(
		services.NewProductService(
			repository.NewInMemoryProductRepository(),
		),
	)

	var newProduct = types.HttpServerPort{
		SilentRoute: true,
		Name:        "products",
		Path:        "products",
		Verb:        types.POST,
		Adapter:     adapters.Fiber,
		Controller:  productController.CreateNewProduct,
	}

	var GetProductBySku = types.HttpServerPort{
		SilentRoute: true,
		Name:        "products",
		Path:        "products/:sku",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productController.GetBySku,
	}

	var GetAllProducts = types.HttpServerPort{
		SilentRoute: true,
		Name:        "products",
		Path:        "products",
		Verb:        types.GET,
		Adapter:     adapters.Fiber,
		Controller:  productController.GetAll,
	}

	newProduct.Start()
	GetProductBySku.Start()
	GetAllProducts.Start()
}
