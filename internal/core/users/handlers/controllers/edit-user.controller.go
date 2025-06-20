package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
	"github.com/joaopedro489/back-golang/internal/core/users/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type EditUserController struct {
	editUser usecases.EditUser
}

func NewEditUserController(editUser usecases.EditUser) *EditUserController {
	return &EditUserController{
		editUser: editUser,
	}
}

func (h *EditUserController) EditUser(c *gin.Context) {
	var params input.ParamUser
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	var body input.EditUser
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := h.editUser.Execute(body, params)
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
