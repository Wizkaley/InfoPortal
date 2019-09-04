package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2/json"
	"RestApp/model"
	"log"
	"fmt"
)

var db *mgo.Database // Global Variable to Hold Database Connection

func Init(){ 
 	connect()
}

func connect(){
	session, err := mgo.Dial("localhost:27017") //session Dialed to mongo server
	if err == nil{
		db = session.DB("myDB")
		fmt.Println("Successfully connected")
	}
	
		
}
// Add Student 
func AddStudent(stud model.Student)error{
	return db.C("Student").Insert(stud)
}


//Remove student 
func RemoveByName(s string)error{
	stu,err := GetByName(s)
	if err!=nil{
		log.Printf("")
		return err
	}
	err = db.C("Student").Remove(stu)
	if err != nil {
		//log
		return err
	}
	return nil
}


//Get Student By Name
func GetByName(i string) (stu model.Student, err error){
	fmt.Println(i)
	err = db.C("Student").Find(bson.M{"studentName":i}).One(&stu)
	return 
}

//Get All the Students
func GetAll()(students [] model.Student, err error ) {
	err = db.C("Student").Find(bson.M{}).All(&students)
	return
}


//Update Student
func UpdateStudent(student model.Student)(err error){
	err = db.C("Student").Update(
			bson.M{"studentName":student.StudentName},
			student)
	return 
}

