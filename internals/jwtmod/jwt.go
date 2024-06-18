package jwtmod

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/obiMadu/goJWTAuth/internals/models"
)

var JwtKey []byte

type JwtClaim struct {
	UserID int    `json:"userId" binding:"required"`
	Email  string `json:"email" binding:"required"`
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	// create expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JwtClaim{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
