package models

type Link struct {
	ID	int 			`json:"id" db:"id"`
	URL string 			`json:"url" db:"url"`
	Tag string			`json:"tag" db:"short_tag"`
	CreatedAt string	`json:"created_at" db:"created_at"`
}

type LinkRequest struct {
	URL string `json:"url" binding:"omitempty"`
}