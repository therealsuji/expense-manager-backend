package core

import (
	"os"

	"github.com/rs/zerolog"
)

type AppLogger struct {
	*zerolog.Logger
}

func (l *AppLogger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l *AppLogger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

func NewLogger() AppLogger {
	applogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return AppLogger{
		Logger: &applogger,
	}
}
