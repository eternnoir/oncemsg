package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "app29209595"
const colName = "secmsg"

func getSesseion() *mgo.Session {
	uri := os.Getenv("MONGOLAB_URI")
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
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB(dbName).C(colName)
	err := collection.Insert(*msg)
	sess.Close()
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		return false
	}
	return true
}

func DeleteSceMessage(uniid string) bool {
	sess := getSesseion()
	if sess == nil {
		return false
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB(dbName).C(colName)
	err := collection.Remove(bson.M{"unid": uniid})
	sess.Close()
	if err != nil {
		fmt.Printf("Can't remove document: %v\n", err)
		return false
	}
	return true
}

func GetSceMessage(unid string) *SceMessage {
	sess := getSesseion()
	if sess == nil {
		return nil
	}
	defer sess.Close()
	var ret SceMessage
	collection := sess.DB(dbName).C(colName)
	err := collection.Find(bson.M{"unid": unid}).One(&ret)
	if err != nil {
		fmt.Printf("got an error finding a doc %v\n")
		return nil
	}
	return &ret
}
