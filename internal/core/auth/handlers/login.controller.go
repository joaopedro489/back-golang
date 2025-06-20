package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/auth/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/auth/use-cases"
)

type LoginController struct {
	Login usecases.Login
}

func NewLoginController(login usecases.Login) *LoginController {
	return &LoginController{
		Login: login,
	}
}

func (h *LoginController) LoginUser(c *gin.Context) {
	var body input.LoginInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.Login.Execute(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
