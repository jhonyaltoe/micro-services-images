package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/micro-service-create-carouselimage/handlers"
	"github.com/micro-service-create-carouselimage/repositories"
)

var handler = &handlers.HandlerAdd{
	Repo: repositories.Repo,
}

func ImgRoute(router *gin.Engine) {
	router.POST("/assets/:company/ourWorks", handler.AddMany)
	router.POST("/assets/:company/logo", handler.AddOne)
	router.POST("/assets/:company/products", handler.AddMany)
	router.POST("/assets/:company/video", handler.AddMany)
}
