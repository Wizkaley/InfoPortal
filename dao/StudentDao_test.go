package dao

import (
	mocks "RESTApp/mocks"
	"RESTApp/model"
	mongo "RESTApp/utils/mongo"
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/golang/mock/gomock"
)

func TestAddStudent(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	tst := model.Student{
		StudentName:  "Pretty",
		StudentAge:   56,
		StudentMarks: 99,
	}

	err = AddStudent(tst, ds, testingdb)
	if err != nil {
		fmt.Printf("Error Was Expected : %v", err)
	}

}

func TestAddStudentErr(t *testing.T) {
	ds, _ := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()

	testStu := model.Student{
		StudentName:  "test",
		StudentAge:   24,
		StudentMarks: 24,
	}

	err := AddStudent(testStu, ds, testingdb)
	if err != nil {
		t.Errorf("Failed: %v", err)
	}
	//NewMongoDAL = mockMongo
}

func TestRemoveStudent(t *testing.T) {
	var name = "Pretty"
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	err = RemoveByName(name, ds, testingdb)
	if err != nil {
		t.Errorf("Error not Expected but : %v", err)
	}
}

func TestRemoveStudentErr(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var name = "sajdlas"
	err = RemoveByName(name, ds, testingdb)
	if err != nil {
		fmt.Printf("Error Expected : %v", err)
	}

}
func TestGetByName(t *testing.T) {
	ds, _ := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var n = "test"

	stud, err := GetByName(n, ds, "testing")
	if err != nil {
		t.Errorf("Error Not Expected but : %v", err)
	}
	fmt.Println(stud)
	//------------------------------------------
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// db := mocks.NewMockMgoDBDAL(mockCtrl)
	// coll := mocks.NewMockMgoCollectionDAL(mockCtrl)

	// db.EXPECT().C("testing").Return(coll).Times(1)

	// coll.EXPECT().Find("ASAP").Return(nil).Times(1)

	// _, _ = GetByName("ASAP", ds, testingdb)
}

func TestGetByNameErr(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var n = "jdfhsdjfhks"

	stud, err := GetByName(n, ds, testingdb)
	if err != nil {
		fmt.Printf("Error Expected : %v", err)
	}
	fmt.Println(stud)
}

func TestGetAll(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	var s []model.Student
	s, err = GetAll(ds, testingdb)
	if err != nil {
		t.Errorf("Error Not Expected but : %v", err)
	}
	fmt.Println(s)
}

func TestUpdateStudent(t *testing.T) {
	ds, err := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	tst := model.Student{
		StudentName:  "test",
		StudentAge:   28,
		StudentMarks: 99,
	}

	err = UpdateStudent(tst, ds, testingdb)
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

	err = UpdateStudent(tst, ds, testingdb)
	if err != nil {
		fmt.Printf("Error Not Expected but : %v", err)
	}
}

func TestRemoveByNameErr(t *testing.T) {
	ds, _ := mongo.GetDataBaseSession("localhost:27017")
	defer ds.Close()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	coll := mocks.NewMockMgoCollectionDAL(mockCtrl)
	coll.EXPECT().Remove(gomock.Any()).Return(mgo.ErrNotFound).Times(1)
	coll.Remove(gomock.Any())
}
