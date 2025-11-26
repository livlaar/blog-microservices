package gateway

import (
	"context"
	"time"

	pb "github.com/livlaar/blog-microservices/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type usersGatewayGRPC struct {
	client pb.UsersClient
}

func NewUsersGatewayGRPC(addr string) (UsersGateway, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUsersClient(conn)
	return &usersGatewayGRPC{client: client}, nil
}

func (g *usersGatewayGRPC) CheckUserExists(userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := g.client.GetUser(ctx, &pb.GetUserRequest{Id: userID})
	if err != nil {
		// Si el servicio responde NOT_FOUND, se interpreta como "no existe"
		return false, nil
	}

	// existe si resp.User != nil
	return resp.User != nil && resp.User.Id != "", nil
}
