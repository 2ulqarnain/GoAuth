package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewAuthRouter(svc *AuthService) chi.Router {
	r := chi.NewRouter()

	authHandler := NewAuthHandler(svc)

	r.Use(middleware.Logger)
	r.Get("/", RootHandler)
	r.Post("/login", authHandler.LoginHandler)

	return r
}
