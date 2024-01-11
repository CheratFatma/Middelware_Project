package services

import (
	"Middleware_Project/Tchipify/internal/models"
	repositories "Middleware_Project/Tchipify/internal/repositories/songs"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// SongsService gère la logique métier liée aux chansons
type SongsService struct {
	Repo *repositories.SongsRepository
}

//Récupèrer toutes les chansons
func (s *SongsService) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := s.Repo.GetAllSongs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//Récupèrer une chanson par son ID
func (s *SongsService) GetSongByID(w http.ResponseWriter, r *http.Request) {
	songIDParam := chi.URLParam(r, "id")
	songID, err := strconv.Atoi(songIDParam)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	song, err := s.Repo.GetSongByID(songID)
	if err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(song)
}

//Ajoute une nouvelle chanson
func (s *SongsService) AddNewSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Repo.AddNewSong(newSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New song added successfully"))
}

//Mettre à jour une chanson par son ID
func (s *SongsService) UpdateSongByID(w http.ResponseWriter, r *http.Request) {
	songIDParam := chi.URLParam(r, "id")
	songID, err := strconv.Atoi(songIDParam)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	var updatedSong models.Song
	err = json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Repo.UpdateSongByID(songID, updatedSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Song updated successfully"))
}

//Supprime une chanson par son ID
func (s *SongsService) DeleteSongByID(w http.ResponseWriter, r *http.Request) {
	songIDParam := chi.URLParam(r, "id")
	songID, err := strconv.Atoi(songIDParam)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	err = s.Repo.DeleteSongByID(songID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Song deleted successfully"))
}
