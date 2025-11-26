package controller

import (
	"github.com/google/uuid"
	"github.com/livlaar/blog-microservices/comments/internal/repository"
	model "github.com/livlaar/blog-microservices/shared/models"
)

type CommentController struct {
	repo repository.CommentRepository
}

func NewCommentController(r repository.CommentRepository) *CommentController {
	return &CommentController{repo: r}
}

func (c *CommentController) GetCommentByID(id string) (model.Comment, error) {
	return c.repo.GetByID(id)
}

func (c *CommentController) GetCommentsByPost(postID string) ([]model.Comment, error) {
	return c.repo.GetByPostID(postID)
}

func (c *CommentController) CreateComment(comment model.Comment) error {
	// Asigna ID si viene vacío
	if comment.ID == "" {
		comment.ID = uuid.NewString()
	}
	return c.repo.Create(comment)
}
