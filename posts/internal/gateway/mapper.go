package gateway

import (
	model "github.com/livlaar/blog-microservices/shared/models"
	pb "github.com/livlaar/blog-microservices/shared/proto"
)

// Exported helpers (mayúscula) para evitar problemas de visibilidad entre archivos.
func CommentProtoToDomain(c *pb.CommentDTO) model.Comment {
	return model.Comment{
		ID:        c.Id,
		PostID:    c.PostId,
		AuthorID:  c.AuthorId,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

func CommentsProtoListToDomain(list []*pb.CommentDTO) []model.Comment {
	out := make([]model.Comment, 0, len(list))
	for _, c := range list {
		out = append(out, CommentProtoToDomain(c))
	}
	return out
}

func UserDomainToProto(u model.User) *pb.UserDTO {
	return &pb.UserDTO{
		Id:   u.ID,
		Name: u.Name,
	}
}
