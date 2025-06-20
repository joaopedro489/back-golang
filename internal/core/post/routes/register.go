package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/post/handlers"
)

func RegisterPostRoutes(r *gin.Engine, postHandlers *handlers.PostHandlers) {
	postRoutes := r.Group("/posts")
	postRoutes.Use(auth.JWTMiddleware())
	postRoutes.GET("/", postHandlers.Browse.BrowsePosts)
	postRoutes.GET("/:id", postHandlers.Read.ReadPost)
	postRoutes.POST("/", postHandlers.Add.AddPost)
	postRoutes.POST("/:id/like", postHandlers.Like.LikePost)
	postRoutes.DELETE("/:id", postHandlers.Delete.DeletePost)
}
