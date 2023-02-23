package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/go-jwt-project/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Acsess granted for api v1"})
	})

	router.GET("/api/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Acsess granted for api v2"})
	})

	router.Run(":" + port)
}
