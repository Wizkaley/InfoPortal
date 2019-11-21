package utils

import (
	"gopkg.in/mgo.v2"
)

// MgoDial ...
var MgoDial = mgo.Dial

// GetDataBaseSession returns a session from the DB
func GetDataBaseSession(uri string) (session *mgo.Session, err error) {
	//var dialInfo *mgo.DialInfo

	session, err = mgo.Dial(uri)
	if err != nil {
		//log.Print("Error connecting to Database")
		//return nil, errors.Wrap(err, "Couldn't connect to DB")
		panic("Couldn't connect to DB")
	}
	return session, nil
}
