package main

func main() {
	initializeApplication := InitializeApplication()

	//products.Run(bootstrap)
	//users.Run(bootstrap)
	if err := initializeApplication.Bootstrap.App.Listen(":3000"); err != nil {
		initializeApplication.Bootstrap.Logger.Fatal(err)
		panic(err)
	}
}
