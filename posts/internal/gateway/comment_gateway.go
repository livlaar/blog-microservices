package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	model "github.com/livlaar/blog-microservices/shared/models"
)

type CommentsHTTPGateway struct {
	baseURL string
	client  *http.Client
}

func NewCommentsHTTPGateway(baseURL string) *CommentsHTTPGateway {
	return &CommentsHTTPGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (g *CommentsHTTPGateway) GetCommentsByPost(postID string) ([]model.Comment, error) {
	url := fmt.Sprintf("%s/posts/%s/comments", g.baseURL, postID)
	resp, err := g.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("comments service returned %d", resp.StatusCode)
	}

	// 👇 ESTE TIPO ES CRÍTICO
	var list []model.Comment

	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}

	return list, nil
}
