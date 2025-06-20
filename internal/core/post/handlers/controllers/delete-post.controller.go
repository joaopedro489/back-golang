package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
)

type DeletePostController struct {
	deletePost usecases.DeletePost
}

func NewDeletePostController(deletePost usecases.DeletePost) *DeletePostController {
	return &DeletePostController{
		deletePost: deletePost,
	}
}

func (h *DeletePostController) DeletePost(c *gin.Context) {
	authCtx, ok := auth.FromContext(c.Request.Context())
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}
	userId := authCtx.GetUserID()

	var params input.ParamPost
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	if err := h.deletePost.Execute(params, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
