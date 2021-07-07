package sqlstore

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     5432,
		User:     "Postgres",
		Password: "Postgres",
		DBName:   "URLShortener",
	}
}
