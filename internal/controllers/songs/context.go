package songs

import (
	services "Middleware_Project/Tchipify/internal/services/songs"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// Context = Services nécessaires pour gérer les requêtes liées aux chansons
type Context struct {
	DB           *sql.DB
	SongsService *services.SongsService
}

// NewContext crée une nouvelle instance de Context
func NewContext(db *sql.DB, songsService *services.SongsService) *Context {
	return &Context{
		DB:           db,
		SongsService: songsService,
	}
}

// Les routes pour les gestionnaires de chansons
func (c *Context) SetRoutes(r *chi.Mux) {
	r.Get("/songs", c.SongsService.GetAllSongs)
	r.Get("/songs/{id}", c.SongsService.GetSongByID)
	r.Post("/songs", c.SongsService.AddNewSong)
	r.Put("/songs/{id}", c.SongsService.UpdateSongByID)
	r.Delete("/songs/{id}", c.SongsService.DeleteSongByID)
}
