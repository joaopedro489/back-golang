package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
	"github.com/joaopedro489/back-golang/internal/core/users/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
)

type BrowseUserController struct {
	browseUser usecases.BrowseUser
}

func NewBrowseUserController(browseUser usecases.BrowseUser) *BrowseUserController {
	return &BrowseUserController{
		browseUser: browseUser,
	}
}

func (h *BrowseUserController) BrowseUsers(c *gin.Context) {
	var params input.QueryUser
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query invalida"})
		return
	}

	users, err := h.browseUser.Execute(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := output.NewUsersResponse(users)
	c.JSON(http.StatusOK, response)
}
