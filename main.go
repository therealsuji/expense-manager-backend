package main

import (
	"expense-manager-backend/core"
	"expense-manager-backend/middleware"
	"expense-manager-backend/routes"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)

	cfg := core.Environment{}
	cfg.Load()
	if err := cfg.Validate(); err != nil {
		logger.Fatal().Msg(err.Error())
	}

	run(&cfg)
}

func run(cfg *core.Environment) {

	core.InitLogger()
	appLogger := core.GetLogger()
	middlewareLogger := middleware.NewMiddlewareLogger(appLogger)
	app := core.NewServer(*cfg)
	routes.ConfigureRoutes(app)
	app.SetMiddleware(middlewareLogger.Logging)
	app.Start(cfg.Port)
}
