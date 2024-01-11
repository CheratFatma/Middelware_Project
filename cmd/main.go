package main

import (
	"Middleware_Project/Tchipify/internal/controllers/songs"
	"Middleware_Project/Tchipify/internal/controllers/users"
	"Middleware_Project/Tchipify/internal/helpers"

	repositoriesU "Middleware_Project/Tchipify/internal/repositories/users"
	servicesU "Middleware_Project/Tchipify/internal/services/users"

	repositoriesS "Middleware_Project/Tchipify/internal/repositories/songs"
	servicesS "Middleware_Project/Tchipify/internal/services/songs"

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

	//USERS
	usersRepository := &repositoriesU.UsersRepository{DB: db}
	usersService := &servicesU.UsersService{Repo: usersRepository}
	usersContext := users.NewContext(db, usersService)

	//SONGS
	songsRepository := &repositoriesS.SongsRepository{DB: db}
	songsService := &servicesS.SongsService{Repo: songsRepository}
	songsContext := songs.NewContext(db, songsService)

	// Configuration du routeur
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configuration des routes
	usersContext.SetRoutes(r)
	songsContext.SetRoutes(r)

	// Démarrage du serveur web
	log.Println("[INFO] Serveur web démarré. En écoute sur *:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
