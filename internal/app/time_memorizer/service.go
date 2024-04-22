package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/radiance822/time_memorizer/internal/app/model"
	grpc "google.golang.org/grpc"
)

type Server struct {
	UnimplementedTimeMemorizerServer
	Storage model.CategoryStorage
}

func (s Server) AddTime(ctx context.Context, in *CategoryAndTime) (*Message, error) {
	s.Storage.Add(in.Category, int(in.TimeSpent))
	return &Message{
		Body: "Time added",
	}, nil
}

func (s Server) SubstractTime(ctx context.Context, in *CategoryAndTime) (*Message, error) {
	s.Storage.Subtract(in.Category, int(in.TimeSpent))
	return &Message{
		Body: "Time substracted",
	}, nil
}
func InitHttpHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/AddTime", AddCategoryTime)
	return mux
}

func AddCategoryTime(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Category string `json:"category"`
		Time     int    `json:"time"`
	}
	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Wrong request", http.StatusBadRequest)
	}
	converted := CategoryAndTime{
		Category:  request.Category,
		TimeSpent: int32(request.Time),
	}
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to dial gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := NewTimeMemorizerClient(conn)

	response, err := client.AddTime(r.Context(), &converted)
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
