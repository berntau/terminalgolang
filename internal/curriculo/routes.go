package curriculo

import (
	"net/http"
	"strings"

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

	r.Static("/assets", "./frontend/dist/assets")
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.File("./frontend/dist/index.html")
	})
}
