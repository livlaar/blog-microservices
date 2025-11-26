package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"

	"github.com/livlaar/blog-microservices/users/internal/controller"
	grpcserver "github.com/livlaar/blog-microservices/users/internal/grpc"
	"github.com/livlaar/blog-microservices/users/internal/repository"

	pb "github.com/livlaar/blog-microservices/shared/proto"
	"google.golang.org/grpc"
)

func main() {

	// === CONFIG PUERTO ===
	port := "50051"

	// === CONSUL ===
	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "http://consul:8500"
	}

	// === REPO ===
	repo, _ := repository.NewFileRepo("/app/data/users.json")
	ctrl := controller.NewUserController(repo)

	// === GRPC SERVER ===
	grpcSrv := grpcserver.NewUserGRPCServer(ctrl)
	grpcServer := grpc.NewServer()

	pb.RegisterUsersServer(grpcServer, grpcSrv)

	// Escuchar puerto
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error listening on port %s: %v", port, err)
	}

	log.Printf("Users gRPC service running on port %s", port)

	// Correr GRPC en goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// === HTTP healthcheck (obligatorio para Consul) ===
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	go func() {
		log.Println("Healthcheck HTTP server running on :8080")
		http.ListenAndServe(":8080", r)
	}()

	// === REGISTER SERVICE IN CONSUL ===
	config := api.DefaultConfig()
	config.Address = consulAddr

	client, err := api.NewClient(config)
	if err == nil {
		reg := &api.AgentServiceRegistration{
			ID:   "users",
			Name: "users",
			Port: 50051,
			Check: &api.AgentServiceCheck{
				HTTP:     "http://users:8080/health",
				Interval: "10s",
				Timeout:  "2s",
			},
		}
		client.Agent().ServiceRegister(reg)
		log.Println("Registered users service in Consul")
	} else {
		log.Println("WARNING: Could not register service in Consul:", err)
	}

	// Bloquea la ejecución (el programa no termina)
	select {}
}
