package repositories

import (
	"Middleware_Project/Tchipify/internal/models"
	"database/sql"
	"fmt"
)

// SongsRepository gère l'accès à la BDD pour les chansons
type SongsRepository struct {
	DB *sql.DB
}

//Récupèrer toutes les chansons depuis la BDD
func (r *SongsRepository) GetAllSongs() ([]models.Song, error) {
	rows, err := r.DB.Query("SELECT id, title, author FROM songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song

	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Author); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

//Récupèrer une chanson par son ID depuis la BDD
func (r *SongsRepository) GetSongByID(songID int) (*models.Song, error) {
	var song models.Song
	err := r.DB.QueryRow("SELECT id, title, author FROM songs WHERE id = ?", songID).Scan(&song.ID, &song.Title, &song.Author)
	if err != nil {
		fmt.Println("Error fetching song:", err)
		return nil, err
	}

	fmt.Println("Song found:", song)
	return &song, nil
}

//Insèrer une nouvelle chanson dans la BDD
func (r *SongsRepository) AddNewSong(song models.Song) error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM songs WHERE title = ?", song.Title).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = r.DB.Exec(`INSERT INTO songs (title, author) VALUES (?, ?)`, song.Title, song.Author)
		if err != nil {
			return err
		}
		fmt.Println("Chanson insérée avec succès")
	}
	return nil
}

//Mettre à jour une chanson par son ID dans la BDD
func (r *SongsRepository) UpdateSongByID(songID int, updatedSong models.Song) error {
	_, err := r.DB.Exec("UPDATE songs SET title = ?, author = ? WHERE id = ?", updatedSong.Title, updatedSong.Author, songID)
	if err != nil {
		return err
	}
	return nil
}

//Supprime une chanson par son ID de la BDD
func (r *SongsRepository) DeleteSongByID(songID int) error {
	_, err := r.DB.Exec("DELETE FROM songs WHERE id = ?", songID)
	if err != nil {
		return err
	}
	return nil
}
