package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
)

type ReadPostController struct {
	readPost usecases.ReadPost
}

func NewReadPostController(readPost usecases.ReadPost) *ReadPostController {
	return &ReadPostController{
		readPost: readPost,
	}
}

func (h *ReadPostController) ReadPost(c *gin.Context) {
	var params input.ParamPost
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	post, err := h.readPost.Execute(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
