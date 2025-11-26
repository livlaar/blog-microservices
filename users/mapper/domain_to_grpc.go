package mapper

import (
	proto "github.com/livlaar/blog-microservices/shared/proto"
	"github.com/livlaar/blog-microservices/users/domain"
)

func DomainToGrpc(u *domain.User) *proto.UserDTO {
	return &proto.UserDTO{
		Id:   u.ID,
		Name: u.Name,
	}
}
