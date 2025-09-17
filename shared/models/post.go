package model

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
}
