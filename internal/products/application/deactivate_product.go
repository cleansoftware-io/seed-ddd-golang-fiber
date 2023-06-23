package application

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/response"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
)

func (u DeactivateProductUseCase) DeactivateProduct(productDto request.DeactivateProductDto) (error, *response.DeactivateProductDto) {
	u.logger.Info("DeactivateProductUseCase")
	u.logger.Info(productDto.ID)
	err, responseDto := u.service.DeactivateProduct(productDto)
	if err != nil {
		return err, nil
	}
	return nil, &responseDto

}

type DeactivateProductUseCase struct {
	logger  ports.Logger
	service ports.Products
}

func NewDeactivateProductUseCase(logger ports.Logger, products ports.Products) *DeactivateProductUseCase {
	return &DeactivateProductUseCase{
		logger:  logger,
		service: products,
	}
}
