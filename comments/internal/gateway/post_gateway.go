package gateway

import (
	"fmt"
	"net/http"
	"time"
)

// PostsGateway hace llamadas HTTP al servicio Posts
type PostsGateway struct {
	baseURL string
	client  *http.Client
}

func NewPostsGateway(baseURL string) *PostsGateway {
	return &PostsGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

// CheckPostExists valida si un post existe antes de crear un comentario
func (g *PostsGateway) CheckPostExists(postID string) (bool, error) {
	resp, err := g.client.Get(fmt.Sprintf("%s/posts/%s", g.baseURL, postID))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return resp.StatusCode == http.StatusOK, nil
}
