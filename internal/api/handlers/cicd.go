package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pigen-dev/pigen-core/internal/core/cicd"
	shared "github.com/pigen-dev/shared"
)

func ConnectRepo(c *gin.Context){
	var pigenStepsFile shared.PigenStepsFile
	err := c.ShouldBindBodyWithJSON(&pigenStepsFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid connecting repo request", "error": err.Error()})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to connect repo", "error": r})
		}
	}()
	resp := cicd.ConnectRepo(pigenStepsFile)
	if resp.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to connect repo", "error": resp.Error.Error()})
		return
	}
	if resp.ActionUrl != "" {
		c.JSON(http.StatusOK, gin.H{"Action is required": resp.ActionUrl})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repo connected successfully"})
}

func CreateTrigger(c *gin.Context){
	var pigenStepsFile shared.PigenStepsFile
	err := c.ShouldBindBodyWithJSON(&pigenStepsFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid trigger creation request", "error": err.Error()})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create trigger", "error": r})
		}
	}()
	err = cicd.CreateTrigger(pigenStepsFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create trigger", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repo connected successfully"})
}

func PipelineNotifier(c *gin.Context){
	var pipelineNotification shared.PipelineNotification
	err := c.ShouldBindBodyWithJSON(&pipelineNotification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid Notification request", "error": err.Error()})
		return
	}
	err = cicd.PipelineNotifier(pipelineNotification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to handle pipeline notification", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification handled successfully"})
}