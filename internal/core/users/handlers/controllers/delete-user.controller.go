package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type DeleteUserController struct {
	deleteUser usecases.DeleteUser
}

func NewDeleteUserController(deleteUser usecases.DeleteUser) *DeleteUserController {
	return &DeleteUserController{
		deleteUser: deleteUser,
	}
}

func (h *DeleteUserController) DeleteUser(c *gin.Context) {
	var params input.ParamUser
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametros inválidos"})
		return
	}

	if err := h.deleteUser.Execute(params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
