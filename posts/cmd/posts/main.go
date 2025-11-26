package main

import (
	"log"
	"net"

	"github.com/livlaar/blog-microservices/posts/internal/server"
	pb "github.com/livlaar/blog-microservices/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Posts service listening on :50052")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostsServer(s, server.NewPostsServer())

	reflection.Register(s)

	log.Println("Posts service listening on :50052")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
