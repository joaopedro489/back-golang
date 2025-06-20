package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	types "github.com/joaopedro489/back-golang/internal/core/users/types/input"
	"github.com/joaopedro489/back-golang/internal/core/users/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type AddUserController struct {
	addUser usecases.AddUser
}

func NewUserController(addUser usecases.AddUser) *AddUserController {
	return &AddUserController{addUser: addUser}
}

func (h *AddUserController) AddUser(c *gin.Context) {
	var params types.AddUser
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.addUser.Execute(params)
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

	c.JSON(http.StatusCreated, response)
}
