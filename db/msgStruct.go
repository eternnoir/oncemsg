package db

import (
	"gopkg.in/mgo.v2/bson"
	"oncemsg/utility"
	"time"
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
	ret.Content, _ = utility.CryAse(Content)
	ret.Type = Type
	ret.Id = bson.NewObjectId()
	ret.CreateDate = time.Now()
	return ret
}
