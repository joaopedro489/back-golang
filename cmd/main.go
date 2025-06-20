package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaopedro489/back-golang/internal/config"
	"github.com/joaopedro489/back-golang/internal/core/auth"
	authRoutes "github.com/joaopedro489/back-golang/internal/core/auth/routes"
	"github.com/joaopedro489/back-golang/internal/core/post"
	postRoutes "github.com/joaopedro489/back-golang/internal/core/post/routes"
	"github.com/joaopedro489/back-golang/internal/core/users"
	userRoutes "github.com/joaopedro489/back-golang/internal/core/users/routes"

	"github.com/joaopedro489/back-golang/internal/db"
	"github.com/joaopedro489/back-golang/internal/models"
)

func main() {
	config.Load()

	dataBase := db.NewPostgresDatabase().GetDB()
	models.Migrate(dataBase)

	userHandler := users.Config(dataBase)
	postHandler := post.Config(dataBase)
	loginHandler := auth.Config(dataBase)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	userRoutes.RegisterUserRoutes(r, userHandler)
	postRoutes.RegisterPostRoutes(r, postHandler)
	authRoutes.RegisterAuthRoutes(r, loginHandler)

	r.Run(":" + config.Env.Port)
}
