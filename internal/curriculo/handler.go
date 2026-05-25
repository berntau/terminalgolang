package curriculo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	service        *ResumeService
	commandService *CommandService
}

func NewResumeHandler(service *ResumeService) *ResumeHandler {
	return &ResumeHandler{
		service:        service,
		commandService: NewCommandService(service.GetResume()),
	}
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetResume())
}

func (h *ResumeHandler) GetSkills(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetSkills())
}

func (h *ResumeHandler) GetExperience(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetExperience())
}

func (h *ResumeHandler) GetEducation(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetEducation())
}

func (h *ResumeHandler) GetContact(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetContact())
}

func (h *ResumeHandler) ExecuteCommand(c *gin.Context) {
	var request CommandRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comando inválido"})
		return
	}

	output := h.commandService.Execute(request.Command)
	c.JSON(http.StatusOK, CommandResponse{Output: output})
}
