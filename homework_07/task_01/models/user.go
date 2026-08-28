package models

// 'User' represents the system user
type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
	IsActive     bool
}
