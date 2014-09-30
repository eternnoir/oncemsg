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

func CreateLogMsg(Msg string, logTime time.Time) *LogMsg {
	ret := &LogMsg{}
	ret.LogTime = logTime
	ret.Msg = Msg
	return ret
}
