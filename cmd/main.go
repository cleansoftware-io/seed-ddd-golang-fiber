package main

import (
	"github.con/tgarcia/seed-golang-server/cmd/products"
	"github.con/tgarcia/seed-golang-server/cmd/users"
)

func main() {
	// Configurar la inyección de dependencias y otros aspectos globales de la aplicación

	// Inicializar el contexto de usuarios
	users.Run()

	// Inicializar el contexto de productos
	products.Run()

	// Otras inicializaciones o lógica de la aplicación
}
