package curriculo_test

import (
	"strings"
	"testing"

	"github.com/berntau/curriculo-terminal/internal/curriculo"
)

func TestGetResume_LoadsResumeFromJSON(t *testing.T) {
	service, err := curriculo.NewResumeService("data/resume.json")
	if err != nil {
		t.Fatalf("falha ao criar service: %v", err)
	}

	resume := service.GetResume()

	if resume.Name != "Bernardo" {
		t.Fatalf("nome inesperado: %q", resume.Name)
	}

	if !strings.Contains(resume.Bio, "Desenvolvedor Júnior na CODATA") {
		t.Fatalf("bio inesperada: %q", resume.Bio)
	}

	if resume.Location != "João Pessoa, PB, Brazil" {
		t.Fatalf("localizacao inesperada: %q", resume.Location)
	}

	if resume.Contact.Github != "github.com/berntau" {
		t.Fatalf("github inesperado: %q", resume.Contact.Github)
	}

	if resume.Contact.Linkedin != "linkedin.com/in/tauabernardo" {
		t.Fatalf("linkedin inesperado: %q", resume.Contact.Linkedin)
	}

	if len(resume.Skills) != 3 {
		t.Fatalf("quantidade de skills inesperada: %d", len(resume.Skills))
	}

	if resume.Skills[0].Category != "Java" {
		t.Fatalf("categoria da primeira skill inesperada: %q", resume.Skills[0].Category)
	}

	if len(resume.Skills[0].Items) == 0 || resume.Skills[0].Items[0] != "Java" {
		t.Fatalf("items da primeira skill inesperados: %#v", resume.Skills[0].Items)
	}

	if len(resume.Experience) != 1 {
		t.Fatalf("quantidade de experiences inesperada: %d", len(resume.Experience))
	}

	if resume.Experience[0].Company != "Companhia de Processamento de Dados da Paraíba - CODATA" {
		t.Fatalf("empresa inesperada: %q", resume.Experience[0].Company)
	}

	if resume.Experience[0].Role != "Desenvolvedor Júnior" {
		t.Fatalf("cargo inesperado: %q", resume.Experience[0].Role)
	}

	if !resume.Experience[0].Current {
		t.Fatal("esperava experience atual marcada como current=true")
	}

	if len(resume.Education) != 0 {
		t.Fatalf("education deveria estar vazia, veio com %d itens", len(resume.Education))
	}
}

func TestNewResumeService_ReturnsErrorWhenFileDoesNotExist(t *testing.T) {
	_, err := curriculo.NewResumeService("data/arquivo-inexistente.json")
	if err == nil {
		t.Fatal("esperava erro ao carregar arquivo inexistente")
	}
}
