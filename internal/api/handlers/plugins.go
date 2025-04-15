package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pigen-dev/pigen-core/internal/core/plugins"
	shared "github.com/pigen-dev/shared"
)

func SetupPlugin(c *gin.Context){
	var pluginStruct shared.PluginStruct
	err := c.ShouldBindBodyWithJSON(&pluginStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid setup plugin request", "error": err.Error()})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to setup plugin", "error": r})
		}
	}()
	err = plugins.SetupPlugins(pluginStruct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to setup plugin", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Plugins setup complete"})
}

func DestroyPlugin(c *gin.Context){
	var pluginStruct shared.PluginStruct
	err := c.ShouldBindBodyWithJSON(&pluginStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid destroy plugin request", "error": err.Error()})
		return
	}
	err = plugins.DestroyPlugin(pluginStruct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to destroy plugin", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Plugins destructed successfully"})
}

func GetOutput(c *gin.Context){
	var pluginStruct shared.PluginStruct
	err := c.ShouldBindBodyWithJSON(&pluginStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid destroy plugin request", "error": err.Error()})
		return
	}
	output := plugins.GetOutput(pluginStruct)
	if output.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get plugin output", "error": output.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}