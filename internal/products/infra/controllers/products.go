package controllers

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/application"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/dto/request"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"github.com/gofiber/fiber/v2"
)

func NewProducts(app *fiber.App, logger ports.Logger, deactivateProducts *application.DeactivateProductUseCase) *Products {
	return &Products{
		app:                app,
		logger:             logger,
		deactivateProducts: deactivateProducts,
	}
}

func (p *Products) RegisterRoutes() {
	p.logger.Info("RegisterRoutes")
	p.app.Delete("/products", p.Deactivate)
}

func (p *Products) Deactivate(context *fiber.Ctx) error {
	p.logger.Info("Deactivate controller")
	requestDto := request.DeactivateProductDto{}
	if err := context.BodyParser(&requestDto); err != nil {
		p.logger.Error(err)
		return err
	}
	err, responseDto := p.deactivateProducts.DeactivateProduct(requestDto)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	return context.JSON(responseDto)
}

type Products struct {
	app                *fiber.App
	logger             ports.Logger
	deactivateProducts *application.DeactivateProductUseCase
}
