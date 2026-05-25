package curriculo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrarRotas(r *gin.Engine, handler *ResumeHandler) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.GET("/resume", handler.GetResume)
		api.GET("/skills", handler.GetSkills)
		api.GET("/experience", handler.GetExperience)
		api.GET("/education", handler.GetEducation)
		api.GET("/contact", handler.GetContact)
		api.POST("/command", handler.ExecuteCommand)
	}
}
