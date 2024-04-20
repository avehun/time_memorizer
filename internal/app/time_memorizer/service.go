package service

import (
	"context"
	"encoding/json"
	"net/http"

	grpc "google.golang.org/grpc"
)

type Server struct {
}

func (s Server) mustEmbedUnimplementedTimeMemorizerServer() {
	panic("unimplemented")
}

func (s Server) SimpleResponse(ctx context.Context, message *Message) (*Message, error) {
	return &Message{Body: "SimpleResponse handler"}, nil
}

func InitHttpHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/SimpleResponse", SimpleResponseHandler)
	return mux
}

func SimpleResponseHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to dial gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := NewTimeMemorizerClient(conn)

	message := &Message{
		Body: "Your message here",
	}

	response, err := client.SimpleResponse(r.Context(), message)
	if err != nil {
		http.Error(w, "Failed to call gRPC method", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
