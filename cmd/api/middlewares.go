package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/obiMadu/goJWTAuth/internals/jwtmod"
)

func AuthMiddleware() gin.HandlerFunc {
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

		claims := &jwtmod.JwtClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtmod.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, jsonResponse{
				Status:  "error",
				Message: "Authorization header missing.",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
