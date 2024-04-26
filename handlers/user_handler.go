package handlers

import (
	"expense-manager-backend/core"
	"expense-manager-backend/services/user"
	"net/http"
)

type UserHandler struct {
	coreHandler *core.ApiHandler
	userService *user.UserService
}

func NewUserHandler(coreHandler *core.ApiHandler, userService *user.UserService) UserHandler {
	return UserHandler{
		coreHandler: coreHandler,
		userService: userService,
	}
}

func (h *UserHandler) RegisterRoutes() http.Handler {
	handler := http.NewServeMux()
	h.coreHandler.AddRoute("POST", "/", h.GetAllUser, handler)
	return handler
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}
