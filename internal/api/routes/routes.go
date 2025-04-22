package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()
	//To use a global middelware router.Use(middlewares.CORSMiddleware())
	// CORS configuration to allow all origins
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//API routes
	api:= router.Group("api/v1")
	SetupPluginRoutes(api)
	SetupCICDRoutes(api)
	return router
}