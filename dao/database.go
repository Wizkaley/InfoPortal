package dao

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DataStore Struct to handle mongo connectivity
type DataStore struct {
	db *mgo.Session
}

// Init ...
func Init(path string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(path)
	if err != nil {
		log.Printf("Error while getting session : %v", err)
	}
	return
}
