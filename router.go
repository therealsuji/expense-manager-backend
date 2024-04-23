package main

import (
	"expense-manager-backend/core"
	"expense-manager-backend/handlers/income"
	"expense-manager-backend/handlers/user"
	"net/http"
)

func ConfigureRoutes(server *core.Server) http.Handler {
	router := http.NewServeMux()
	apiv1 := http.NewServeMux()

	commonHandler := core.ApiHandler{Logger: server.AppLogger}
	userHandler := user.UserHandler{ApiHandler: &commonHandler}
	incomeHandler := income.IncomeHandler{ApiHandler: &commonHandler}

	apiv1.Handle("/users", userHandler.RegisterRoutes())
	apiv1.Handle("/income", incomeHandler.RegisterRoutes())

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiv1))

	return router
}
