package curriculo

// ResumeService representa a camada de serviço da funcionalidade de currículo.
// A responsabilidade dessa camada é concentrar a lógica da aplicação,
// deixando o handler focado apenas em HTTP.
type ResumeService struct {
	// Aqui o service mantém em memória os dados do currículo.
	resume Resume
}

// NewResumeService cria e devolve uma nova instância do service.
// Neste projeto, ele já inicializa o campo resume chamando NewResume().
func NewResumeService() *ResumeService {
	return &ResumeService{
		resume: NewResume(),
	}
}

// GetResume devolve o currículo armazenado dentro do service.
// O (s *ResumeService) é o receiver do método, ou seja,
// indica que essa função pertence à struct ResumeService.
// Hoje ela apenas retorna um dado em memória, mas no futuro
// poderia buscar essas informações em banco, arquivo ou API externa.
func (s *ResumeService) GetResume() Resume {
	return s.resume
}
