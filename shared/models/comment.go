package model

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
