package controller

import (
	"fmt"

	"github.com/livlaar/blog-microservices/posts/internal/gateway"
	"github.com/livlaar/blog-microservices/posts/internal/repository"
	"github.com/livlaar/blog-microservices/shared/models"
)

type PostController struct {
	repo       repository.PostRepository
	commentsGw gateway.CommentsGateway
	usersGw    gateway.UsersGateway
}

func NewPostController(r repository.PostRepository, cg gateway.CommentsGateway, ug gateway.UsersGateway) *PostController {
	return &PostController{
		repo:       r,
		commentsGw: cg,
		usersGw:    ug,
	}
}

func (c *PostController) GetPostWithComments(postID string) (models.Post, []models.Comment, error) {
	post, err := c.repo.GetByID(postID)
	if err != nil {
		return models.Post{}, nil, err
	}

	comments, err := c.commentsGw.GetCommentsByPost(postID)
	if err != nil {
		return post, nil, err
	}

	return post, comments, nil
}

func (c *PostController) CreatePost(post models.Post) error {
	exists, err := c.usersGw.CheckUserExists(post.AuthorID)
	if err != nil {
		return fmt.Errorf("error verificando usuario: %w", err)
	}

	if !exists {
		return fmt.Errorf("el usuario %s no existe", post.AuthorID)
	}

	return c.repo.Create(post)
}
