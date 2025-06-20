package output

import "github.com/joaopedro489/back-golang/internal/core/post/domain/entity"

type BrowsePostResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

func NewPostsResponse(posts []entity.Post) []BrowsePostResponse {
	var responses []BrowsePostResponse
	for _, post := range posts {
		postResponse := BrowsePostResponse{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
			Likes:   len(post.Likes),
		}
		responses = append(responses, postResponse)
	}
	return responses
}
