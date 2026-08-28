package responses

// 'ErrorResponse' represents the error structure
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// 'SuccessResponse' represents a successful response
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// 'UsersResponse' represents a response with a list of users
type UsersResponse struct {
	Users []interface{} `json:"users"`
	Count int           `json:"count"`
}
