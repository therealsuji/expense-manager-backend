package core

import (
	"net/http"
)

type Server struct {
	Environment Environment
	HttpServer  http.Server
	AppLogger   AppLogger
}

func NewServer(env Environment, appLogger AppLogger) *Server {
	return &Server{
		Environment: env,
		HttpServer:  http.Server{},
		AppLogger:   appLogger,
	}
}

func (server *Server) Start(addr string) error {
	server.AppLogger.Info("Server started on port " + addr)
	server.HttpServer.Addr = ":" + addr
	return server.HttpServer.ListenAndServe()
}
