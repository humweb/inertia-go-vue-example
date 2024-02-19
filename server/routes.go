package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
)

func BindRoutes(s *Server) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {

		r.Use(
			s.Inertia.Middleware,
		)

		usr := NewUsersHandler(s)
		hr := NewHomeRoutes(s)

		r.Group(func(r chi.Router) {
			r.Get("/users", usr.HandleGetUsers)
			r.Get("/about", hr.HandleAbout)

			r.Get("/", hr.HandleIndex)
		})
	})

	// Static files

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "client/dist/assets"))
	fs := http.FileServer(filesDir)
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
	return r
}
