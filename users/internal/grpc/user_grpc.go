package grpc

import (
	"context"

	model "github.com/livlaar/blog-microservices/shared/models"
	pb "github.com/livlaar/blog-microservices/shared/proto"
	"github.com/livlaar/blog-microservices/users/internal/controller"
)

type UserGRPCServer struct {
	pb.UnimplementedUsersServer
	ctrl *controller.UserController
}

func NewUserGRPCServer(ctrl *controller.UserController) *UserGRPCServer {
	return &UserGRPCServer{ctrl: ctrl}
}

func (s *UserGRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := s.ctrl.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.UserDTO{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	u := model.User{
		ID:   req.User.Id,
		Name: req.User.Name,
	}

	err := s.ctrl.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Ok: true,
	}, nil
}
