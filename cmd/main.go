package main

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/initialization"
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/products"
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/users"
)

func main() {
	// Configurar la inyección de dependencias y otros aspectos globales de la aplicación
	bootstrap, err := initialization.NewBootstrap()
	if err != nil {
		panic(err)
	}

	// Inicializar el contexto de usuarios
	users.Run(bootstrap)

	// Inicializar el contexto de productos
	products.Run(bootstrap)

	// Otras inicializaciones o lógica de la aplicación

	if err := bootstrap.App.Listen(":3000"); err != nil {
		bootstrap.Logger.Fatal(err)
		panic(err)
	}
}
