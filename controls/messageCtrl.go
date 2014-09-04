package controls

import (
	"math/rand"
	uti "oncemsg/utility"
	"time"
)

const layout = "2006-01-02T15:04:05.999999-07:00"

func SaveMsg(msg string) (string, error) {
	rand.Seed(42)
	ha1 := uti.GetUniHashString(msg)
	timeStr := time.Now().Format(layout)
	ranNum := string(rand.Int31())
	uniString := uti.GetUniHashString(ha1 + timeStr + ranNum)

	return uniString, nil
}
