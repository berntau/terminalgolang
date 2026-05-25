package curriculo

import (
	"fmt"
	"strings"
)

type CommandService struct {
	commands map[string]Command // map — equivalente ao HashMap<String, Command> do Java
	resume   Resume
}

func NewCommandService(resume Resume) *CommandService {
	cs := &CommandService{
		commands: make(map[string]Command),
		resume:   resume,
	}

	// Registra todos os comandos disponíveis
	cs.register(Command{
		Name:        "help",
		Description: "lista todos os comandos disponíveis",
		Handler:     cs.handleHelp,
	})

	cs.register(Command{
		Name:        "about",
		Description: "exibe informações sobre mim",
		Handler:     cs.handleAbout,
	})

	cs.register(Command{
		Name:        "skills",
		Description: "lista minhas habilidades",
		Handler:     cs.handleSkills,
	})

	cs.register(Command{
		Name:        "experience",
		Description: "mostra minha experiência profissional",
		Handler:     cs.handleExperience,
	})

	cs.register(Command{
		Name:        "education",
		Description: "mostra minha formação acadêmica",
		Handler:     cs.handleEducation,
	})

	cs.register(Command{
		Name:        "contact",
		Description: "exibe formas de contato",
		Handler:     cs.handleContact,
	})

	return cs
}

// Execute processa o input do usuário e retorna a resposta
func (cs *CommandService) Execute(input string) string {
	// strings.TrimSpace remove espaços — strings.ToLower normaliza
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "" {
		return "digite 'help' para ver os comandos disponíveis"
	}

	// Busca o comando no map
	command, exists := cs.commands[input]
	if !exists {
		return fmt.Sprintf("comando '%s' não encontrado. Digite 'help' para ver os comandos disponíveis", input)
	}

	// Chama a função guardada na struct
	return command.Handler()
}

// register adiciona um comando ao map
func (cs *CommandService) register(cmd Command) {
	cs.commands[cmd.Name] = cmd
}

// GetCommands retorna todos os comandos — usado pelo handler
func (cs *CommandService) GetCommands() map[string]Command {
	return cs.commands
}

// --- Handlers de cada comando ---

func (cs *CommandService) handleHelp() string {
	var sb strings.Builder

	sb.WriteString("comandos disponíveis:\n\n")

	for name, cmd := range cs.commands {
		sb.WriteString(fmt.Sprintf("  %-12s %s\n", name, cmd.Description))
	}

	return sb.String()
}

func (cs *CommandService) handleAbout() string {
	return fmt.Sprintf(
		"nome:      %s\nbio:       %s\nlocalização: %s",
		cs.resume.Name,
		cs.resume.Bio,
		cs.resume.Location,
	)
}

func (cs *CommandService) handleSkills() string {
	var sb strings.Builder

	sb.WriteString("habilidades:\n\n")

	for _, skill := range cs.resume.Skills {
		sb.WriteString(fmt.Sprintf("  %s\n", skill.Category))
		for _, item := range skill.Items {
			sb.WriteString(fmt.Sprintf("    → %s\n", item))
		}
	}

	return sb.String()
}

func (cs *CommandService) handleExperience() string {
	var sb strings.Builder

	sb.WriteString("experiência profissional:\n\n")

	for _, exp := range cs.resume.Experience {
		status := "atual"
		if !exp.Current && exp.EndDate != "" {
			status = exp.EndDate
		}

		sb.WriteString(fmt.Sprintf(
			"  %s — %s\n  %s até %s\n  %s\n\n",
			exp.Company,
			exp.Role,
			exp.StartDate,
			status,
			exp.Description,
		))
	}

	return sb.String()
}

func (cs *CommandService) handleEducation() string {
	var sb strings.Builder

	sb.WriteString("formação acadêmica:\n\n")

	for _, edu := range cs.resume.Education {
		sb.WriteString(fmt.Sprintf(
			"  %s\n  %s em %s (%d)\n\n",
			edu.Institution,
			edu.Degree,
			edu.Field,
			edu.Year,
		))
	}

	return sb.String()
}

func (cs *CommandService) handleContact() string {
	return fmt.Sprintf(
		"contato:\n\n  email:    %s\n  github:   %s\n  linkedin: %s",
		cs.resume.Contact.Email,
		cs.resume.Contact.Github,
		cs.resume.Contact.Linkedin,
	)
}
