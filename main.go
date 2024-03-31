package main

import (
	"expense-manager-backend/api/handlers/income"
	"expense-manager-backend/api/handlers/user"
	"flag"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	listenAddr := flag.String("port", ":4000", "port to listen on")
	flag.Parse()
	logger := zerolog.New(os.Stdout)
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	api := e.Group("api")
	userApi := api.Group("/users")
	userApi.POST("", user.CreateUser)
	userApi.GET("", user.GetAllUser)
	userApi.GET("/:id", user.GetUser)
	userApi.PUT("/:id", user.UpdateUser)
	userApi.DELETE("/:id", user.DeleteUser)

	incomeApi := api.Group("/income")
	incomeApi.POST("", income.CreateIncome)
	incomeApi.GET("/", income.GetAllIncome)
	incomeApi.GET("/:id", income.GetIncome)
	incomeApi.PUT("/:id", income.UpdateIncome)
	incomeApi.DELETE("/:id", income.DeleteIncome)

	for _, route := range e.Routes() {
		logger.Info().Msg("Route Registered: " + route.Path + " " + route.Method)
	}
	e.Logger.Fatal(e.Start(*listenAddr))
}
