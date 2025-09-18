package controller

import (
	"fmt"

	"github.com/livlaar/blog-microservices/posts/internal/gateway"
	"github.com/livlaar/blog-microservices/posts/internal/repository"
	model "github.com/livlaar/blog-microservices/shared/models"
)

type PostController struct {
	repo       repository.PostRepository
	commentsGw *gateway.CommentsGateway
	usersGw    *gateway.UsersGateway
}

func NewPostController(r repository.PostRepository, cg *gateway.CommentsGateway, ug *gateway.UsersGateway) *PostController {
	return &PostController{
		repo:       r,
		commentsGw: cg,
		usersGw:    ug,
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
		err := c.usersGw.CreateUser(model.User{
			ID:   post.AuthorID,
			Name: fmt.Sprintf("User-%s", post.AuthorID),
		})
		if err != nil {
			return fmt.Errorf("no se pudo crear el usuario: %w", err)
		}
	}

	return c.repo.Create(post)
}
