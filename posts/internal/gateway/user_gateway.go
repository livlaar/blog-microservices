package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersHTTPGateway struct {
	baseURL string
	client  *http.Client
}

func NewUsersHTTPGateway(baseURL string) *UsersHTTPGateway {
	return &UsersHTTPGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (g *UsersHTTPGateway) GetUserByID(userID string) (UserDTO, error) {
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

func (g *UsersHTTPGateway) CheckUserExists(userID string) (bool, error) {
	resp, err := g.client.Get(fmt.Sprintf("%s/users/%s", g.baseURL, userID))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return resp.StatusCode == http.StatusOK, nil
}
