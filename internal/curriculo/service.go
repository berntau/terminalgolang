package curriculo

import (
	"encoding/json"
	"fmt"
	"os"
)

type ResumeService struct {
	resume Resume
}

func NewResumeService(dataPath string) (*ResumeService, error) {
	resume, err := loadFromJSON(dataPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar currículo: %w", err)
	}
	return &ResumeService{resume: resume}, nil
}

func (s *ResumeService) GetResume() Resume {
	return s.resume
}

// Métodos específicos por recurso
func (s *ResumeService) GetSkills() []Skill {
	return s.resume.Skills
}

func (s *ResumeService) GetExperience() []Experience {
	return s.resume.Experience
}

func (s *ResumeService) GetEducation() []Education {
	return s.resume.Education
}

func (s *ResumeService) GetContact() Contact {
	return s.resume.Contact
}

func loadFromJSON(path string) (Resume, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Resume{}, fmt.Errorf("erro ao ler arquivo '%s': %w", path, err)
	}

	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return Resume{}, fmt.Errorf("erro ao parsear JSON: %w", err)
	}

	return resume, nil
}
