package services

import (
	"basicMicroservice/config"
	"basicMicroservice/models"
	"errors"
	"log"
	"strconv"
)

// Get all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, name, email FROM users`

	err := config.DB.Select(&users, query)
	if err != nil {
		log.Printf("❌ Failed to get users: %v", err)
		return nil, err
	}

	return users, nil
}

// Get a user by ID (string input, comparing with int ID)
func GetUserById(id string) (models.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, errors.New("invalid ID format")
	}

	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`

	err = config.DB.Get(&user, query, intID)
	if err != nil {
		log.Printf("❌ Error fetching user with id %d: %v", intID, err)
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

// Create a new user
func CreateUser(user models.User) (models.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, user.Name, user.Email).Scan(&user.Id)
	if err != nil {
		log.Printf("❌ Error inserting user: %v\n", err)
		return models.User{}, err
	}
	return user, nil
}
