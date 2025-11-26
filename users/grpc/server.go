package grpcserver

import (
	"context"

	proto "github.com/livlaar/blog-microservices/shared/proto"
	"github.com/livlaar/blog-microservices/users/mapper"
	"github.com/livlaar/blog-microservices/users/service"
)

type UserGrpcServer struct {
	proto.UnimplementedUsersServer
	svc *service.UserService
}

func New(s *service.UserService) *UserGrpcServer {
	return &UserGrpcServer{svc: s}
}

func (g *UserGrpcServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := g.svc.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	// 👇 conversión: domain → protobuf
	return &proto.GetUserResponse{
		User: mapper.DomainToGrpc(user),
	}, nil
}

func (g *UserGrpcServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	// 👇 conversión: protobuf → domain
	u := mapper.GrpcToDomain(req.User)

	err := g.svc.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{Ok: true}, nil
}
