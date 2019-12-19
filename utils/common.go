package utils

import (
	mgo "gopkg.in/mgo.v2"
	"strconv"
)

// MgoDial ...
var MgoDial = mgo.DialWithInfo

// GetDataBaseSession returns a session from the DB
func GetDataBaseSession() (session *mgo.Session, err error) {
	Host := []string{
		Config.DatabaseHost + ":" + strconv.Itoa(Config.DatabasePort),
	}
	session, err = MgoDial(&mgo.DialInfo{
		Addrs: Host,
	})
	return
}

// GetDataBaseSessionWithURI returns a session from the DB
func GetDataBaseSessionWithURI(uri string) (session *mgo.Session, err error) {
	//var dialInfo *mgo.DialInfo

	session, err = mgo.Dial(uri)
	if err != nil {
		//log.Print("Error connecting to Database")
		//return nil, errors.Wrap(err, "Couldn't connect to DB")
		panic("Couldn't connect to DB")
	}
	return session, nil
}
