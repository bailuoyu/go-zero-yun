package model

import (
	"go-zero-yun/public/logic/client/monlgc"
	"go.mongodb.org/mongo-driver/mongo"
)

func MonDb() *mongo.Database {
	return monlgc.Core()
}
