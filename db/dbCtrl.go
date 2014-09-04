package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

const dbName = "app29209595"
const colName = "secmsg"

func getSesseion() *mgo.Session {
	uri := os.Getenv("MONGOHQ_URL")
	fmt.Println(uri)
	if uri == "" {
		fmt.Println("no connection string provided")
		return nil
	}
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		return nil
	}
	return sess
}

func SaveSceMessage(msg *SceMessage) bool {
	sess := getSesseion()
	if sess == nil {
		return false
	}
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB(dbName).C(colName)
	err := collection.Insert(*msg)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		return false
	}
	return true
}
