package main

import (
	"os"

	"github.com/agarwalson02/go_restaurant_management_system/routes"
	"github.com/gin-gonic/gin"
	"go.mongogb.org/mongo-driver/mongo"
)

func main(){
	port:=os.Getenv("PORT")
	if port==""{
		port="8000"
	}

	router:=gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
}