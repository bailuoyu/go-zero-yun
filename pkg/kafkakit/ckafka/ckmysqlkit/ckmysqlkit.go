// Package ckmysqlkit https://cloud.tencent.com/document/product/597/83342#.E5.88.A0.E9.99.A4.E4.BA.8B.E4.BB.B6.EF.BC.88delete-events.EF.BC.89
package ckmysqlkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	OptInsert = "c"
	OptUpdate = "u"
	OptDelete = "d"
)

// EventKey 事件Key
type EventKey struct {
	Schema  EventSchema     `json:"schema"`
	Payload json.RawMessage `json:"payload"`
}

type EventValue struct {
	Schema  EventSchema     `json:"schema"`
	Payload json.RawMessage `json:"payload"`
}

// EventSchema 事件提要
type EventSchema struct {
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
	Name     string `json:"name"`
	Fields   []struct {
		Field    string `json:"field"`
		Type     string `json:"type"`
		Optional bool   `json:"optional"`
	} `json:"fields"`
}

// EVPayload 事件payload
type EVPayload struct {
	Op     string          `json:"op"`
	TsMs   int64           `json:"ts_ms"`
	Before json.RawMessage `json:"before"`
	After  json.RawMessage `json:"after"`
	Source struct {
		Version   string `json:"version"`
		Connector string `json:"connector"`
		Name      string `json:"name"`
		TsMs      int64  `json:"ts_ms"`
		//Snapshot  string `json:"snapshot"`	//文档显示bool实际返回string,注释不用
		Db       string `json:"db"`
		Table    string `json:"table"`
		ServerId int64  `json:"server_id"`
		Gtid     string `json:"gtid"`
		File     string `json:"file"`
		Pos      int64  `json:"pos"`
		Row      int    `json:"row"`
		Thread   int64  `json:"thread"`
		Query    string `json:"query"`
	} `json:"source"`
}

// GetVersionNum 获取唯一版本号，用于保证顺序一致性
func (payload EVPayload) GetVersionNum() (uint64, error) {
	if payload.Source.File == "" {
		return 0, errors.New("empty source file")
	}
	inx := strings.Index(payload.Source.File, ".")
	if inx <= 0 {
		return 0, fmt.Errorf("error source file:%s", payload.Source.File)
	}
	logNumStr := payload.Source.File[inx+1:]
	logNum, err := strconv.Atoi(logNumStr)
	if err != nil {
		return 0, err
	}
	ver := uint64(logNum)*1e10 + uint64(payload.Source.Pos)
	return ver, nil
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
	err = json.Unmarshal(evv.Payload, &payload)
	return payload, err
}
