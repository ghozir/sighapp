package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password,omitempty"`
}

type SafeUser struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Email string             `bson:"email" json:"email"`
}

func (u *User) ToSafeUser() SafeUser {
	return SafeUser{
		ID:    u.ID,
		Email: u.Email,
	}
}
