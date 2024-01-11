package users

import (
	services "Middleware_Project/Tchipify/internal/services/users"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// Context = Services nécessaires pour gérer les requêtes liées aux users
type Context struct {
	DB           *sql.DB
	UsersService *services.UsersService
}

// NewContext crée une nouvelle instance de Context
func NewContext(db *sql.DB, usersService *services.UsersService) *Context {
	return &Context{
		DB:           db,
		UsersService: usersService,
	}
}

// Les routes pour les gestionnaires d'utilisateurs
func (c *Context) SetRoutes(r *chi.Mux) {
	r.Get("/users", c.UsersService.GetAllUsers)
	r.Get("/users/{id}", c.UsersService.GetUserByID)
	r.Post("/users", c.UsersService.AddNewUser)
	r.Put("/users/{id}", c.UsersService.UpdateUserByID)
	r.Delete("/users/{id}", c.UsersService.DeleteUserByID)
}
