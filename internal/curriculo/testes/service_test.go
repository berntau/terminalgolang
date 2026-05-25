package testes

import (
	"testing"

	"github.com/berntau/curriculo-terminal/internal/curriculo"
)

func TestGetResume(t *testing.T) {
	service := curriculo.NewResumeService()
	resume := service.GetResume()

	if resume.Name == "" {
		t.Errorf("esperava um nome no currículo, mas estava vazio")
	}

	if resume.Email == "" {
		t.Errorf("esperava um email no currículo, mas estava vazio")
	}
}
