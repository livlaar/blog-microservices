package service

import "github.com/livlaar/blog-microservices/users/domain"

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	Save(user *domain.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(u *domain.User) error {
	return s.repo.Save(u)
}
