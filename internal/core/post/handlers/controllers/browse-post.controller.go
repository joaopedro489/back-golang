package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	"github.com/joaopedro489/back-golang/internal/core/post/types/output"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
)

type BrowsePostController struct {
	browsePost usecases.BrowsePost
}

func NewBrowsePostController(browsePost usecases.BrowsePost) *BrowsePostController {
	return &BrowsePostController{
		browsePost: browsePost,
	}
}

func (h *BrowsePostController) BrowsePosts(c *gin.Context) {
	var params input.QueryPost
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	posts, err := h.browsePost.Execute(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := output.NewPostsResponse(posts)
	c.JSON(http.StatusOK, response)
}
