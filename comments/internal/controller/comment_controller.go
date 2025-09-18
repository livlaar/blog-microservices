package controller

import (
	"fmt"

	"github.com/livlaar/blog-microservices/comments/internal/gateway"
	"github.com/livlaar/blog-microservices/comments/internal/repository"
	model "github.com/livlaar/blog-microservices/shared/models"
)

type CommentController struct {
	repo    repository.CommentRepository
	postsGw *gateway.PostsGateway
}

func NewCommentController(repo repository.CommentRepository, postsGw *gateway.PostsGateway) *CommentController {
	return &CommentController{repo: repo, postsGw: postsGw}
}

func (c *CommentController) GetCommentByID(id string) (model.Comment, error) {
	return c.repo.GetByID(id)
}

func (c *CommentController) GetCommentsByPost(postID string) ([]model.Comment, error) {
	return c.repo.GetByPostID(postID)
}

func (c *CommentController) CreateComment(comment model.Comment) error {
	exists, err := c.postsGw.CheckPostExists(comment.PostID)
	if err != nil {
		return fmt.Errorf("error verificando post: %w", err)
	}
	if !exists {
		return fmt.Errorf("post no encontrado")
	}

	return c.repo.Create(comment)
}
