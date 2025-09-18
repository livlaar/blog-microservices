package controller

import (
	model "github.com/livlaar/blog-microservices/shared/models"
	"github.com/livlaar/blog-microservices/users/internal/repository"
)

type UserController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) GetUserByID(id string) (model.User, error) {
	return c.repo.GetByID(id)
}

func (c *UserController) CreateUser(user model.User) error {
	return c.repo.Create(user)
}
