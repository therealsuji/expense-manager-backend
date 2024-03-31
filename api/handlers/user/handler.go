package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Create user")
}

func GetAllUser(c echo.Context) error {
	return c.String(http.StatusOK, "Get all user")
}

func GetUser(c echo.Context) error {
	userId := c.Param("id")
	return c.String(http.StatusOK, "Get user"+userId)
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Update user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Delete user")
}
