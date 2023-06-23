package ports

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/response"
)

type Products interface {
	DeactivateProduct(productDto request.DeactivateProductDto) (error, response.DeactivateProductDto)
}
