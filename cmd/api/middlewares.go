package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/obiMadu/goJWTAuth/internals/jwtmod"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, jsonResponse{
				Status:  "error",
				Message: "Authorization header missing.",
			})
			c.Abort()
			return
		}

		claims := jwtmod.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtmod.JwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, jsonResponse{
				Status:  "error",
				Message: "Error validating token.",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, jsonResponse{
				Status:  "error",
				Message: "Invalid token.",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
