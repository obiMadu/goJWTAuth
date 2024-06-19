package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/obiMadu/goJWTAuth/internals/db"
	"github.com/obiMadu/goJWTAuth/internals/jwtmod"
	"github.com/obiMadu/goJWTAuth/internals/models"
)

// TODO: create handlers
func login(c *gin.Context) {

	requestPayload := jsonRequest{}

	err := c.ShouldBindBodyWithJSON(&requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsonResponse{
			Status:  "error",
			Message: "Invalid request.",
		})
		return
	}

	// validate the user against the database
	user, err := models.GetByEmail(db.RawDB(), requestPayload.Email)
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

	token, err := jwtmod.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{
			Status:  "error",
			Message: "Failed to generate token.",
		})
	}

	payload := jsonResponse{
		Status:  "success",
		Message: "Logged In successfully.",
		Data: gin.H{
			"token": token,
		},
	}

	c.JSON(http.StatusOK, payload)
}

func getProfile(c *gin.Context) {
	var claimsJSON jwtmod.JwtClaim

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, jsonResponse{
			Status:  "error",
			Message: "Unable to retrieve claims.",
		})
	}

	// Convert the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{
			Status:  "error",
			Message: "Could not marshal claims.",
		})
	}

	if err := json.Unmarshal(jsonBytes, &claimsJSON); err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{
			Status:  "error",
			Message: "Could not unmarshal claims.",
		})
	}

	user, err := models.GetByEmail(db.RawDB(), claimsJSON.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{
			Status:  "error",
			Message: "Couldn't retrieve user.",
		})
		return
	}

	payload := jsonResponse{
		Status:  "success",
		Message: "User authenticated successfully.",
		Data: gin.H{
			"user": user,
		},
	}

	c.JSON(http.StatusOK, payload)
}
