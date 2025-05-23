package auth

import (
	"github.com/ghozir/sighapp/database/mongodb"
	"github.com/ghozir/sighapp/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	FindUserByEmail(email string) (*entities.User, error)
	FindOneUser(filter bson.M) (*entities.User, error)
	FindOneToken(filter bson.M) (*entities.Session, error)
	InsertOneSession(data entities.Session) (*entities.Session, error)
}

type authRepository struct {
	userSvc    *mongodb.MongoService
	sessionSvc *mongodb.MongoService
}

func NewAuthRepository() Repository {
	return &authRepository{
		userSvc:    mongodb.NewMongoService(mongodb.MongoDB.Collection("user")),
		sessionSvc: mongodb.NewMongoService(mongodb.MongoDB.Collection("session")),
	}
}

func (r *authRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.userSvc.FindOne(bson.M{"email": email}, nil, &user)
	return &user, err
}

func (r *authRepository) FindOneUser(filter bson.M) (*entities.User, error) {
	var user entities.User
	err := r.userSvc.FindOne(filter, nil, &user)
	return &user, err
}

func (r *authRepository) FindOneToken(filter bson.M) (*entities.Session, error) {
	var session entities.Session
	err := r.sessionSvc.FindOne(filter, nil, &session)
	return &session, err
}

func (r *authRepository) InsertOneSession(data entities.Session) (*entities.Session, error) {
	res, err := r.sessionSvc.InsertOne(data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return &data, nil
}
