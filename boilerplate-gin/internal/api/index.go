package api

import (
	"github.com/gin-gonic/gin"
)

func (m *ServiceServer) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to my channel",
	})
}
