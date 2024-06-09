package main

import (
	"log"

	"github.com/radiance822/time_memorizer/internal/pkg/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatal("failed initializing app: %v", err)
	}
	app.Run()
}
