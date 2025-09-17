package repository

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/livlaar/blog-microservices/users/models/model"
)

type FileRepo struct {
	file string
	mu   sync.Mutex
	data map[string]model.User
}

func NewFileRepo(filename string) (*FileRepo, error) {
	r := &FileRepo{
		file: "/app/data/users.json",
		data: make(map[string]model.User),
	}

	// Si existe el archivo, cargar datos
	if _, err := os.Stat(filename); err == nil {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(bytes, &r.data)
	}
	r.load()
	return r, nil
}

func (r *FileRepo) saveToFile() error {
	bytes, err := json.MarshalIndent(r.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.file, bytes, 0644)
}

func (r *FileRepo) GetByID(id string) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.data[id]
	if !ok {
		return model.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}

func (r *FileRepo) Create(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[user.ID] = user
	return r.saveToFile()
}

func (r *FileRepo) load() {
	if _, err := os.Stat(r.file); os.IsNotExist(err) {
		// archivo no existe, inicializamos vacío
		r.data = make(map[string]model.User)
		return
	}

	bytes, err := os.ReadFile(r.file)
	if err != nil {
		log.Println("Error leyendo users.json:", err)
		r.data = make(map[string]model.User)
		return
	}

	if err := json.Unmarshal(bytes, &r.data); err != nil {
		log.Println("Error parseando users.json:", err)
		r.data = make(map[string]model.User)
	}
}

type UserRepository interface {
	GetByID(id string) (model.User, error)
	Create(user model.User) error
}
