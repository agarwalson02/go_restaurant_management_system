package routes

import (
	"github.coim/gin-gonic/gin"
	controller "github.com/agarwalson02/go_restaurant_management_system/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
