package auth

import (
	"GoAuth/internal/errs"
	"GoAuth/internal/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewAuthHandler(auth *Service) *Handler {
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
		switch {
		case errors.Is(err, errs.ErrInvalidPassword):
			utils.WriteError(w, 400, errs.ErrInvalidPassword)
		case errors.Is(err, errors.New("no rows in result set")):
			utils.WriteError(w, 400, errs.ErrUserNotFound)
		default:
			log.Printf("svc login error: %v", err)
			utils.WriteError(w, 500, errs.ErrInternalServerError)
		}
		return
	}
	utils.WriteJSON(w, 200, map[string]any{
		"accessToken":  "Hello From Access Token",
		"refreshToken": "Hello From Refresh Token",
	})
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
		Data:    user,
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Printf("failed to write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
