package server

type Config struct {
	Port         int `json:"port"`
	WriteTimeout int `json:"write_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	IdleTimeout  int `json:"idle_timeout"`
}

func NewServerConfig() *Config {
	return &Config{
		Port:         8080,
		WriteTimeout: 15,
		ReadTimeout:  15,
		IdleTimeout:  15,
	}
}
