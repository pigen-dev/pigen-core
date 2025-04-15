package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pigen-dev/pigen-core/internal/api/handlers"
)

func SetupPluginRoutes(api *gin.RouterGroup) {
	plugin_api := api.Group("/plugin")
  plugin_api.POST("/setup_plugin", handlers.SetupPlugin)
	plugin_api.POST("/destroy_plugin", handlers.DestroyPlugin)
	plugin_api.GET("/get_output", handlers.GetOutput)
}