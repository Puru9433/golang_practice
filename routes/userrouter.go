package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/puru/golang-jwt-project/controllers"
	"github.com/puru/golang-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
