package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"

	"github.com/livlaar/blog-microservices/users/internal/controller"
	"github.com/livlaar/blog-microservices/users/internal/handler"
	"github.com/livlaar/blog-microservices/users/internal/repository"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "http://consul:8500"
	}

	repo, _ := repository.NewFileRepo("/app/data/users.json")
	ctrl := controller.NewUserController(repo)
	h := handler.NewUserHandler(ctrl)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/health", h.Health)

	// Registro en Consul
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err == nil {
		reg := &api.AgentServiceRegistration{
			ID:   "users-service",
			Name: "users",
			Port: 8001,
			Check: &api.AgentServiceCheck{
				HTTP:     "http://users:8001/health",
				Interval: "10s",
				Timeout:  "2s",
			},
		}
		client.Agent().ServiceRegister(reg)
	}

	log.Printf("Users service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
