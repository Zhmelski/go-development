package database

import (
	"database/sql"
	"fmt"

	"task_01/models"

	_ "modernc.org/sqlite"
)

const dbName = "users.db"

// 'InitDatabase' creates a database file and a Users table
func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		name TEXT NOT NULL,
		is_active BIT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("table creation error: %w", err)
	}

	fmt.Println("Database and 'Users' table have been successfully created.")
	return db, nil
}

// 'InsertUser' adds a user to the database
func InsertUser(db *sql.DB, user *models.User) error {
	insertSQL := `INSERT INTO Users (email, password_hash, name, is_active) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, user.Email, user.PasswordHash, user.Name, user.IsActive)
	if err != nil {
		return fmt.Errorf("user insert error: %w", err)
	}
	return nil
}

// 'GetAllUsers' returns all users from the database
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, email, password_hash, name, is_active FROM Users")
	if err != nil {
		return nil, fmt.Errorf("user reading error: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("line scan error: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
