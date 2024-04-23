package core

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler
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

func createStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}
func (server *Server) SetMiddleware(xs ...Middleware) {
	stack := createStack(xs...)
	server.HttpServer.Handler = stack(server.HttpServer.Handler)
}
