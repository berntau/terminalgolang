package curriculo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResumeHTTPHandler representa a camada HTTP dessa funcionalidade.
// A ideia do handler é receber a requisição, conversar com o service
// e devolver uma resposta HTTP para o cliente.
type ResumeHTTPHandler struct {
	// Guardamos um ponteiro para o service para reutilizar a mesma instância
	// sem copiar a struct inteira.
	service *ResumeService
}

// NewResumeHTTPHandler é uma função construtora por convenção.
// Ela recebe as dependências do handler e devolve um ponteiro para ele.
func NewResumeHTTPHandler(service *ResumeService) *ResumeHTTPHandler {
	return &ResumeHTTPHandler{
		service: service,
	}
}

// GetResume é o handler para a rota GET /resume
// Ele chama o serviço para obter o currículo e retorna como JSON
// O (h *ResumeHTTPHandler) antes do nome é o "receiver" do método:
// isso significa que GetResume pertence à struct ResumeHTTPHandler.
// O gin.Context carrega os dados da requisição e também ajuda
// a montar a resposta que será enviada ao cliente.
func (h *ResumeHTTPHandler) GetResume(c *gin.Context) {
	// O Gin já separa handlers por método HTTP quando usamos r.GET(...),
	// então essa checagem é redundante neste caso.
	// Ainda assim, ela serve como exemplo de como acessar a requisição bruta.
	if c.Request.Method != http.MethodGet {
		http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Aqui o handler delega a regra de negócio para o service.
	resume := h.service.GetResume()

	// c.JSON envia a resposta com status HTTP 200 e converte a struct
	// para JSON usando as tags `json:"..."` definidas no model.
	c.JSON(http.StatusOK, resume)
}
