package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/auth/handlers"
)

func RegisterAuthRoutes(r *gin.Engine, authHandlers *handlers.LoginController) {
	r.POST("/login", authHandlers.LoginUser)
}
