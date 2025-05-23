package authcontract

import (
	"github.com/ghozir/sighapp/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	FindUserByEmail(email string) (*entities.User, error)
	FindOneUser(params bson.M) (*entities.User, error)
	FindOneToken(params bson.M) (*entities.Session, error)
	InsertOneSession(data entities.Session) (*entities.Session, error)
}
