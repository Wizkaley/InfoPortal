package dao

import(
	"testing"
	"RestApp/model"
	"fmt"
	"github.ibm.com/dash/dash_utils/dashtest"
)


func TestConnect(t *testing.T){
	connect()
}

func TestAddStudent(t *testing.T){
	Init()
	tst := model.Student{
		StudentName:"Pretty",
		StudentAge:56,
		StudentMarks:99,
	}

	err := AddStudent(tst); if err!=nil{
		fmt.Printf("Error Was Expected : %v",err)
	}

}

func TestRemoveStudent(t * testing.T){
	var name = "Light Yagami"

	err := RemoveByName(name); if err!= nil{
		t.Errorf("Error not Expected but : %v",err)
	}
}

func TestRemoveStudentErr(t * testing.T){
	var name ="sajdlas"
	err := RemoveByName(name); if err!= nil{
		fmt.Printf("Error Expected : %v",err)
	}
	
}
func TestGetByName(t *testing.T){
	var n = "Devansh"

	stud,err := GetByName(n); if err!=nil{
		t.Errorf("Error Not Expected but : %v",err)
	}
	fmt.Println(stud)
}

func TestGetByNameErr(t *testing.T){
	var n = "jdfhsdjfhks"

	stud,err := GetByName(n); if err!=nil{
		fmt.Printf("Error Expected : %v",err)
	}
	fmt.Println(stud)
}

func TestGetAll(t * testing.T){
	var s [] model.Student
	s,err := GetAll(); if err!=nil{
		t.Errorf("Error Not Expected but : %v",err)
	}
	fmt.Println(s)
}

func TestUpdateStudent(t *testing.T){
	tst := model.Student{
		StudentName:"Devansh",
		StudentAge:28,
		StudentMarks:99,
	} 

	err := UpdateStudent(tst); if err!=nil{
		t.Errorf("Error Not Expected but : %v",err)
	}
}

func TestUpdateStudentErr(t *testing.T){
	tst := model.Student{
		StudentName:"as",
		StudentAge:28,
		StudentMarks:99,
	} 

	err := UpdateStudent(tst); if err!=nil{
		fmt.Printf("Error Not Expected but : %v",err)
	}
}

func TestMain(m *testing.M){
	dashtest.ControlCoverage(m)
}