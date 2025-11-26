package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	blogpb "github.com/livlaar/blog-microservices/shared/proto"
)

var (
	usersClient    blogpb.UsersClient
	postsClient    blogpb.PostsClient
	commentsClient blogpb.CommentsClient
)

func main() {
	fmt.Println("Starting API Gateway on port 8080...")

	// USERS service
	usersConn, err := grpc.Dial("users-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to Users service: %v", err)
	}
	usersClient = blogpb.NewUsersClient(usersConn)

	// POSTS service
	postsConn, err := grpc.Dial("posts-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to Posts service: %v", err)
	}
	postsClient = blogpb.NewPostsClient(postsConn)

	// COMMENTS service
	commentsConn, err := grpc.Dial("comments-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to Comments service: %v", err)
	}
	commentsClient = blogpb.NewCommentsClient(commentsConn)

	// Routes
	http.HandleFunc("/users/", handleGetUser)
	http.HandleFunc("/posts/", handleGetPost)
	http.HandleFunc("/comments/", handleGetComments)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):] // /users/123

	resp, err := usersClient.GetUser(context.Background(), &blogpb.GetUserRequest{
		Id: id,
	})
	if err != nil {
		http.Error(w, "Error fetching user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User: %v", resp)
}

func handleGetPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/posts/"):]

	resp, err := postsClient.GetPostWithComments(context.Background(), &blogpb.GetPostRequest{
		Id: id,
	})
	if err != nil {
		http.Error(w, "Error fetching post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Post with comments: %v", resp)
}

func handleGetComments(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post")

	resp, err := commentsClient.GetCommentsByPost(context.Background(), &blogpb.GetCommentsRequest{
		PostId: postID,
	})
	if err != nil {
		http.Error(w, "Error fetching comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Comments: %v", resp)
}
