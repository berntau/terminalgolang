package curriculo

type Config struct {
	Port     string
	Env      string
	DataPath string
}

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		Env:      "development",
		DataPath: "internal/curriculo/data/resume.json",
	}
}
