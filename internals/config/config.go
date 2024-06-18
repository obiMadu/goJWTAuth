package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/obiMadu/goJWTAuth/internals/db"
	"github.com/obiMadu/goJWTAuth/internals/jwtmod"
)

func Config() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load .env file %s\n", err.Error())
	}

	// init db
	db.InitDB()

	// init jwt
	jwtmod.JwtKey = []byte(os.Getenv("JWTKEY"))
}
