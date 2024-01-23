// Package ckmonkit https://cloud.tencent.com/document/product/597/83341#.E6.97.A0.E6.95.88.E4.BA.8B.E4.BB.B6.EF.BC.88invalidate-event.EF.BC.89
package ckmonkit

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	OptInsert       = "insert"
	OptUpdate       = "update"
	OptDelete       = "delete"
	OptReplace      = "replace"
	OptDrop         = "drop"
	OptRename       = "rename"
	OptDropDatabase = "dropDatabase"
	OptInvalidate   = "invalidate"
)

// EventKey 事件Key
type EventKey struct {
	Schema  EventSchema `json:"schema"`
	Payload string      `json:"payload"`
}

type EventValue struct {
	Schema  EventSchema `json:"schema"`
	Payload string      `json:"payload"`
}

// EventSchema 事件提要
type EventSchema struct {
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
}

// EVPayload 事件value
type EVPayload struct {
	ID struct {
		Data string `bson:"_data" json:"_data"`
	} `bson:"_id" json:"_id"`
	OperationType string   `bson:"operationType" json:"operationType"`
	FullDocument  bson.Raw `bson:"fullDocument" json:"fullDocument"`
	Ns            struct {
		Db   string `bson:"db" json:"db"`
		Coll string `bson:"coll" json:"coll"`
	} `bson:"ns" json:"ns"`
	To struct {
		Db   string `bson:"db" json:"db"`
		Coll string `bson:"coll" json:"coll"`
	} `bson:"to" json:"to"`
	DocumentKey struct {
		ID primitive.ObjectID `bson:"_id" json:"_id"`
	} `bson:"documentKey" json:"documentKey"`
	UpdateDescription struct {
		UpdatedFields   bson.Raw `bson:"updatedFields" json:"updatedFields"`
		RemovedFields   []string `bson:"removedFields" json:"removedFields"`
		TruncatedArrays []struct {
			Field   string `bson:"field" json:"field"`
			NewSize int    `bson:"newSize" json:"newSize"`
		} `bson:"truncatedArrays" json:"truncatedArrays"`
	} `bson:"updateDescription" json:"updateDescription"`
	ClusterTime primitive.Timestamp `bson:"clusterTime" json:"clusterTime"`
	TxnNumber   int64               `bson:"txnNumber" json:"txnNumber"`
	Lsid        struct {
		Id  string        `bson:"id" json:"id"`
		Uid bson.RawValue `bson:"uid" json:"uid"`
	} `bson:"lsid" json:"lsid"`
}

// GetVersionNum 获取唯一版本号，用于保证顺序一致性
func (payload EVPayload) GetVersionNum() uint64 {
	ver := uint64(payload.ClusterTime.T)*1e8 + uint64(payload.ClusterTime.I)
	return ver
}

// GetEventKey 获取事件key
func GetEventKey(kb []byte) (EventKey, error) {
	var evk EventKey
	err := json.Unmarshal(kb, &evk)
	return evk, err
}

// GetEventValue 获取事件value
func GetEventValue(vb []byte) (EventValue, error) {
	var evv EventValue
	err := json.Unmarshal(vb, &evv)
	return evv, err
}

// GetEVPayload 解析内容payload
func GetEVPayload(vb []byte) (EVPayload, error) {
	var payload EVPayload
	evv, err := GetEventValue(vb)
	if err != nil {
		return payload, err
	}
	err = bson.UnmarshalExtJSON([]byte(evv.Payload), false, &payload)
	return payload, err
}
