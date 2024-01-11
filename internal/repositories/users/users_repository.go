package repositories

import (
	"Middleware_Project/Tchipify/internal/models"
	"database/sql"
	"fmt"
)

// UsersRepository gère l'accès à la BDD pour les utilisateurs
type UsersRepository struct {
	DB *sql.DB
}

//Récupèrer tous les utilisateurs depuis la BDD
func (r *UsersRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT ID, Username, Email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//Récupèrer un utilisateur par son ID depuis la BDD
func (r *UsersRepository) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT ID, Username, Email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, err
	}

	fmt.Println("User found:", user)
	return &user, nil
}

//Insèrer un nouveau utilisateur dans la BDD
func (r *UsersRepository) AddNewUser(user models.User) error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE Username = ?", user.Username).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = r.DB.Exec(`INSERT INTO users (username, email) VALUES (?, ?)`, user.Username, user.Email)
		if err != nil {
			return err
		}
		fmt.Println("Utilisateur inséré avec succès")
	}
	return nil
}

//Mettre à jour un uilisateur par son ID dans la BDD
func (r *UsersRepository) UpdateUserByID(userID int, updatedUser models.User) error {
	_, err := r.DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", updatedUser.Username, updatedUser.Email, userID)
	if err != nil {
		return err
	}
	return nil
}

//Supprimer un user par son ID de la BDD
func (r *UsersRepository) DeleteUserByID(userID int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
