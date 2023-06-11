package model

// User represents a user account in the database.
type User struct {
	ID       int64  `json:"id"`
	GoogleID string `json:"google_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
