package test

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	model "go-zero-yun/public/model/mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMongoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MongoLogic {
	return &MongoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MongoLogic) Mongo(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	Model := model.NewDemoModel()
	var ds []model.Demo
	//err = Model.Conn().Find(l.ctx, &ds, bson.M{"age": bson.M{"$gt": 20}})
	opt := mopt.Find().SetSkip(1)
	err = Model.Conn().Find(l.ctx, &ds, bson.M{}, opt)
	if err != nil {
		return
	}
	fmt.Println(ds)
	return
}
