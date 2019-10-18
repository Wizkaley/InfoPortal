package dao

import (
	"RESTApp/model"
	mongo "RESTApp/utils/mongo"
	"fmt"
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestAddStudent(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	tst := model.Student{
		StudentName:  "Pretty",
		StudentAge:   56,
		StudentMarks: 99,
	}

	err = AddStudent(tst, ds)
	if err != nil {
		fmt.Printf("Error Was Expected : %v", err)
	}

}

func TestRemoveStudent(t *testing.T) {
	var name = "Pretty"
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	err = RemoveByName(name, ds)
	if err != nil {
		t.Errorf("Error not Expected but : %v", err)
	}
}

func TestRemoveStudentErr(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var name = "sajdlas"
	err = RemoveByName(name, ds)
	if err != nil {
		fmt.Printf("Error Expected : %v", err)
	}

}
func TestGetByName(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var n = "ASAP"

	stud, err := GetByName(n, ds)
	if err != nil {
		t.Errorf("Error Not Expected but : %v", err)
	}
	fmt.Println(stud)
}

func TestGetByNameErr(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var n = "jdfhsdjfhks"

	stud, err := GetByName(n, ds)
	if err != nil {
		fmt.Printf("Error Expected : %v", err)
	}
	fmt.Println(stud)
}

func TestGetAll(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var s []model.Student
	s, err = GetAll(ds)
	if err != nil {
		t.Errorf("Error Not Expected but : %v", err)
	}
	fmt.Println(s)
}

func TestUpdateStudent(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	tst := model.Student{
		StudentName:  "ASAP",
		StudentAge:   28,
		StudentMarks: 99,
	}

	err = UpdateStudent(tst, ds)
	if err != nil {
		t.Errorf("Error Not Expected but : %v", err)
	}
}

func TestUpdateStudentErr(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	tst := model.Student{
		StudentName:  "as",
		StudentAge:   28,
		StudentMarks: 99,
	}

	err = UpdateStudent(tst, ds)
	if err != nil {
		fmt.Printf("Error Not Expected but : %v", err)
	}
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
