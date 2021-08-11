package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(port int) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "Welcome to episode-parser v1.0.0",
			"server_time": time.Now().String(),
		})
	})
	router.POST("/parse", JSONParser())

	listenPort := fmt.Sprintf(":%d", port)

	return &http.Server{
		Addr:    listenPort,
		Handler: router,
	}
}
