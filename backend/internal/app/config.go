package app

import (
	"chatServer/internal/logger"
	"chatServer/internal/server"
	"chatServer/pkg/db/sqlstore"
)

type Config struct {
	Server *server.Config
	Store  *sqlstore.Config
	Logger *logger.Config
}

func NewConfig() *Config {
	return &Config{
		Server: server.NewServerConfig(),
		Store:  sqlstore.NewConfig(),
	}
}
