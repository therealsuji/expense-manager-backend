package income

import (
	"expense-manager-backend/core"
	"net/http"
)

type IncomeHandler struct {
	*core.ApiHandler
}

func (h *IncomeHandler) RegisterRoutes() http.Handler {
	handler := http.NewServeMux()
	h.AddRoute("GET", "/income", h.CreateIncome, handler)
	return handler
}

func (h *IncomeHandler) CreateIncome(w http.ResponseWriter, r *http.Request) error {
	h.Logger.Info("Create income")

	w.Write([]byte("Create income"))

	return nil
}

func (h *IncomeHandler) GetAllIncome(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *IncomeHandler) GetIncome(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *IncomeHandler) UpdateIncome(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *IncomeHandler) DeleteIncome(w http.ResponseWriter, r *http.Request) error {
	return nil
}
