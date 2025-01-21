package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signup", h.handleSignup).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {

}