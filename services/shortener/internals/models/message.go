package models

// To implement basic rate limit.
type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}
