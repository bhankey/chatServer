package server

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	server http.Server
}

func NewServer(c *Config, h http.Handler) *Server {
	return &Server{
		server: http.Server{
			Addr: ":" + strconv.Itoa(c.Port),

			WriteTimeout: time.Second * time.Duration(c.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(c.ReadTimeout),
			IdleTimeout:  time.Second * time.Duration(c.IdleTimeout),

			Handler: h,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
