package repository

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/livlaar/blog-microservices/comments/models/model"
)

type CommentRepository interface {
	GetByID(id string) (model.Comment, error)
	GetByPostID(postID string) ([]model.Comment, error)
	Create(comment model.Comment) error
}

type FileRepo struct {
	file string
	data map[string]model.Comment
}

func NewFileRepo() *FileRepo {
	r := &FileRepo{
		file: "/app/data/comments.json",
		data: make(map[string]model.Comment),
	}
	r.load()
	return r
}

func (r *FileRepo) load() {
	if _, err := os.Stat(r.file); os.IsNotExist(err) {
		r.data = make(map[string]model.Comment)
		return
	}

	bytes, err := os.ReadFile(r.file)
	if err != nil {
		log.Println("Error leyendo comments.json:", err)
		return
	}

	if err := json.Unmarshal(bytes, &r.data); err != nil {
		log.Println("Error parseando comments.json:", err)
		r.data = make(map[string]model.Comment)
	}
}

func (r *FileRepo) save() {
	bytes, err := json.MarshalIndent(r.data, "", "  ")
	if err != nil {
		log.Println("Error serializando comments.json:", err)
		return
	}
	if err := os.WriteFile(r.file, bytes, 0644); err != nil {
		log.Println("Error escribiendo comments.json:", err)
	}
}

func (r *FileRepo) GetByID(id string) (model.Comment, error) {
	c, ok := r.data[id]
	if !ok {
		return model.Comment{}, errors.New("comentario no encontrado")
	}
	return c, nil
}

func (r *FileRepo) GetByPostID(postID string) ([]model.Comment, error) {
	list := []model.Comment{}
	for _, c := range r.data {
		if c.PostID == postID {
			list = append(list, c)
		}
	}
	return list, nil
}

func (r *FileRepo) Create(comment model.Comment) error {
	if _, exists := r.data[comment.ID]; exists {
		return errors.New("comentario ya existe")
	}
	r.data[comment.ID] = comment
	r.save()
	return nil
}
