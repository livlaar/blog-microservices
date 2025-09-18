package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"

	"github.com/livlaar/blog-microservices/comments/internal/controller"
	"github.com/livlaar/blog-microservices/comments/internal/gateway"
	"github.com/livlaar/blog-microservices/comments/internal/handler"
	"github.com/livlaar/blog-microservices/comments/internal/repository"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "http://consul:8500"
	}

	repo := repository.NewFileRepo()

	postsAddr := "http://localhost:8002"
	postsGw := gateway.NewPostsGateway(postsAddr)

	ctrl := controller.NewCommentController(repo, postsGw)
	h := handler.NewCommentHandler(ctrl)

	r := mux.NewRouter()
	r.HandleFunc("/comments/{id}", h.GetComment).Methods("GET")
	r.HandleFunc("/posts/{id}/comments", h.GetCommentsByPost).Methods("GET")
	r.HandleFunc("/comments", h.CreateComment).Methods("POST")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err == nil {
		reg := &api.AgentServiceRegistration{
			ID:   "comments-service",
			Name: "comments",
			Port: 8003,
			Check: &api.AgentServiceCheck{
				HTTP:     "http://blog-microservices-comments-1:8003/health",
				Interval: "10s",
				Timeout:  "2s",
			},
		}
		client.Agent().ServiceRegister(reg)
	}

	log.Printf("Comments service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
