package services

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
)

type ServiceList struct {
	DeactivateProduct ports.Products
	//... inject other services
}

func ProvideServiceList(deactivateProduct adapters.DeactivateProductsIml) ServiceList {
	return ServiceList{
		DeactivateProduct: deactivateProduct,
		//... inject other services
	}
}
