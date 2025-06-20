package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
)

type AddPostController struct {
	addPost usecases.AddPost
}

func NewAddPostController(addPost usecases.AddPost) *AddPostController {
	return &AddPostController{
		addPost: addPost,
	}
}

func (h *AddPostController) AddPost(c *gin.Context) {
	authCtx, ok := auth.FromContext(c.Request.Context())
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}
	userId := authCtx.GetUserID()

	var params input.AddPost
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.addPost.Execute(params, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
