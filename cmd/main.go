package main

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/products"
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/users"
)

func main() {
	bootstrap := InitializeApplication()
	products.Run(bootstrap)
	users.Run(bootstrap)
	if err := bootstrap.App.Listen(":3000"); err != nil {
		bootstrap.Logger.Fatal(err)
		panic(err)
	}
}
