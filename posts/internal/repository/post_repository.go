package repository

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	model "github.com/livlaar/blog-microservices/shared/models"
)

type PostRepository interface {
	GetByID(id string) (model.Post, error)
	Create(post model.Post) error
	GetAll() []model.Post
}

type FileRepo struct {
	mu   sync.Mutex
	file string
	data map[string]model.Post
}

func NewFileRepo() *FileRepo {
	r := &FileRepo{
		file: "/app/data/posts.json",
		data: make(map[string]model.Post),
	}
	r.load()
	return r
}

func (r *FileRepo) load() {
	if _, err := os.Stat(r.file); os.IsNotExist(err) {
		r.data = make(map[string]model.Post)
		return
	}

	bytes, err := os.ReadFile(r.file)
	if err != nil {
		log.Println("Error leyendo posts.json:", err)
		r.data = make(map[string]model.Post)
		return
	}

	if err := json.Unmarshal(bytes, &r.data); err != nil {
		log.Println("Error parseando posts.json:", err)
		r.data = make(map[string]model.Post)
	}
}

func (r *FileRepo) save() {
	bytes, err := json.MarshalIndent(r.data, "", "  ")
	if err != nil {
		log.Println("Error serializando posts.json:", err)
		return
	}

	if err := os.WriteFile(r.file, bytes, 0644); err != nil {
		log.Println("Error guardando posts.json:", err)
	}
}

func (r *FileRepo) GetByID(id string) (model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post, ok := r.data[id]
	if !ok {
		return model.Post{}, ErrPostNotFound
	}
	return post, nil
}

func (r *FileRepo) Create(post model.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[post.ID] = post
	r.save()
	return nil
}

func (r *FileRepo) GetAll() []model.Post {
	r.mu.Lock()
	defer r.mu.Unlock()

	posts := make([]model.Post, 0, len(r.data))
	for _, p := range r.data {
		posts = append(posts, p)
	}
	return posts
}

// Error de post no encontrado
var ErrPostNotFound = &RepositoryError{"post no encontrado"}

type RepositoryError struct {
	Msg string
}

func (e *RepositoryError) Error() string {
	return e.Msg
}
