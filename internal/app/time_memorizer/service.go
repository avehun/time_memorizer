package service

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/radiance822/time_memorizer/internal/app/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	StartGrpcServer = startGrpcServer
	StartHttpServer = startHttpServer
)

func StartServers(grpcString, httpString string, storage *model.CategoryStorage) {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting gRPC server")
		startGrpcServer(grpcString, storage)
		log.Println("gRPC server gracefully shut down")
	}()
	go startHttpServer(httpString, storage)
	wg.Wait()
}
func startGrpcServer(adress string, storage *model.CategoryStorage) {
	list, err := net.Listen("tcp", adress)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	RegisterTimeMemorizerServer(server, &timeMemorizerServer{Storage: storage})
	reflection.Register(server)

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt
		server.GracefulStop()
	}()

	server.Serve(list)
}
func startHttpServer(adress string, storage *model.CategoryStorage) {
	list, err := net.Listen("tcp", adress)
	if err != nil {
		return
	}
	mux := InitHttpHandler(storage)
	err = http.Serve(list, mux)
	if err != nil {
		log.Fatal("error serving http")
	}
}
