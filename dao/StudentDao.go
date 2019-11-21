
package dao

import (
	"RESTApp/model"
	"RESTApp/utils/mongodal"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//  

// NewMongoDAL ...
var NewMongoDAL = mongodal.NewMongoDBDAL

// var (
// 	workIndices = []testIndices{
// 		{
// 			Name: "Marks",
// 			Index: mgo.Index{
// 				Key:        []string{"studentMarks"},
// 				Unique:     false,
// 				Background: true,
// 				DropDups:   false,
// 				Sparse:     false,
// 			},
// 		},
// 	}
// )

// AddStudent ...
func AddStudent(stud model.Student, ds *mgo.Session, db string) error {
	s := ds.Clone()
	defer s.Close()

	datab := s.DB(db)
	dal := NewMongoDAL(datab)
	err := dal.C("Student").Insert(stud)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//RemoveByName ...
func RemoveByName(s string, ds *mgo.Session, db string) error {

	se := ds.Clone()
	defer se.Close()
	err := se.DB(db).C("Student").Remove(bson.M{"studentName": s})
	if err != nil || err == mgo.ErrNotFound {
		//log
		return err
	}
	return nil
}

//GetByName ...
func GetByName(i string, ds *mgo.Session, db string) (stu model.Student, err error) {
	fmt.Println(i)
	s := ds.Clone()
	defer s.Close()
	err = NewMongoDAL(s.DB(db)).C("Student").Find(bson.M{"studentName": i}).One(&stu)
	//err = db.DB("trial").C("Student").Find(bson.M{"studentName": i}).One(&stu)
	return stu, err
}

//GetAll ...
func GetAll(ds *mgo.Session, db string) (students []model.Student, err error) {
	s := ds.Clone()
	defer s.Close()
	err = s.DB(db).C("Student").Find(bson.M{}).All(&students)
	return
}

//UpdateStudent ...
func UpdateStudent(student model.Student, ds *mgo.Session, db string) (err error) {
	s := ds.Clone()
	defer s.Close()
	err = s.DB(db).C("Student").Update(
		bson.M{"studentName": student.StudentName},
		student)
	return
}

// StudentAggregates returns Aggregates based on Marks
func StudentAggregates(marks int32, ds *mgo.Session) (err error) {
	db := ds.Clone()
	defer db.Close()
	//c := db.DB("trial").C("Student")
	//c.EnsureIndex(workIndices[0].Index)
	//pipe := c.Pipe(bson.M{"$count":marks})
	return
}
