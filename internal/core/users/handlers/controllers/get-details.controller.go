package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/users/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type GetDetailsController struct {
	getDetails usecases.GetDetails
}

func NewGetDetailsController(readUser usecases.GetDetails) *GetDetailsController {
	return &GetDetailsController{
		getDetails: readUser,
	}
}

func (h *GetDetailsController) GetDetails(c *gin.Context) {
	authCtx, ok := auth.FromContext(c.Request.Context())
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}
	userId := authCtx.GetUserID()

	user, err := h.getDetails.Execute(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := output.NewUserResponse(
		user.Name,
		user.Email,
		user.Birthday,
		user.Id,
	)

	c.JSON(http.StatusOK, response)
}
