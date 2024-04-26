package core

import (
	db_connection "expense-manager-backend/db"
	db "expense-manager-backend/db/sqlc"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New(validator.WithRequiredStructEnabled())

type Middleware func(http.Handler) http.Handler
type Server struct {
	Environment Environment
	HttpServer  http.Server
	Logger      Logger
	Validator   *validator.Validate
	Queries     *db.Queries
}

func NewServer(env Environment) *Server {
	db_connection.Init(
		env.DB.User,
		env.DB.Password,
		env.DB.Host,
		env.DB.Port,
		env.DB.Name,
	)
	return &Server{
		Environment: env,
		HttpServer:  http.Server{},
		Logger:      GetLogger(),
		Validator:   Validator,
		Queries:     db.New(db_connection.GetDB()),
	}
}

func (server *Server) Start(addr string) error {
	server.Logger.Info("Server started on port " + addr)
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
