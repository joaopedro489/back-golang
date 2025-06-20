package entity

type Post struct {
	Id      int
	Title   string
	Content string
	UserId  int
	Likes   []Like
}

func NewPost(title, content string, userId int) *Post {
	return &Post{
		Title:   title,
		Content: content,
		UserId:  userId,
	}
}
