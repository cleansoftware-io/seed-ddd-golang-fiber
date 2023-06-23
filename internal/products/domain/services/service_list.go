package services

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
)

type ServiceList struct {
	DeactivateProduct ports.Products
}

func ProvideServiceList(deactivateProduct adapters.DeactivateProductsIml) *ServiceList {
	return &ServiceList{
		DeactivateProduct: deactivateProduct,
	}
}
