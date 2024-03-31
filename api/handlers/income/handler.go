package income

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateIncome(c echo.Context) error {
	return c.String(http.StatusOK, "Create income")
}

func GetAllIncome(c echo.Context) error {
	return c.String(http.StatusOK, "Get all income")
}

func GetIncome(c echo.Context) error {
	incomeId := c.Param("id")
	return c.String(http.StatusOK, "Get income "+incomeId)
}

func UpdateIncome(c echo.Context) error {
	return c.String(http.StatusOK, "Update income")
}

func DeleteIncome(c echo.Context) error {
	return c.String(http.StatusOK, "Delete income")
}
