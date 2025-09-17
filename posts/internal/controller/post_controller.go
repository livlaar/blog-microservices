package controller

import (
	"fmt"

	"github.com/livlaar/blog-microservices/posts/internal/gateway"
	"github.com/livlaar/blog-microservices/posts/internal/repository"
	"github.com/livlaar/blog-microservices/posts/models/model"
)

type PostController struct {
	repo       repository.PostRepository
	commentsGw *gateway.CommentsGateway
}

func NewPostController(r repository.PostRepository, cg *gateway.CommentsGateway) *PostController {
	return &PostController{
		repo:       r,
		commentsGw: cg,
	}
}

func (c *PostController) GetPostWithComments(postID string) (model.Post, []gateway.CommentDTO, error) {
	post, err := c.repo.GetByID(postID)
	if err != nil {
		return model.Post{}, nil, err
	}

	comments, err := c.commentsGw.GetCommentsByPost(postID)
	if err != nil {
		return post, nil, nil
	}
	return post, comments, nil
}

func (c *PostController) CreatePost(post model.Post) error {
	exists, err := c.usersGw.CheckUserExists(post.AuthorID)
	if err != nil {
		return fmt.Errorf("error verificando usuario: %w", err)
	}

	if !exists {
		// Crear usuario automáticamente
		err := c.usersGw.CreateUser(model.User{
			ID:   post.AuthorID,
			Name: fmt.Sprintf("User-%s", post.AuthorID), // o tomas el valor del JSON
		})
		if err != nil {
			return fmt.Errorf("no se pudo crear el usuario: %w", err)
		}
	}

	return c.repo.Create(post)
}
