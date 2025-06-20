package output

type PostResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserName string `json:"userName"`
	Likes    int    `json:"likes"`
}

func NewPostResponse(id int, title, content, userName string, likes int) *PostResponse {
	return &PostResponse{
		Id:       id,
		Title:    title,
		Content:  content,
		UserName: userName,
		Likes:    likes,
	}
}
