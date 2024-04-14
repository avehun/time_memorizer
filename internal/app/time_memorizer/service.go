package service

import (
	"context"
	"log"
	"net/http"
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
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("HttpHanlder response"))
		if err != nil {
			log.Print("handler error occured")
		}

	})
	return mux
}
