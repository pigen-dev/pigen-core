package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pigen-dev/pigen-core/internal/api/handlers"
)

func SetupCICDRoutes(api *gin.RouterGroup) {
	cicd_api := api.Group("/cicd")
  cicd_api.POST("/connect_repo", handlers.ConnectRepo)
	cicd_api.POST("/create_trigger", handlers.CreateTrigger)
	cicd_api.POST("/pipeline_notifier", handlers.PipelineNotifier)
	cicd_api.POST("/gen_script", handlers.GenerateScript)
}