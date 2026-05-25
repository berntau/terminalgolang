package main

import (
	"log"

	"github.com/berntau/curriculo-terminal/internal/curriculo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := curriculo.NewConfig()

	// dependencias
	service := curriculo.NewResumeService()
	handler := curriculo.NewResumeHTTPHandler(service)

	//rotas
	r.GET("/resume", handler.GetResume)

	log.Printf("Servidor rodando http://localhost:%s", config.Port)
	r.Run(":" + config.Port)

}
