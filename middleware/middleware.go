package middleware

import (
	"expense-manager-backend/core"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}
func SetMiddleware(server *core.Server, xs ...Middleware) {
	stack := CreateStack(xs...)
	server.HttpServer.Handler = stack(server.HttpServer.Handler)
}
