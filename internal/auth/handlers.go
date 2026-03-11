package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	svc *AuthService
}

func NewAuthHandler(auth *AuthService) *Handler {
	return &Handler{svc: auth}
}

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("auth ok"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload loginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.svc.Login(r.Context(), payload)
	if err != nil {
		if err.Error() == "incorrect password" {
			http.Error(w, "incorrect password", http.StatusUnauthorized)
		} else if err.Error() == "no rows in result set" {
			http.Error(w, "No account against provided credentials", http.StatusUnauthorized)
		} else {
			fmt.Printf("failed to login: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	w.Write([]byte("ok"))
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var payload signupPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload format not correct", http.StatusBadRequest)
	}
	user, err := h.svc.Signup(r.Context(), payload)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &signupResponse{
		Ok:      true,
		Message: "signup ok",
		Data: &signupResponseData{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Printf("failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
