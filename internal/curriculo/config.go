package curriculo

type Config struct {
	Port string
	Env  string
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
		Env:  "development",
	}
}
