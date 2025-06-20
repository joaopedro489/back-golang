package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/users/handlers"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handlers.UserHandlers) {
	r.POST("/users", userHandler.Add.AddUser)
	r.GET("/users/:id", userHandler.Read.ReadUser)
	r.DELETE("/users/:id", userHandler.Delete.DeleteUser)
	r.GET("/users", userHandler.Browse.BrowseUsers)
	r.PATCH("/users/:id", userHandler.Edit.EditUser)

	authGroup := r.Group("/api")
	authGroup.Use(auth.JWTMiddleware())
	authGroup.GET("/details", userHandler.GetDetails.GetDetails)
}
