package models

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PlaylistTracksResponse struct {
	Name   string   `json:"name"`
	Tracks []string `json:"tracks"`
}
