package repository

import (
	"errors"

	model "github.com/livlaar/blog-microservices/shared/models"
)

type commentsMemoryRepository struct {
	data []model.Comment
}

func NewCommentsMemoryRepository() CommentRepository {
	return &commentsMemoryRepository{
		data: make([]model.Comment, 0),
	}
}

func (r *commentsMemoryRepository) GetByID(id string) (model.Comment, error) {
	for _, c := range r.data {
		if c.ID == id {
			return c, nil
		}
	}
	return model.Comment{}, errors.New("comentario no encontrado")
}

func (r *commentsMemoryRepository) GetByPostID(postID string) ([]model.Comment, error) {
	out := []model.Comment{}
	for _, c := range r.data {
		if c.PostID == postID {
			out = append(out, c)
		}
	}
	return out, nil
}

func (r *commentsMemoryRepository) Create(c model.Comment) error {
	// validar duplicado
	for _, existing := range r.data {
		if existing.ID == c.ID {
			return errors.New("comentario ya existe")
		}
	}

	r.data = append(r.data, c)
	return nil
}
