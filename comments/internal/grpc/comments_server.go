package grpc

import (
	"context"

	"github.com/livlaar/blog-microservices/comments/internal/controller"
	"github.com/livlaar/blog-microservices/comments/internal/gateway"
	pb "github.com/livlaar/blog-microservices/shared/proto"
)

type CommentsGRPCServer struct {
	pb.UnimplementedCommentsServer
	ctrl *controller.CommentController
}

func NewCommentsGRPCServer(c *controller.CommentController) *CommentsGRPCServer {
	return &CommentsGRPCServer{ctrl: c}
}

func (s *CommentsGRPCServer) GetCommentsByPost(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	comms, err := s.ctrl.GetCommentsByPost(req.PostId)
	if err != nil {
		return nil, err
	}

	protoList := make([]*pb.CommentDTO, 0, len(comms))
	for _, c := range comms {
		protoList = append(protoList, gateway.CommentDomainToProto(c))
	}

	return &pb.GetCommentsResponse{Comments: protoList}, nil
}

func (s *CommentsGRPCServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	domain := gateway.CommentProtoToDomain(req.Comment)
	err := s.ctrl.CreateComment(domain)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{Ok: true}, nil
}
