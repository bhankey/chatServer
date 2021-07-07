package logger

type Config struct {
	LogLevel string `yaml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "info",
	}
}
