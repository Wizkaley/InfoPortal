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
