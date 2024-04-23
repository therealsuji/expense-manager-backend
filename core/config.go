package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Port string
}

func (e *Environment) Validate() error {
	if e.Port == "" {
		return AppError{"Port is required"}
	}

	return nil
}

func (e *Environment) Load() error {
	err := godotenv.Load()
	if err != nil {
		return AppError{"Error loading .env file"}
	}

	e.Port = os.Getenv("PORT")

	return nil
}
