package adapters

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/response"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
)

type DeactivateProductsIml struct {
	logger ports.Logger
}

func ProvideDeactivateProductsIml(logger ports.Logger) DeactivateProductsIml {
	return DeactivateProductsIml{
		logger: logger,
	}
}

func (u DeactivateProductsIml) DeactivateProduct(productDto request.DeactivateProductDto) (error, response.DeactivateProductDto) {
	u.logger.Info("DeactivateProduct executed")
	return nil, response.DeactivateProductDto{
		ID: "a uuid from product",
	}
}
