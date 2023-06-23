package services

import (
	"cleansoftware.io/ddd/fiber/seed/internal/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/domain/dto/response"
)

type Products interface {
	DeactivateProduct(productDto request.DeactivateProductDto) (error, response.DeactivateProductDto)
}
