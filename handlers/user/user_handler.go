package user

import (
	"expense-manager-backend/core"
	"expense-manager-backend/requests"
	"expense-manager-backend/utils"
	"net/http"
)

type UserHandler struct {
	*core.ApiHandler
}

func (h *UserHandler) RegisterRoutes() http.Handler {
	handler := http.NewServeMux()
	h.AddRoute("POST", "/users", h.CreateUser, handler)
	return handler
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var createUserRequest requests.CreateUserRequest
	if err := utils.ParseJSON(r, &createUserRequest); err != nil {
		return utils.BadRequest
	}

	if err := h.Validator.Struct(createUserRequest); err != nil {
		return utils.BadRequest.UnprocessableEntity(err.Error())
	}

	return nil
}

func GetAllUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil

}
