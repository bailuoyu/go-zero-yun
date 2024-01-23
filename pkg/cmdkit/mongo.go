package cmdkit

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/syncx"
	mon "go-zero-yun/pkg/monkit"
	pconf "go-zero-yun/public/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientManager = syncx.NewResourceManager()

// initMongo 初始化mongo
func initMongo(check bool) {
	pconf.ClientCfg.Mongo = make(map[string]*pconf.ClientMongoConfig)
	for _, v := range pconf.Cfg.Client.Mongo {
		pconf.ClientCfg.Mongo[v.Name] = &pconf.ClientMongoConfig{
			MongoCfg: v,
		}
		if check && !IsEnvLocal() { //检查连接
			_, err := MongoConnect(pconf.ClientCfg.Mongo[v.Name])
			if err != nil {
				panic(errors.New(fmt.Sprintf("mongo fatal ping, name: %s, error: %s", v.Name, err.Error())))
			}
		}
	}
}

// closeMongo 关闭Mongo
func closeMongo() {
	for _, v := range pconf.ClientCfg.Mongo {
		if v != nil && v.Database != nil {
			v.Database.Client().Disconnect(context.Background())
		}
	}
}

// MongoConnect 检查连接是否可用
func MongoConnect(v *pconf.ClientMongoConfig) (*mongo.Database, error) {
	client, err := mon.GetClient(v.Uri,
		func(opts *mon.Options) {
			opts.SetMinPoolSize(v.MinPoolSize)
			opts.SetMaxPoolSize(v.MaxPoolSize)
			opts.SetReadPreference(readpref.SecondaryPreferred())
		},
	)
	if err != nil {
		return nil, err
	}
	db := client.Database(v.Db)
	v.Database = db
	db.ReadPreference()
	return db, nil
}
