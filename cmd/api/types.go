package main

type jsonRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type jsonResponse struct {
	Status  string         `json:"status" binding:"required"`
	Message string         `json:"message" binding:"required"`
	Data    map[string]any `json:"data" binding:"required"`
	Error   map[string]any `json:"error,omitempty"`
}
