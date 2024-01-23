package model

import (
	"context"
	"errors"
	mon "go-zero-yun/pkg/monkit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

const AccessTokenCollectionName = "access_token"

var _ AccessTokenModel = (*customAccessTokenModel)(nil)

type (
	// AccessTokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccessTokenModel.
	AccessTokenModel interface {
		authModel
		TokenLimit(ctx context.Context, userId int, max int64) error
		CheckToken(ctx context.Context, userId int, token string) (AccessToken, error)
	}

	customAccessTokenModel struct {
		*defaultAccessTokenModel
	}
)

// NewAccessTokenModel returns a model for the mongo.
func NewAccessTokenModel() AccessTokenModel {
	db := MonDb()
	conn := mon.MustNewModel(db, AccessTokenCollectionName)
	return &customAccessTokenModel{
		defaultAccessTokenModel: newDefaultAccessTokenModel(conn),
	}
}

// TokenLimit 自定义函数
func (m *customAccessTokenModel) TokenLimit(ctx context.Context, userId int, max int64) error {
	if max < 1 {
		max = 3
	}
	// 最多保留10个token
	var auth AccessToken
	opts := mopt.FindOne().SetSkip(max).SetSort(bson.M{"_id": -1})
	if err := m.conn.FindOne(ctx, &auth, bson.M{"user_id": userId}, opts); err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	} else {
		m.conn.DeleteMany(ctx, bson.D{{"user_id", userId}, {"_id", bson.M{"$lte": auth.ID}}})
	}
	return nil
}

// CheckToken 自定义函数
func (m *customAccessTokenModel) CheckToken(ctx context.Context, userId int, token string) (AccessToken, error) {
	var act AccessToken
	err := m.conn.FindOne(ctx, &act, bson.D{{"user_id", userId}, {"token", token}})
	if err != nil && err == mongo.ErrNoDocuments {
		return act, errors.New("不存在的token")
	}
	return act, err
}
