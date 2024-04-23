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
	middlewareLogger := middleware.MiddlewareLogger{
		AppLogger: appLogger,
	}

	stack := middleware.CreateStack(
		middlewareLogger.Logging,
	)
	app := core.NewServer(*cfg, appLogger)
	ConfigureRoutes(app)
	middleware.SetMiddleware(app, stack)
	app.Start(cfg.Port)
}
