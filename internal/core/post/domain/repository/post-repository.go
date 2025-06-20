package repository

import "github.com/joaopedro489/back-golang/internal/core/post/domain/entity"

type Filters struct {
	UserId *int
	Limit  int
	Offset int
}

type PostResult struct {
	Post *entity.Post
	Err  error
}

type IPostRepository interface {
	GetPostById(id int) (*entity.Post, error)
	BrowsePosts(filters Filters) ([]entity.Post, error)
	CreatePost(post entity.Post) (*entity.Post, error)
	LikePost(postId int, userId int) error
	DeletePost(id int) error
	CheckAlreadyLiked(postId int, userId int) (bool, error)
}
