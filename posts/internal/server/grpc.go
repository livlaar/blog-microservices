package server

import (
	"context"
	"log"

	pb "github.com/livlaar/blog-microservices/shared/proto"
)

type PostsServer struct {
	pb.UnimplementedPostsServer
}

func NewPostsServer() *PostsServer {
	return &PostsServer{}
}

func (s *PostsServer) GetPostWithComments(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	log.Println("GetPostWithComments called for ID:", req.Id)

	// RESPONSE DE PRUEBA
	return &pb.GetPostResponse{
		Post: &pb.PostDTO{
			Id:        req.Id,
			Title:     "Post demo",
			Content:   "Contenido de ejemplo",
			AuthorId:  "1",
			CreatedAt: "2025-01-01",
		},
		Comments: []*pb.CommentDTO{},
	}, nil
}

func (s *PostsServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	log.Println("CreatePost called")
	return &pb.CreatePostResponse{Ok: true}, nil
}
