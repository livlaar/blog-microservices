package models

type Post struct {
	ID        string
	Title     string
	Content   string
	AuthorID  string
	CreatedAt string
}

type Comment struct {
	ID        string
	PostID    string
	AuthorID  string
	Content   string
	CreatedAt string
}
