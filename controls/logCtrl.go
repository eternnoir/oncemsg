package controls

import (
	"oncemsg/db"
	"time"
)


func LogInfo(MsgContent string) bool {
	log := db.CreateLogMsg("INFO", MsgContent, time.Now())
	return db.SaveLogMsg(log)
}

func LogWarn(MsgContent string) bool {
	log := db.CreateLogMsg("WARN", MsgContent, time.Now())
	return db.SaveLogMsg(log)
}

func LogError(err error) bool {
	log := db.CreateLogMsg("ERR", err.Error(), time.Now())
	return db.SaveLogMsg(log)
}
