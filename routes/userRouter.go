package routes

import (
	controller "github.com/Christomesh/go-jwt-project/controllers"
	"github.com/Christomesh/go-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("users/", controller.GetAllUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
