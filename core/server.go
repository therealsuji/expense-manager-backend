package core

import (
	database "expense-manager-backend/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New(validator.WithRequiredStructEnabled())

type Middleware func(http.Handler) http.Handler
type Server struct {
	Environment Environment
	HttpServer  http.Server
	AppLogger   AppLogger
	Validator   *validator.Validate
	DB          *database.DB
}

func NewServer(env Environment, appLogger AppLogger) *Server {
	return &Server{
		Environment: env,
		HttpServer:  http.Server{},
		AppLogger:   appLogger,
		Validator:   Validator,
		DB:          database.GetDB(),
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
