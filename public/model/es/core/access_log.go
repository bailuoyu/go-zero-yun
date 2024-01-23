package core

import (
	"time"
)

type AccessLog struct {
	Time    time.Time `json:"@timestamp"` //date
	Trace   string    `json:"trace"`      //keyword
	Span    string    `json:"span"`       //span
	Route   string    `json:"route"`      //keyword
	UserId  int       `json:"user_id"`    //integer
	Level   string    `json:"level"`      //keyword
	Type    string    `json:"type"`       //keyword
	Runtime float64   `json:"runtime"`    //float
	Content string    `json:"content"`    //text
	Caller  string    `json:"caller"`     //text
}

func (ix AccessLog) IndexName() string {
	return "access_log"
}

/**
建表语句
PUT /access_log
{
  "settings": {
    "refresh_interval": "5s",
    "number_of_shards": 6,
	"number_of_replicas": 2
  },
  "mappings": {
	"dynamic": false,
    "properties": {
      "@timestamp": {
        "type": "date"
      },
      "trace": {
        "type": "keyword"
      },
      "span": {
        "type": "keyword"
      },
      "route": {
        "type": "keyword"
      },
      "user_id": {
        "type": "integer"
      },
      "level": {
        "type": "keyword"
      },
      "type": {
        "type": "keyword"
      },
      "runtime": {
        "type": "float"
      },
      "content": {
        "type": "text"
      },
      "caller": {
        "type": "text","index": false
      }
    }
  }
}
*/
