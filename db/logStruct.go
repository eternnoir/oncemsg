package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type LogMsg struct {
	Id      bson.ObjectId `bson:"_id"`
	LogTime time.Time
	Msg     string
}
