package main

type jsonRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type jsonResponse struct {
	Status  string `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
	Data    any    `json:"data" binding:"required"`
	Error   any    `json:"error,omitempty"`
}
