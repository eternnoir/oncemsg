package controls

import (
	"errors"
	"math/rand"
	"oncemsg/db"
	uti "oncemsg/utility"
	"time"
)

const layout = "2006-01-02T15:04:05.999999-07:00"

func SaveMsg(msg, Type string) (string, error) {
	uniString := getUniString(msg)
	if saveMsgToDb(uniString, msg, Type) == false {
		return "", errors.New("Insert To Db Error")
	}
	return uniString, nil
}

func saveMsgToDb(unid, content, Type string) bool {
	scemsg := db.CreateSceMessage(unid, content, Type)
	return db.SaveSceMessage(scemsg)
}

func getUniString(msg string) string {
	rand.Seed(42)
	ha1 := uti.GetUniHashString(msg)
	timeStr := time.Now().Format(layout)
	ranNum := string(rand.Int31())
	uniString := uti.GetUniHashString(ha1 + timeStr + ranNum)
	return uniString
}
