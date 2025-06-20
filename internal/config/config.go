package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	SecretJWT   string
	JWTExpires  string
}

var Env Config

func Load() {
	println("Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}
	Env.Port = os.Getenv("PORT")
	Env.DatabaseURL = os.Getenv("DATABASE_URL")
	Env.SecretJWT = os.Getenv("SECRET_JWT")
	Env.JWTExpires = os.Getenv("JWT_EXPIRES_IN")
	println("Environment variables loaded successfully.")
}
