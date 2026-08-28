package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"task_02/database"
	"task_02/responses"
)

// 'UserHandler' processes all requests to /users
type UserHandler struct {
	DB *sql.DB
}

// 'NewUserHandler' creates a new handler
func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// 'ServeHTTP' processes HTTP requests
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the path: /users or /users/{id}
	path := strings.TrimPrefix(r.URL.Path, "/users")
	path = strings.Trim(path, "/")

	switch r.Method {
	case http.MethodGet:
		if path == "" {
			// GET /users - get all users
			h.GetAllUsers(w, r)
		} else {
			// GET /users/{id} - get user by ID
			h.GetUserByID(w, r, path)
		}
	case http.MethodDelete:
		if path != "" {
			// DELETE /users/{id} - delete user
			h.DeleteUser(w, r, path)
		} else {
			h.sendError(w, http.StatusBadRequest, "User ID not specified")
		}
	default:
		h.sendError(w, http.StatusMethodNotAllowed, "Method not supported")
	}
}

// 'GetAllUsers' processes GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers(h.DB)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		h.sendError(w, http.StatusInternalServerError, "Error receiving data")
		return
	}

	// Convert to UserResponse (without passwords)
	userResponses := make([]interface{}, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	response := responses.UsersResponse{
		Users: userResponses,
		Count: len(users),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 'GetUserByID' processes GET /users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := database.GetUserByID(h.DB, id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		h.sendError(w, http.StatusInternalServerError, "Error receiving data")
		return
	}

	if user == nil {
		h.sendError(w, http.StatusNotFound, fmt.Sprintf("User with ID %d not found", id))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.ToResponse())
}

// 'DeleteUser' processes DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	deleted, err := database.DeleteUserByID(h.DB, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		h.sendError(w, http.StatusInternalServerError, "Error deleting user")
		return
	}

	if !deleted {
		h.sendError(w, http.StatusNotFound, fmt.Sprintf("User with ID %d not found", id))
		return
	}

	response := responses.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("User with ID %d has been successfully deleted", id),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 'sendError' sends an error in JSON format
func (h *UserHandler) sendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errorResponse := responses.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
