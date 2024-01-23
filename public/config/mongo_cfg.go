package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type MongoCfg struct {
	Name        string `json:"Name"`
	Uri         string `json:"Uri"`
	Db          string `json:"Db"`
	MinPoolSize uint64 `json:"MinPoolSize,default=5"`
	MaxPoolSize uint64 `json:"MaxPoolSize,default=50"`
}

// ClientMongoConfig 客户端mysql配置
type ClientMongoConfig struct {
	MongoCfg
	Database *mongo.Database
	RwMutex  sync.RWMutex
}
