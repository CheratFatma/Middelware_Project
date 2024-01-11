package main

import (
	"Middleware_Project/Tchipify/internal/controllers/songs"
	"Middleware_Project/Tchipify/internal/helpers"
	repositories "Middleware_Project/Tchipify/internal/repositories/songs"
	services "Middleware_Project/Tchipify/internal/services/songs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialisation de la base de données
	db, err := helpers.InitDatabase()
	if err != nil {
		log.Fatal("Erreur lors de l'initialisation de la base de données :", err)
	}
	defer db.Close()

	// Initialisation du repository songsRepository
	songsRepository := &repositories.SongsRepository{DB: db}

	// Initialisation du service songsService
	songsService := &services.SongsService{Repo: songsRepository}

	// Création du contexte songsContext
	songsContext := songs.NewContext(db, songsService)

	// Configuration du routeur
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configuration des routes pour le contexte songsContext
	songsContext.SetRoutes(r)

	// Démarrage du serveur web
	log.Println("[INFO] Serveur web démarré. En écoute sur *:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
