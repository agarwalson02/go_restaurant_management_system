package main

import (
	"os"
	middleware "github.com/agarwalson02/go_restaurant_management_system/middleware"
	"github.com/agarwalson02/go_restaurant_management_system/database"
	"github.com/agarwalson02/go_restaurant_management_system/helpers"
	"github.com/agarwalson02/go_restaurant_management_system/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = helpers.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}
