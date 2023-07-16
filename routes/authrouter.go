package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/puru/golang-jwt-project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/Login", controller.Login())

}
