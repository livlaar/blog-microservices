package handler

import (
	"encoding/json"
	"net/http"

	model "github.com/livlaar/blog-microservices/shared/models"
	"github.com/livlaar/blog-microservices/users/internal/controller"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	ctrl *controller.UserController
}

func NewUserHandler(ctrl *controller.UserController) *UserHandler {
	return &UserHandler{ctrl: ctrl}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.ctrl.GetUserByID(id)
	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.ctrl.CreateUser(user); err != nil {
		http.Error(w, "Error guardando usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
