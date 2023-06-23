package services

import "cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"

type ServiceList struct {
	DeactivateProduct Products
}

func ProvideServiceList(deactivateProduct adapters.DeactivateProductsIml) *ServiceList {
	return &ServiceList{
		DeactivateProduct: deactivateProduct,
	}
}
