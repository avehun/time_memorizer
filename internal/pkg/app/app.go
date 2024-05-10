package app

import (
	"github.com/radiance822/time_memorizer/internal/app/model"
	service "github.com/radiance822/time_memorizer/internal/app/time_memorizer"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app App) Run() {
	storage := model.CategoryStorage{}
	grpcAdress, httpAdress := ":8081", ":8080"
	service.StartServers(grpcAdress, httpAdress, &storage)
}
