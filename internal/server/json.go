package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

func JSONParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":      "parsing",
			"server_time": time.Now().String(),
		})
	}
}
