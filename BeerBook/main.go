package main

import "BeerBook/internal/routes"

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
