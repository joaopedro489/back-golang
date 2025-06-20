package domain

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/entity"
	"github.com/joaopedro489/back-golang/internal/models"
)

func PostModelToDomain(postModel models.Post) entity.Post {
	likes := make([]entity.Like, 0, len(postModel.Likes))
	for _, likeModel := range postModel.Likes {
		likes = append(likes, entity.Like{
			PostId: int(likeModel.PostId),
			UserId: int(likeModel.UserId),
		})
	}

	return entity.Post{
		Id:      int(postModel.ID),
		Title:   postModel.Title,
		Content: postModel.Content,
		UserId:  int(postModel.UserId),
		Likes:   likes,
	}
}

func PostModelsToDomain(posts []models.Post) []entity.Post {
	domainPosts := make([]entity.Post, len(posts))
	for i, postModel := range posts {
		domainPosts[i] = PostModelToDomain(postModel)
	}
	return domainPosts
}

func PostDomainToModel(post entity.Post) models.Post {
	return models.Post{
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserId,
	}
}
