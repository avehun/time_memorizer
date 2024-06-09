package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	service "github.com/radiance822/time_memorizer/internal/app/time_memorizer"
	pb "github.com/radiance822/time_memorizer/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.Server
}

var (
	grpcPort = ":8081"
	httpPort = ":8080"
)

func getDbConnection() (*sqlx.DB, error) {
	sqlxDB, err := sqlx.Connect("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}
	return sqlxDB, nil
}

func getGrpcServer() (*grpc.Server, error) {
	db, err := getDbConnection()
	if err != nil {
		return nil, err
	}
	impl := service.NewImplementation(db)
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(server)
	pb.RegisterTimeMemorizerServer(server, impl)
	return server, nil
}

func getHttpServer() (*http.Server, error) {
	server := grpc.NewServer()
	reflection.Register(server)
	list, err := net.Listen("tpc", grpcPort)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	gatewayConn, err := grpc.DialContext(
		ctx,
		list.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	grpcMux := http.NewServeMux()
	err = pb.RegisterTimeMemorizerHandler(ctx, grpcMux, gatewayConn)
	if err != nil {
		return nil, err
	}
	return &http.Server{
		Addr:    httpPort,
		Handler: grpcMux,
	}, nil
}
func NewApp() (*App, error) {
	httpSrv, err := getHttpServer()
	if err != nil {
		return nil, err
	}

	grpcServer, err := getGrpcServer()
	if err != nil {
		return nil, err
	}

	return &App{
		grpcServer: grpcServer,
		httpServer: httpSrv,
	}, nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return nil
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return nil
	}

	return nil
}

func (a *App) runHTTPServer() error {
	return a.httpServer.ListenAndServe()
}

func (a *App) Run() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.runHTTPServer()
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
