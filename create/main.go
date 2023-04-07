package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro-service-create-carouselimage/repositories"
	routes "github.com/micro-service-create-carouselimage/router"
)

func main() {
	router := gin.Default()
	repositories.Opts.Ping()
	routes.ImgRoute(router)
	router.Run(":5000")
}
