package routes

import (
	"github.com/gin-gonic/gin"
    "BeerBook/internal/controllers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/beers", controllers.GetBeers)
	router.GET("/beers/:id", controllers.GetBeer)

	return router
}
