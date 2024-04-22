package app

import (
	"log"
	"net"
	"net/http"

	"github.com/radiance822/time_memorizer/internal/app/model"
	service "github.com/radiance822/time_memorizer/internal/app/time_memorizer"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app App) Run() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Unable to listen this port")
	}

	mux := cmux.New(listener)
	grpcListener := mux.Match(cmux.HTTP2())
	httpListener := mux.Match(cmux.HTTP1Fast())

	go serveGrpc(grpcListener)
	go serveHttp(httpListener)

	if err := mux.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func serveGrpc(listener net.Listener) {
	grpcServer := grpc.NewServer()
	timeMemServ := service.Server{Storage: model.CategoryStorage{}}
	service.RegisterTimeMemorizerServer(grpcServer, timeMemServ)
	reflection.Register(grpcServer)
	err := grpcServer.Serve(listener)
	if err != nil {
		log.Print("error serving grpc")
	}

}

func serveHttp(listener net.Listener) {
	mux := service.InitHttpHandler()
	if err := http.Serve(listener, mux); err != nil {
		log.Fatal("Failed to serve http")
	}
}
