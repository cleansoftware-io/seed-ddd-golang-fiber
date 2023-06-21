package services

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/request"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/domain/dto/response"
)

type Products interface {
	DeactivateProduct(productDto request.DeactivateProductDto) (error, response.DeactivateProductDto)
}
