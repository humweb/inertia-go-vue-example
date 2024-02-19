package server

import (
	"github.com/humweb/inertia-go"
	"github.com/humweb/inertia-go-example/server/models"
	"github.com/humweb/inertia-go-example/server/resources"
	"net/http"
)

type HomeHandler struct {
	App *Server
}

func NewHomeRoutes(server *Server) *HomeHandler {
	return &HomeHandler{
		App: server,
	}
}
func (h HomeHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {

	h.App.Inertia.Render(w, r, "Index", inertia.Props{})
}

func (h HomeHandler) HandleAbout(w http.ResponseWriter, r *http.Request) {

	h.App.Inertia.Render(w, r, "About", inertia.Props{})
}

type UsersHandler struct {
	App *Server
}

func NewUsersHandler(server *Server) *UsersHandler {
	return &UsersHandler{
		App: server,
	}
}

// HandleGetUsers Table listing of users
func (h UsersHandler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {

	resource := resources.NewUserResource(h.App.Db, r)

	var model []models.User
	response, _ := resource.Paginate(resource, model)

	_ = h.App.Inertia.Render(w, r, "Users", response)
}
