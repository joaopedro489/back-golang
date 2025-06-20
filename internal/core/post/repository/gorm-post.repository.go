package repository

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/entity"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
	"github.com/joaopedro489/back-golang/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(post entity.Post) (*entity.Post, error) {
	dbPost := domain.PostDomainToModel(post)
	if err := r.db.Create(&dbPost).Error; err != nil {
		return nil, err
	}
	createdPost := domain.PostModelToDomain(dbPost)
	return &createdPost, nil
}

func (r *PostRepository) GetPostById(id int) (*entity.Post, error) {
	var dbPost models.Post
	if err := r.db.Preload("Likes").First(&dbPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	post := domain.PostModelToDomain(dbPost)
	return &post, nil
}

func (r *PostRepository) DeletePost(id int) error {
	if err := r.db.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) BrowsePosts(filters repository.Filters) ([]entity.Post, error) {
	var dbPosts []models.Post

	query := r.db.Model(&models.Post{})
	query = query.Preload("Likes")
	query = query.Limit(filters.Limit).Offset(filters.Offset)
	if filters.UserId != nil {
		query = query.Where("user_id = ?", *filters.UserId)
	}

	if err := query.Find(&dbPosts).Error; err != nil {
		return nil, err
	}
	posts := domain.PostModelsToDomain(dbPosts)
	return posts, nil
}

func (r *PostRepository) LikePost(postId int, userId int) error {
	like := models.Like{
		PostId: postId,
		UserId: userId,
	}

	if err := r.db.Create(&like).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) CheckAlreadyLiked(postId int, userId int) (bool, error) {
	var like models.Like
	if err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
