package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
	"github.com/joaopedro489/back-golang/internal/core/users/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type ReadUserController struct {
	readUser usecases.ReadUser
}

func NewReadUserController(readUser usecases.ReadUser) *ReadUserController {
	return &ReadUserController{
		readUser: readUser,
	}
}

func (h *ReadUserController) ReadUser(c *gin.Context) {
	var params input.ParamUser
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametros inválidos"})
		return
	}

	user, err := h.readUser.Execute(params)
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
