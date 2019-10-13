package dao

import (
	"RESTApp/model"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Init ...
// func Init() {
// 	connect()
// }

// AddStudent ...
func AddStudent(stud model.Student, ds *mgo.Session) error {
	db := ds.Clone()
	defer db.Close()
	return db.DB("trial").C("Student").Insert(stud)
}

//RemoveByName ...
func RemoveByName(s string, ds *mgo.Session) error {

	db := ds.Clone()
	defer db.Close()
	stu, err := GetByName(s, ds)
	if err != nil {
		log.Printf("")
		return err
	}
	err = db.DB("trial").C("Student").Remove(stu)
	if err != nil {
		//log
		return err
	}
	return nil
}

//GetByName ...
func GetByName(i string, ds *mgo.Session) (stu model.Student, err error) {
	fmt.Println(i)
	db := ds.Clone()
	defer db.Close()
	err = db.DB("trial").C("Student").Find(bson.M{"studentName": i}).One(&stu)
	return
}

//GetAll ...
func GetAll(ds *mgo.Session) (students []model.Student, err error) {
	db := ds.Clone()
	defer db.Close()
	err = db.DB("trial").C("Student").Find(bson.M{}).All(&students)
	return
}

//UpdateStudent ...
func UpdateStudent(student model.Student, ds *mgo.Session) (err error) {
	db := ds.Clone()
	err = db.DB("trial").C("Student").Update(
		bson.M{"studentName": student.StudentName},
		student)
	return
}
