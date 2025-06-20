package handlers

import "github.com/joaopedro489/back-golang/internal/core/post/handlers/controllers"

type PostHandlers struct {
	Add    *controllers.AddPostController
	Browse *controllers.BrowsePostController
	Like   *controllers.LikePostController
	Delete *controllers.DeletePostController
	Read   *controllers.ReadPostController
}

func NewPostHandlers(
	add *controllers.AddPostController,
	browse *controllers.BrowsePostController,
	like *controllers.LikePostController,
	delete *controllers.DeletePostController,
	read *controllers.ReadPostController,
) *PostHandlers {
	return &PostHandlers{
		Add:    add,
		Browse: browse,
		Like:   like,
		Delete: delete,
		Read:   read,
	}
}
