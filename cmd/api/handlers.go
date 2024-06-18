package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/obiMadu/goJWTAuth/internals/db"
	"github.com/obiMadu/goJWTAuth/internals/models"
)

// TODO: create handlers
func login(c *gin.Context) {

	db := db.RawDB()

	requestPayload := jsonRequest{}

	err := c.ShouldBindBodyWithJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Status:  "error",
			Message: "Invalid credentials.",
		})
		return
	}

	// validate the user against the database
	user, err := models.GetByEmail(db, requestPayload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Status:  "error",
			Message: "Invalid credentials.",
		})
		return
	}

	valid, err := models.PasswordMatches(user, requestPayload.Password)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Status:  "error",
			Message: "Invalid credentials.",
		})
		return
	}

	payload := jsonResponse{
		Status:  "success",
		Message: "Logged In successfully.",
		Data: gin.H{
			"token": "jwt",
		},
	}

	c.JSON(http.StatusOK, payload)
}
