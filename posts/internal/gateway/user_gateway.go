package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	model "github.com/livlaar/blog-microservices/shared/models"
)

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersGateway struct {
	baseURL string
	client  *http.Client
}

func NewUsersGateway(baseURL string) *UsersGateway {
	return &UsersGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (g *UsersGateway) GetUserByID(userID string) (UserDTO, error) {
	url := fmt.Sprintf("%s/users/%s", g.baseURL, userID)
	resp, err := g.client.Get(url)
	if err != nil {
		return UserDTO{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserDTO{}, fmt.Errorf("users service returned %d", resp.StatusCode)
	}

	var user UserDTO
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return UserDTO{}, err
	}

	return user, nil
}

func (g *UsersGateway) CheckUserExists(userID string) (bool, error) {
	resp, err := g.client.Get(fmt.Sprintf("%s/users/%s", g.baseURL, userID))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return true, nil
}

func (g *UsersGateway) CreateUser(user model.User) error {
	data, _ := json.Marshal(user)
	resp, err := g.client.Post(fmt.Sprintf("%s/users", g.baseURL), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
