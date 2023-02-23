package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meshachdamilare/go-jwt-project/controllers"
	"github.com/meshachdamilare/go-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controller.GetAllUsers())
	incomingRoutes.GET("/users/:userId", controller.GetUser())
}
