package configs

import "github.com/gin-gonic/gin"

func CustomGin() *gin.Engine {
	ginMode := "release"
	gin.SetMode((ginMode))
	r := gin.New()
	r.Use(gin.Recovery())
	return r
}
