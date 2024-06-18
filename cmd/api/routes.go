package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	mux := gin.Default()

	mux.Use(cors.Default())

	// TODO: create routes
	mux.POST("/login", login)

	return mux
}
