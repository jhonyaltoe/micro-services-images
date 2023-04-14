package repositories

import (
	"context"
	"time"

	"github.com/micro-service-create-carouselimage/entities"
	"github.com/micro-service-create-carouselimage/responses"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepoAttrs) AddMany(company string, documentName string, imgs []*entities.Image) (*responses.DbResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := r.Client.Database(company).Collection("test")
	filter := bson.M{"name": documentName}

	update := bson.M{
		"$push": bson.M{"images": bson.M{"$each": imgs}},
	}

	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		Log.Errorf("Error insert to MongoDB: %v", err)
		return &responses.DbResponse{Message: responses.MsgDb.InternalError}, err
	}

	if result.MatchedCount == 1 {
		return &responses.DbResponse{Message: responses.MsgDb.SuccessUpdate}, nil
	}

	return &responses.DbResponse{Message: responses.MsgDb.FailureUpdate}, nil
}
