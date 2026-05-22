package curriculo

// ResumeService é responsável por lidar com a lógica de negócios relacionada ao currículo
// Em Go é comum usar structs para organizar serviços e suas dependências
type ResumeService struct {
	resume Resume
}

// NewResumeService é o "construtor" do service
func NewResumeService() *ResumeService {
	return &ResumeService{
		resume: NewResume(),
	}
}

// GetResume retorna o currículo armazenado no serviço
// Note: método em Go é uma função com "receiver" — o (s *ResumeService) antes do nome
// É assim que Go faz "orientação a objetos" sem classes
func (s *ResumeService) GetResume() Resume {
	return s.resume
}
