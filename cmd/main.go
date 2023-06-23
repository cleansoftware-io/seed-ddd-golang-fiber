package main

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/products"
	"cleansoftware.io/ddd/fiber/seed/cmd/users"
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
