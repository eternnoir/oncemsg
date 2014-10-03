package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type LogMsg struct {
	Id      bson.ObjectId `bson:"_id"`
	LogTime time.Time
	Msg     string
	LogTyep string
}

func CreateLogMsg(LogType string,Msg string, logTime time.Time) *LogMsg {
	ret := &LogMsg{}
	ret.Id = bson.NewObjectId()
	ret.LogTime = logTime
	ret.LogTyep = LogType
	ret.Msg = Msg
	return ret
}
