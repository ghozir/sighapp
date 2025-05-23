package auth

import (
	authcontract "github.com/ghozir/sighapp/contracts"
	"github.com/ghozir/sighapp/database/mongodb"
	"github.com/ghozir/sighapp/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authRepository struct {
	userSvc    *mongodb.MongoService
	sessionSvc *mongodb.MongoService
}

func NewAuthRepository() authcontract.Repository {
	return &authRepository{
		userSvc:    mongodb.NewMongoService(mongodb.MongoDB.Collection("user")),
		sessionSvc: mongodb.NewMongoService(mongodb.MongoDB.Collection("session")),
	}
}

func (r *authRepository) FindOneUser(params bson.M) (*entities.User, error) {
	var user entities.User
	err := r.userSvc.FindOne(params, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) FindOneToken(params bson.M) (*entities.Session, error) {
	var session entities.Session
	err := r.sessionSvc.FindOne(params, nil, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *authRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.userSvc.FindOne(bson.M{"email": email}, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) InsertOneSession(data entities.Session) (*entities.Session, error) {
	res, err := r.sessionSvc.InsertOne(data)
	if err != nil {
		return nil, err
	}

	inserted := data
	inserted.ID = res.InsertedID.(primitive.ObjectID)
	return &inserted, nil
}
