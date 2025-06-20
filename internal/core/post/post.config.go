package post

import (
	"github.com/joaopedro489/back-golang/internal/core/post/handlers"
	"github.com/joaopedro489/back-golang/internal/core/post/handlers/controllers"
	"github.com/joaopedro489/back-golang/internal/core/post/repository"
	usecases "github.com/joaopedro489/back-golang/internal/core/post/use-cases"
	"gorm.io/gorm"
)

func Config(db *gorm.DB) *handlers.PostHandlers {
	postRepo := repository.NewPostRepository(db)
	userRepot := repository.NewUserRepository(db)

	addPostUC := usecases.NewAddPost(postRepo, userRepot)
	addPost := controllers.NewAddPostController(*addPostUC)

	readPostUC := usecases.NewReadPost(postRepo, userRepot)
	readPost := controllers.NewReadPostController(*readPostUC)

	deletePostUC := usecases.NewDeletePost(postRepo)
	deletePost := controllers.NewDeletePostController(*deletePostUC)

	browsePostsUC := usecases.NewBrowsePost(postRepo)
	browsePosts := controllers.NewBrowsePostController(*browsePostsUC)

	likeUc := usecases.NewLikePost(postRepo)
	likePost := controllers.NewLikePostController(*likeUc)

	postHandler := &handlers.PostHandlers{
		Add:    addPost,
		Read:   readPost,
		Delete: deletePost,
		Browse: browsePosts,
		Like:   likePost,
	}

	return postHandler
}
