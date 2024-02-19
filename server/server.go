package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/humweb/inertia-go"
	"github.com/humweb/inertia-go-example/server/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Server struct {
	Port    int
	Router  *chi.Mux
	Inertia *inertia.Inertia
	Db      *gorm.DB
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.Port), s.Router)
}
func NewServer(port int, url string) (*Server, error) {
	app := &Server{
		Port:    port,
		Inertia: inertia.New(url, "./client/index.html", ""),
	}

	app.Db = models.InitDb()

	app.Router = BindRoutes(app)
	return app, nil
}
