package main

import (
	"github.con/tgarcia/seed-golang-server/cmd/products"
	"github.con/tgarcia/seed-golang-server/cmd/users"
	"github.con/tgarcia/seed-golang-server/internal"
)

func main() {
	// Configurar la inyección de dependencias y otros aspectos globales de la aplicación
	bootstrap, err := internal.NewBootstrap()
	if err != nil {
		panic(err)
	}

	// Inicializar el contexto de usuarios
	users.Run(bootstrap)

	// Inicializar el contexto de productos
	products.Run(bootstrap)

	// Otras inicializaciones o lógica de la aplicación
}
