package models

// Modelo interno para un Post
type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  string `json:"author_id"`
	CreatedAt string `json:"created_at"`
}

// Modelo interno para un Comment
type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Modelo interno para un User
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
