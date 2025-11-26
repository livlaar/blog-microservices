package main

import (
	"log"
	"net"

	"github.com/livlaar/blog-microservices/comments/internal/controller"
	commentsgrpc "github.com/livlaar/blog-microservices/comments/internal/grpc"
	"github.com/livlaar/blog-microservices/comments/internal/repository"
	pb "github.com/livlaar/blog-microservices/shared/proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Error al iniciar listener:", err)
	}

	// Repositorio en memoria
	repo := repository.NewCommentsMemoryRepository()

	// Controller correcto (singular)
	ctrl := controller.NewCommentController(repo)

	// Servidor gRPC con controller correcto
	server := commentsgrpc.NewCommentsGRPCServer(ctrl)

	grpcServer := grpc.NewServer()
	pb.RegisterCommentsServer(grpcServer, server)

	log.Println("Comments service escuchando en :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Error al iniciar servidor gRPC:", err)
	}
}
