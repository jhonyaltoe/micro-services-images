package main

import (
	"github.com/micro-service-create-carouselimage/configs"
	"github.com/micro-service-create-carouselimage/repositories"
	routes "github.com/micro-service-create-carouselimage/router"
)

var file, log = configs.Init()

func main() {
	log.Warnf("Service Starting...")
	defer file.Close()
	router := configs.CustomGin()
	repositories.Opts.Ping()
	routes.ImgRoute(router)
	log.Info("Service Ready")
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
