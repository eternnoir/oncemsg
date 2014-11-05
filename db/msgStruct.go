package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SceMessage struct {
	Id          bson.ObjectId `bson:"_id"`
	UnId        string
	Content     string `bson:"msg"`
	Type        string
	ExpiredDate time.Time
	CreateDate  time.Time
}

func CreateSceMessage(Unid, Content, Type string) *SceMessage {
	ret := &SceMessage{}
	ret.UnId = Unid
	ret.Content = Content
	ret.Type = Type
	ret.Id = bson.NewObjectId()
	ret.CreateDate = time.Now()
	return ret
}
