package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ldap-auth/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/health", controllers.HealthCheck)

	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/validate", controllers.ValidateTokenHandler)
	return r
}
