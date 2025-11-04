package models

// ErrorResponse represents a standardized error response for the API.
type ErrorResponse struct {
	Code    int    `json:"code" example:400`
	Message string `json:"message" example:"Bad Request"`
}
