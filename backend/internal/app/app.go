package app

import (
	"chatServer/internal/handler"
	"chatServer/internal/logger"
	"chatServer/internal/repository"
	"chatServer/internal/server"
	"chatServer/internal/service"
	"chatServer/pkg/db/sqlstore"
	"chatServer/pkg/db/sqlstore/postgres"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

type App struct {
	server  Server
	service *service.Service
	logger  *logger.Logger
	db      *sqlstore.Store
}

func NewApp(c *Config) (*App, error) {
	app := &App{}
	var err error

	if app.logger, err = logger.NewLogger(c.Logger); err != nil {
		return nil, fmt.Errorf("invalid logger configuration: %v", err)
	}

	if app.db, err = postgres.NewStore(c.Store); err != nil {
		return nil, fmt.Errorf("cannot create new store: %v", err)
	}

	r := repository.NewRepository(app.db)
	app.service = service.NewService(r)

	h, err := handler.NewHandler(app.service, app.logger)
	if err != nil {
		return nil, fmt.Errorf("cannot create handler: %v", err)
	}

	app.server = server.NewServer(c.Server, h)

	return app, nil
}

func (a *App) Start(c *Config) error {
	a.logger.Info("starting server")

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		a.logger.Info("gracefully shutdown started")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

		go func() {
			defer cancel()
			err := a.server.Stop(ctx)
			if err != nil {
				a.logger.Error(err)
			}
		}()

		<-ctx.Done()

		err := a.db.Close()
		if err != nil {
			a.logger.Error(err)
		}
		a.logger.Info("server was shutdown")
	}()

	return a.server.Start()
}
