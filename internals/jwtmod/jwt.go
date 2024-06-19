package jwtmod

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/obiMadu/goJWTAuth/internals/models"
)

var JwtKey []byte

type JwtClaim struct {
	UserID int `json:"userId" binding:"required"`
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	// create expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JwtClaim{
		UserID: user.ID,
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
