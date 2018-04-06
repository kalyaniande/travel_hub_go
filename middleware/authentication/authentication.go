package authentication

import (
	"azri_hub/helper"
	_ "fmt"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		validate(c)
		//c.Next()
	}
}

func validate(c *gin.Context) {
	api_key := c.Request.Header.Get("api-key")
	agent_info := helper.GetAgentInfo(api_key)

	if agent_info.APIKey == api_key {
		c.Set("agent_details", agent_info)
		ValidateHeaders(c)
	} else {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		c.AbortWithStatus(401)
	}
}

func ValidateHeaders(c *gin.Context) {
	content_type := c.GetHeader("Content-Type")
	accept_type := c.GetHeader("Accept")
	if content_type != "application/json" {
		c.JSON(400, gin.H{"error": "Unsupported media type"})
		c.AbortWithStatus(400)
	}
	if accept_type != "application/json" {
		c.JSON(400, gin.H{"error": "Response content type is not acceptable according to the request accept headers"})
		c.AbortWithStatus(400)
	}
	c.Next()
}
