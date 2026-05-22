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
