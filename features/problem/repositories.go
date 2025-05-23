package problem

import (
	"github.com/ghozir/sighapp/database/mongodb"
	"github.com/ghozir/sighapp/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	InsertOneProblem(data entities.Problem) (*entities.Problem, error)
}

type problemRepository struct {
	problemSvc *mongodb.MongoService
}

func NewProblemRepository() Repository {
	return &problemRepository{
		problemSvc: mongodb.NewMongoService(mongodb.MongoDB.Collection("problem")),
	}
}

func (r *problemRepository) InsertOneProblem(data entities.Problem) (*entities.Problem, error) {
	res, err := r.problemSvc.InsertOne(data)
	if err != nil {
		return nil, err
	}

	inserted := data
	inserted.ID = res.InsertedID.(primitive.ObjectID)
	return &inserted, nil
}
