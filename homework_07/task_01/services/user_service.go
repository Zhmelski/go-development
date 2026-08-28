package services

import (
	"fmt"
	"strings"

	"task_01/models"
	"task_01/utils"

	"golang.org/x/crypto/bcrypt"
)

// 'CreateUserFromPerson' creates a user from person data
func CreateUserFromPerson(person utils.PersonData) (*models.User, error) {
	// Forming an email
	email := fmt.Sprintf("%s%s@coolcompany.com", person.FirstName, person.LastName)

	// Generating a password hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password hash generation error: %w", err)
	}

	// Forming a full name
	fullName := fmt.Sprintf("%s %s", person.FirstName, person.LastName)

	return &models.User{
		Email:        strings.ToLower(email),
		PasswordHash: string(passwordHash),
		Name:         fullName,
		IsActive:     true,
	}, nil
}
