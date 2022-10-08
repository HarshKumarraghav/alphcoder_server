package main

import (
	"os"
	"time"

	routes "example.com/alphacoder-server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowHeaders:     []string{"Accept, Content-Type, Content-Length, Token"},
			AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})
	gin.SetMode(gin.ReleaseMode)
	router.Run(":" + port)
}
