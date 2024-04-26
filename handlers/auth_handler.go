package handlers

import (
	"expense-manager-backend/core"
	"expense-manager-backend/requests"
	"expense-manager-backend/services/user"
	"expense-manager-backend/utils"
	"net/http"
)

type AuthHandler struct {
	coreHandler *core.ApiHandler
	userService *user.UserService
}

func NewAuthHandler(coreHandler *core.ApiHandler, userService *user.UserService) AuthHandler {
	return AuthHandler{
		coreHandler: coreHandler,
		userService: userService,
	}
}

func (h *AuthHandler) RegisterRoutes() http.Handler {
	handler := http.NewServeMux()

	h.coreHandler.AddRoute("POST", "/register", h.RegisterUser, handler)
	h.coreHandler.AddRoute("POST", "/login", h.LoginUser, handler)
	return handler
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) error {
	var createUserRequest requests.CreateUserRequest
	if err := utils.ParseJSON(r, &createUserRequest); err != nil {
		return utils.BadRequest
	}

	if err := h.coreHandler.Validator.Struct(createUserRequest); err != nil {
		return utils.BadRequest.UnprocessableEntity(err.Error())
	}

	_, e := h.userService.CreateUser(r.Context(), user.CreateUserParams{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	})

	if e != nil {
		h.coreHandler.Logger.Error("Error creating user: " + e.Error())
		return utils.InternalServerError
	}

	return nil
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
