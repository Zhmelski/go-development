package models

// 'User' represents the user in the database
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Not included in JSON
	Name         string `json:"name"`
	IsActive     bool   `json:"is_active"`
}

// 'UserResponse' represents a user without a password for the API
type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

// 'ToResponse' converts 'User' to 'UserResponse' (without password)
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		Name:     u.Name,
		IsActive: u.IsActive,
	}
}
