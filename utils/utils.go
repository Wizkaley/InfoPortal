package utils

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// DataStore Struct to handle mongo connectivity
type DataStore interface {
	Init() (db *mgo.Session)
}

// Init ...
func Init(path string) (sess *mgo.Session, err error) {
	sess, err = mgo.Dial(path)
	// if err != nil {
	// 	log.Printf("Error while getting session : %v", err)
	// 	return nil, err
	// }
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Error while getting session : no reachable servers")
		}
	}()
	// ses = ensureInd(sess)
	// if err != nil {
	// 	log.Printf("Error while Applying Indicex : %v", err)
	// }
	//indexes, _ := sess.DB("trial").C("Student").Indexes()

	// fmt.Println("")
	// for _, val := range indexes {
	// 	fmt.Println(val.Name)
	// }
	return
}

// func ensureInd(s *mgo.Session, db string, c string) (sesh *mgo.Session) {
// 	//ses.DB("trial").C("Student").EnsureIndex(workIndices[0].Index)
// 	s.DB(db).C(c).EnsureIndexKey("studentMarks")
// 	sesh = s
// 	return
// }
