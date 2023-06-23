package application

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/response"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/services"
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
	service services.Products
}

func NewDeactivateProductUseCase(logger ports.Logger, products services.Products) *DeactivateProductUseCase {
	return &DeactivateProductUseCase{
		logger:  logger,
		service: products,
	}
}
