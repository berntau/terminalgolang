package curriculo

type Contact struct {
	Email    string `json:"email"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
}

type Skill struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

type Experience struct {
	Company     string `json:"company"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Current     bool   `json:"current"`
}

type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	Year        int    `json:"year"`
}

type Resume struct {
	Name       string       `json:"name"`
	Bio        string       `json:"bio"`
	Location   string       `json:"location"`
	Contact    Contact      `json:"contact"`
	Skills     []Skill      `json:"skills"`
	Experience []Experience `json:"experience"`
	Education  []Education  `json:"education"`
}

// Command representa um comando do terminal
// Handler é uma função guardada como valor — não existe isso em Java sem interface
type Command struct {
	Name        string
	Description string
	Handler     func() string // função sem parâmetros que retorna string
}

type CommandRequest struct {
	Command string `json:"command" binding:"required"`
}

type CommandResponse struct {
	Output string `json:"output"`
}
