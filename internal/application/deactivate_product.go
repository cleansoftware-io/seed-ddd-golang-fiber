package application

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/request"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/response"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/ports"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/services"
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
