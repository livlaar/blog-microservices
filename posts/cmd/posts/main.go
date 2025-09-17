package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"

	"github.com/livlaar/blog-microservices/posts/internal/controller"
	"github.com/livlaar/blog-microservices/posts/internal/gateway"
	"github.com/livlaar/blog-microservices/posts/internal/handler"
	"github.com/livlaar/blog-microservices/posts/internal/repository"
)

func main() {
	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	// Dirección de Consul
	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "http://consul:8500"
	}

	repo := repository.NewFileRepo()
	commentsGw := gateway.NewCommentsGateway("http://comments:8003")
	ctrl := controller.NewPostController(repo, commentsGw)
	h := handler.NewPostHandler(ctrl)

	r := mux.NewRouter()
	r.HandleFunc("/posts/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/posts", h.CreatePost).Methods("POST")

	// Health endpoint para Consul
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Registro en Consul
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		log.Println("Error creando cliente de Consul:", err)
	} else {
		registration := &api.AgentServiceRegistration{
			ID:   "posts-service",
			Name: "posts",
			Port: 8002,
			Check: &api.AgentServiceCheck{
				HTTP:     "http://posts:8002/health",
				Interval: "10s",
				Timeout:  "2s",
			},
		}
		err = client.Agent().ServiceRegister(registration)
		if err != nil {
			log.Println("Error registrando Posts en Consul:", err)
		} else {
			log.Println("Posts registrado en Consul")
		}
	}

	// Health endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Inicia servidor
	log.Printf("Posts service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
