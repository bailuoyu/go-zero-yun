package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessToken struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	UserId   int       `bson:"user_id" json:"user_id"`
	Token    string    `bson:"token" json:"token"`
	ExpireAt time.Time `bson:"expire_at" json:"expire_at"`

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
