package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro-service-create-carouselimage/entities"
	"github.com/micro-service-create-carouselimage/responses"
)

func (a *HandlerAdd) AddOne(c *gin.Context) {
	var img *entities.Image
	if err := c.BindJSON(&img); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: responses.MsgHandler.FailureBindJSON})
		return
	}

	company := c.Param("company")

	path := c.Request.URL.Path
	s := strings.Split(path, "/")
	response, err := a.Repo.AddOne(company, s[3], img)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: response.Message})
		return
	}

	c.JSON(http.StatusOK, &responses.HandleResponse{Message: response.Message, Data: response.Data})
}
