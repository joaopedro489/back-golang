package auth

import (
	"github.com/joaopedro489/back-golang/internal/core/auth/handlers"
	"github.com/joaopedro489/back-golang/internal/core/auth/repository"
	usecases "github.com/joaopedro489/back-golang/internal/core/auth/use-cases"
	"gorm.io/gorm"
)

func Config(db *gorm.DB) *handlers.LoginController {
	authRepo := repository.NewUserRepository(db)

	loginUC := usecases.NewLogin(authRepo)
	login := handlers.NewLoginController(*loginUC)

	return login
}
