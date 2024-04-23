package main

import (
	"expense-manager-backend/core"
	"expense-manager-backend/middleware"
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

	appLogger := core.NewLogger()
	middlewareLogger := middleware.NewMiddlewareLogger(appLogger)
	app := core.NewServer(*cfg, appLogger)
	ConfigureRoutes(app)
	app.SetMiddleware(middlewareLogger.Logging)
	app.Start(cfg.Port)
}
