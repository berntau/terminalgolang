package main

import (
	"net/http"

	"github.com/berntau/curriculo-terminal/internal/curriculo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handler := curriculo.NewResumeHTTPHandler(curriculo.NewResumeService())
	r.GET("/api/resume", handler.GetResume)

	r.Run()
}
