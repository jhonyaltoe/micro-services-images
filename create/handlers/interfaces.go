package handlers

import (
	"github.com/micro-service-create-carouselimage/entities"
	"github.com/micro-service-create-carouselimage/responses"
)

type IRepository interface {
	AddMany(company string, documentName string, img []*entities.Image) (*responses.DbResponse, error)
	AddOne(company string, documentName string, img *entities.Image) (*responses.DbResponse, error)
}

type HandlerAdd struct {
	Repo IRepository
}
