package main

import (
	"expense-manager-backend/middleware"
	"expense-manager-backend/utils"
	"flag"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)
	flag.Parse()

	env := utils.Environment{}
	env.Load()
	if err := env.Validate(); err != nil {
		logger.Fatal().Msg(err.Error())
	}

	run(&env)
}

func run(e *utils.Environment) {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	appLogger := utils.AppLogger{Logger: logger}
	middlewareLogger := middleware.MiddlewareLogger{
		AppLogger: utils.AppLogger{Logger: logger},
	}

	stack := middleware.CreateStack(
		middlewareLogger.Logging,
	)

	server := http.Server{
		Addr:    ":" + e.Port,
		Handler: stack(loadRoutes(appLogger)),
	}

	server.ListenAndServe()

}
