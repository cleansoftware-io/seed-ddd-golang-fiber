package adapters

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/request"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/response"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/ports"
)

type DeactivateProductsIml struct {
	logger ports.Logger
}

func NewDeactivateProductsIml(logger ports.Logger) *DeactivateProductsIml {
	logger.Info("DeactivateProductsIml")
	return &DeactivateProductsIml{
		logger: logger,
	}
}

func (u DeactivateProductsIml) DeactivateProduct(productDto request.DeactivateProductDto) (error, response.DeactivateProductDto) {
	u.logger.Info("DeactivateProduct executed")
	return nil, response.DeactivateProductDto{
		ID: "a uuid from product",
	}
}
