package core

import "time"

type Test struct {
	SourceId    string    `json:"source_id"`
	RoomId      int       `json:"room_id"`
	UserId      int       `json:"user_id"`
	EnteredTime time.Time `json:"entered_time"`
	QuitTime    time.Time `json:"quit_time"`
	QuitType    int       `json:"quit_type"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ix Test) IndexName() string {
	return "test"
}

/*
{
  "settings": {
    "index": {
      "lifecycle": {
        "name": "test-1",
        "rollover_alias": "test"
      },
      "refresh_interval": "1s",
      "number_of_shards": "6",
      "number_of_replicas": "2"
    }
  },
  "mappings": {
    "dynamic": "false",
    "dynamic_templates": [],
    "properties": {
      "created_at": {
        "type": "date"
      },
      "entered_time": {
        "type": "date"
      },
      "quit_time": {
        "type": "date"
      },
      "quit_type": {
        "type": "byte"
      },
      "room_id": {
        "type": "integer"
      },
      "source_id": {
        "type": "keyword"
      },
      "user_id": {
        "type": "integer"
      }
    }
  },
  "aliases": {}
}
*/
