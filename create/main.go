package main

import (
	"github.com/micro-service-create-carouselimage/configs"
	"github.com/micro-service-create-carouselimage/repositories"
	routes "github.com/micro-service-create-carouselimage/router"
)

var file, _ = configs.Init()

func main() {
	defer file.Close()

	// log.Info("Info message")
	// log.Warn("Warn message")
	// log.Error("Error message")
	// log.Fatal("Fatal message")

	router := configs.CustomGin()
	repositories.Opts.Ping()
	routes.ImgRoute(router)
	router.Run(":5000")
}
