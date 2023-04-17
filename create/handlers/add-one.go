package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro-service-create-carouselimage/entities"
	"github.com/micro-service-create-carouselimage/responses"
)

func (a *HandlerAdd) AddOne(c *gin.Context) {
	var img *entities.Image
	if err := c.BindJSON(&img); err != nil {
		Log.Errorf("Error decoding: %v", err)
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: responses.MsgHandler.FailureBindJSON})
		return
	}

	if errs := img.Validate(); errs != nil {
		Log.Errorf("Validation error: %v", errs)
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: responses.MsgHandler.ValidationErr, Error: errs})
		return
	}

	company := c.Param("company")

	path := c.Request.URL.Path
	s := strings.Split(path, "/")
	response, err := a.Repo.AddOne(company, s[3], img)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: response.Message})
		return
	}

	c.JSON(http.StatusOK, &responses.HandleResponse{Message: response.Message, Data: response.Data})
}
