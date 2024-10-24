package models

import "time"

type Response struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *Data      `json:"data,omitempty"`  // Omit if Data is nil
	Error   *ErrorInfo `json:"error,omitempty"` // Omit if Error is nil
}

type Data struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
