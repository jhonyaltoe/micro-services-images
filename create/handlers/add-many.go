package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro-service-create-carouselimage/entities"
	"github.com/micro-service-create-carouselimage/responses"
)

func (a *HandlerAdd) AddMany(c *gin.Context) {
	var imgs []*entities.Image
	if err := c.BindJSON(&imgs); err != nil {
		Log.Errorf("Error decoding: %v", err)
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: responses.MsgHandler.FailureBindJSON})
		return
	}

	for i := range imgs {
		if errs := imgs[i].Validate(); errs != nil {
			Log.Errorf("Validation error: %v", errs)
			c.JSON(http.StatusBadRequest, &responses.HandleResponse{
				Message: responses.MsgHandler.ValidationErr,
				Error:   errs,
				Index:   strconv.Itoa(i),
			})
			return
		}
	}

	company := c.Param("company")

	path := c.Request.URL.Path
	s := strings.Split(path, "/")
	response, err := a.Repo.AddMany(company, s[3], imgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responses.HandleResponse{Message: response.Message})
		return
	}

	c.JSON(http.StatusOK, &responses.HandleResponse{Message: response.Message, Data: response.Data})
}
