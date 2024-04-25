package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, err error, status int) {
	WriteJSON(w, map[string]string{
		"code":    fmt.Sprintf("%d", status),
		"message": err.Error(),
	}, status)
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return ApiError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
	}
	return nil
}
