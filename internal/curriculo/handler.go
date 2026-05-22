package curriculo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	service *ResumeService
}

func NewResumeHandler(service *ResumeService) *ResumeHandler {
	return &ResumeHandler{
		service: service,
	}
}

// GetResume é o handler para a rota GET /resume
// Ele chama o serviço para obter o currículo e retorna como JSON
func (h *ResumeHandler) GetResume(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resume := h.service.GetResume()
	c.JSON(http.StatusOK, resume)
}
