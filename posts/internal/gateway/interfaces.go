package gateway

import "github.com/livlaar/blog-microservices/shared/models"

type CommentsGateway interface {
	GetCommentsByPost(postID string) ([]models.Comment, error)
}

type UsersGateway interface {
	CheckUserExists(userID string) (bool, error)
}
