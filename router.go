package main

import (
	"expense-manager-backend/handlers/user"
	"expense-manager-backend/utils"
	"net/http"
)

func loadRoutes(logger utils.AppLogger) http.Handler {
	router := http.NewServeMux()
	apiv1 := http.NewServeMux()

	commonHandler := utils.ApiHandler{Logger: logger}

	userHandler := user.UserHandler{ApiHandler: commonHandler}

	apiv1.Handle("/users", userHandler.RegisterRoutes())
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiv1))

	return router
}
