package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
)

type LikePostController struct {
	likePost usecases.LikePost
}

func NewLikePostController(likePost usecases.LikePost) *LikePostController {
	return &LikePostController{
		likePost: likePost,
	}
}

func (h *LikePostController) LikePost(c *gin.Context) {
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

	if err := h.likePost.Execute(params.Id, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
