package router

import (
	"net/http"
	"pathshala/internal/handlers"
	"pathshala/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	adminRepo := repository.NewAdminRepository(db)
	authHandler := handlers.NewAuthHandler(adminRepo)

	r.Group(func(r chi.Router) {
		r.Post("/api/auth/login", authHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Get("/api/dashboard", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to Admin Dashboard!"))
		})
	})

	return r
}
