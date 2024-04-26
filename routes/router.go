package routes

import (
	"expense-manager-backend/core"
	"expense-manager-backend/handlers"
	"expense-manager-backend/services/user"
	"net/http"
)

func ConfigureRoutes(server *core.Server) {
	router := http.NewServeMux()
	apiv1 := http.NewServeMux()

	commonHandler := core.NewApiHandler(server.Logger, server.Validator, server.Queries)

	userService := user.NewUserService(server.Queries)

	userHandler := handlers.NewUserHandler(&commonHandler, userService)
	authHandler := handlers.NewAuthHandler(&commonHandler, userService)

	apiv1.Handle("/auth/", http.StripPrefix("/auth", authHandler.RegisterRoutes()))
	apiv1.Handle("/users/", http.StripPrefix("/users", userHandler.RegisterRoutes()))

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiv1))

	server.HttpServer.Handler = router
}
