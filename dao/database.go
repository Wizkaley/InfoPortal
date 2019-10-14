package dao

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// DataStore Struct to handle mongo connectivity
type DataStore struct {
	db *mgo.Session
}

// Init ...
func Init(path string) (ses *mgo.Session, err error) {
	sess, err := mgo.Dial(path)
	if err != nil {
		log.Printf("Error while getting session : %v", err)
	}
	ses = ensureInd(sess)
	if err != nil {
		log.Printf("Error while Applying Indicex : %v", err)
	}
	return
}

func ensureInd(s *mgo.Session) (sesh *mgo.Session) {
	//ses.DB("trial").C("Student").EnsureIndex(workIndices[0].Index)
	s.DB("trial").C("Student").EnsureIndexKey("studentMarks")
	indexes, _ := s.DB("trial").C("Student").Indexes()

	for _, val := range indexes {
		fmt.Println(val)
	}
	return
}
