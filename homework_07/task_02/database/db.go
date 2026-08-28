package database

import (
	"database/sql"
	"fmt"

	"task_02/models"

	_ "modernc.org/sqlite"
)

// 'OpenDatabase' opens an existing database
func OpenDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Checking the connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("DB connection error: %w", err)
	}

	return db, nil
}

// 'GetAllUsers' returns all users (without password hash)
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	query := "SELECT id, email, name, is_active FROM Users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("request execution error: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("line scan error: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading lines: %w", err)
	}

	return users, nil
}

// 'GetUserByID' returns a user by ID (without password hash)
func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	query := "SELECT id, email, name, is_active FROM Users WHERE id = ?"

	var user models.User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.IsActive)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, fmt.Errorf("request execution error: %w", err)
	}

	return &user, nil
}

// 'DeleteUserByID' deletes a user by ID
func DeleteUserByID(db *sql.DB, id int) (bool, error) {
	query := "DELETE FROM Users WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return false, fmt.Errorf("request execution error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error getting number of deleted rows: %w", err)
	}

	return rowsAffected > 0, nil
}
