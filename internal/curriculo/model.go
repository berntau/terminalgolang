package curriculo

import "time"

// Em Go, campos com letra MAIÚSCULA são exportados (públicos)
// campos com letra minúscula são privados ao pacote
type Resume struct {
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	Location  string    `json:"location"`
	Email     string    `json:"email"`
	Github    string    `json:"github"`
	CreatedAt time.Time `json:"created_at"`
}

// Sem banco de dados por enquanto — dados em memória
// Isso é uma função construtora por convenção em Go (New + NomeDaStruct)
func NewResume() Resume {
	return Resume{
		Name:      "Tauã Bernardo",
		Bio:       "Desenvolvedor de software com experiência em Go e Java.",
		Location:  "São Paulo, Brasil",
		Email:     "taua.bernardo@example.com",
		Github:    "https://github.com/taua-bernardo",
		CreatedAt: time.Now(),
	}
}
