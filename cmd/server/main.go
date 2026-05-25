package main

import (
	"log"

	"github.com/berntau/curriculo-terminal/internal/curriculo"
	"github.com/gin-gonic/gin"
)

func main() {
	config := curriculo.NewConfig()

	service, err := curriculo.NewResumeService(config.DataPath)
	if err != nil {
		log.Fatalf("falha ao iniciar serviço: %v", err)
	}

	handler := curriculo.NewResumeHandler(service)

	r := gin.Default()

	// main.go agora só monta — não conhece as rotas
	curriculo.RegistrarRotas(r, handler)

	log.Printf("Servidor rodando em http://localhost:%s", config.Port)
	r.Run(":" + config.Port)
}
