package controller

import (
	"encoding/json"
	"layered-arch/internal/domain"
	"net/http"
	"strconv"
)

type UserHandler struct {
	svc domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	result, err := h.svc.GetDisplayName(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"display_name": result})
}
