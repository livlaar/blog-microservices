package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CommentDTO representa el comentario recibido del servicio Comments
type CommentDTO struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// CommentsGateway hace llamadas HTTP al servicio Comments
type CommentsGateway struct {
	baseURL string
	client  *http.Client
}

// NewCommentsGateway crea un nuevo gateway
func NewCommentsGateway(baseURL string) *CommentsGateway {
	return &CommentsGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

// GetCommentsByPost obtiene los comentarios de un post
func (g *CommentsGateway) GetCommentsByPost(postID string) ([]CommentDTO, error) {
	url := fmt.Sprintf("%s/posts/%s/comments", g.baseURL, postID)
	resp, err := g.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("comments service returned %d", resp.StatusCode)
	}

	var list []CommentDTO
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}

	return list, nil
}
