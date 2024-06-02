package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/radiance822/time_memorizer/internal/app/model"
)

type Request struct {
	Category string `json:"category"`
	Time     int    `json:"time"`
}

type httpServer struct {
	storage *model.CategoryStorage
}

func InitHttpHandler(storage *model.CategoryStorage) *http.ServeMux {
	server := httpServer{storage: storage}
	mux := http.NewServeMux()
	mux.HandleFunc("/AddTime", server.AddCategoryTime)
	mux.HandleFunc("/SubstractTime", server.SubstractCategoryTime)
	mux.HandleFunc("/ShowTime", server.ShowCategoryTime)
	return mux
}

func (s httpServer) AddCategoryTime(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Wrong request", http.StatusBadRequest)
	}
	s.storage.Add(request.Category, request.Time)
	fmt.Fprintf(w, "Added %v for category: %s", request.Time, request.Category)
}
func (s httpServer) SubstractCategoryTime(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Wrong request", http.StatusBadRequest)
	}
	err = s.storage.Subtract(request.Category, request.Time)
	if err != nil {
		http.Error(w, "Unable to substract", 400)
	}
	fmt.Fprint(w, "Succesfully substracted")
}
func (s httpServer) ShowCategoryTime(w http.ResponseWriter, r *http.Request) {
	type showRequest struct {
		Category string `json:"category"`
	}
	request := showRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", 400)
	}
	res := s.storage.Load(request.Category)
	fmt.Fprintf(w, "Category %s contains %v time", request.Category, res)
}
