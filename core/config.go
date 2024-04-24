package core

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type Environment struct {
	Port string
	DB   DBConfig
}

func (e *Environment) Validate() error {
	if e.Port == "" {
		return AppError{"Port is required"}
	}

	if e.DB.User == "" {
		return AppError{"DB User is required"}
	}

	if e.DB.Password == "" {
		return AppError{"DB Password is required"}
	}

	if e.DB.Name == "" {
		return AppError{"DB Name is required"}
	}

	if e.DB.Host == "" {
		return AppError{"DB Host is required"}
	}

	if e.DB.Port == "" {
		return AppError{"DB Port is required"}
	}

	return nil
}

func (e *Environment) Load() error {
	err := godotenv.Load()
	if err != nil {
		return AppError{"Error loading .env file"}
	}

	e.DB.User = os.Getenv("DB_USERNAME")
	e.Port = os.Getenv("PORT")
	e.DB.Password = os.Getenv("DB_PASSWORD")
	e.DB.Name = os.Getenv("DB_DATABASE")
	e.DB.Host = os.Getenv("DB_HOST")
	e.DB.Port = os.Getenv("DB_PORT")

	return nil
}
