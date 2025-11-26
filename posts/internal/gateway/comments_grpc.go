package gateway

import (
	"context"
	"time"

	"github.com/livlaar/blog-microservices/shared/models"
	pb "github.com/livlaar/blog-microservices/shared/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type commentsGatewayGRPC struct {
	client pb.CommentsClient
}

func NewCommentsGatewayGRPC(addr string) (CommentsGateway, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewCommentsClient(conn)
	return &commentsGatewayGRPC{client: client}, nil
}

func (g *commentsGatewayGRPC) GetCommentsByPost(postID string) ([]models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := g.client.GetCommentsByPost(ctx, &pb.GetCommentsRequest{
		PostId: postID,
	})
	if err != nil {
		return nil, err
	}

	return CommentsProtoListToDomain(resp.Comments), nil
}
