package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/micro-service-create-carouselimage/configs"
	"github.com/micro-service-create-carouselimage/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoAttrs struct {
	Client *mongo.Client
}

type IPing interface {
	Ping()
}

type PingAttr struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func new() (handlers.IRepository, IPing) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Client().ApplyURI(configs.EnvMongoURI())
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	return &RepoAttrs{
			Client: client,
		}, &PingAttr{
			Client: client,
			Ctx:    ctx,
			Cancel: cancel,
		}
}

func (r *PingAttr) Ping() {
	if err := r.Client.Ping(r.Ctx, nil); err != nil {
		log.Fatal(err)
		r.Cancel()
	}
	fmt.Println("Connected to MongoDB")
}

var Repo, Opts = new()
