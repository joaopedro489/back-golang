package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaopedro489/back-golang/internal/config"
)

var jwtSecret = []byte(config.Env.SecretJWT)

func GenerateJWT(userID int) (string, error) {
	exp, err := strconv.Atoi(config.Env.JWTExpires)
	if err != nil {
		panic("Invalid JWT expiration time in config: " + config.Env.JWTExpires)
	}
	userIdString := strconv.Itoa(userID)
	claims := jwt.MapClaims{
		"user_id": userIdString,
		"exp":     time.Now().Add(time.Duration(exp) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return jwtSecret, nil
	})
}
