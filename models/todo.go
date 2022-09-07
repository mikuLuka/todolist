package models

// Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	status bool   `json:"status"`
}
