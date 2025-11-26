package mapper

import (
	"github.com/livlaar/blog-microservices/users/domain"

	proto "github.com/livlaar/blog-microservices/shared/proto"
)

func GrpcToDomain(u *proto.UserDTO) *domain.User {
	return &domain.User{
		ID:   u.Id,
		Name: u.Name,
	}
}
