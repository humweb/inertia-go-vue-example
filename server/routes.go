package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/humweb/inertia-go"
	"github.com/humweb/inertia-go-example/server/models"
	"github.com/humweb/inertia-go-example/server/resources"
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

		r.Group(func(r chi.Router) {
			r.Get("/users", func(w http.ResponseWriter, r *http.Request) {

				var model []models.User
				resource := resources.NewUserResource(s.Db, r)
				response, _ := resource.Paginate(resource, model)

				_ = s.Inertia.Render(w, r, "Users", response)
			})

			r.Get("/about", func(w http.ResponseWriter, r *http.Request) {

				s.Inertia.Render(w, r, "About", inertia.Props{})
			})

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				s.Inertia.Render(w, r, "Index", inertia.Props{})
			})
		})
	})

	// Static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "client/dist/assets"))
	fs := http.FileServer(filesDir)
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
	return r
}
