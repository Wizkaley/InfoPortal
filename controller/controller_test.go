package controller

import (
	"RESTApp/dao"
	"RESTApp/model"
	"RESTApp/utils"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	//utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

}

func TestAddPlaneHandler(t *testing.T) {
	//utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	//Add Success
	json1 := []byte(`{"id":4, "name":"F21JET","wheels": 6,"engines":4,"type":"Attack"}`)
	req, _ := http.NewRequest("POST", ser.URL+"/plane", bytes.NewBuffer(json1))
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	// Decoding Error
	json1 = []byte(`{"id":4, "name":"F21JET","wheels": "6","engines":"4","type":types"}`)
	req, _ = http.NewRequest("POST", ser.URL+"/plane", bytes.NewBuffer(json1))
	req.Header.Set("Content-Type", "application/json")
	res, _ = http.DefaultClient.Do(req)
	assert.Error(t, errors.New("Error while Decoding Body"), "")

	//Bad request
	assert.HTTPError(t, AddPlane(sess, "testing"), "DELETE", "http://localhost:8081/planes", nil)

	// Validation Error
	json1 = []byte(`{}`)
	req, _ = http.NewRequest("POST", ser.URL+"/plane", bytes.NewBuffer(json1))
	req.Header.Set("Content-Type", "application/json")
	res, _ = http.DefaultClient.Do(req)
	assert.Equalf(t, 400, res.StatusCode, "Expected %d but got %d ", 400, res.StatusCode)

}

func TestRemovePlaneByName(t *testing.T) {
	//utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// Remove Success
	req, _ := http.NewRequest("DELETE", ser.URL+"/plane/F21JET", nil)
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	// wrong Request Method
	assert.HTTPErrorf(t, RemovePlaneByName(sess, "testing"), "POST", "http://localhost:8081/plane/F21", nil, "")

	// invalid plane name to delete
	assert.HTTPErrorf(t, RemovePlaneByName(sess, "testing"), "DELETE", "http://localhost:8081/plane/asdfsds", nil, "")
}

func TestGetPlanesHandler(t *testing.T) {
	//utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	//success test
	assert.HTTPSuccessf(t, GetPlanesHandler(sess, "testing"), "GET", "http://localhost:8081/planes", nil, "")
	// bad request test
	assert.HTTPErrorf(t, GetPlanesHandler(sess, "testing"), "PATCH", "http://localhost:8081/planes", nil, "")
}

func TestAddStudentHandelr(t *testing.T) {
	// utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// AddStudent success
	json1 := []byte(`{"studentName":"Eshan", "studentAge":"24","studentMarks": "99"}`)
	req, _ := http.NewRequest("POST", ser.URL+"/student", bytes.NewBuffer(json1))
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	// AddStudentErr Bad Request
	assert.HTTPErrorf(t, AddStudent(sess, "testing"), "GET", "http://localhost:8081/student", nil, "")

	// AddStudemtErr validation Error
	emptyJSON := []byte(`{}`)
	req, _ = http.NewRequest("POST", ser.URL+"/student", bytes.NewBuffer(emptyJSON))
	res, _ = http.DefaultClient.Do(req)
	assert.Equalf(t, 422, res.StatusCode, "Expected %d but got %d", 404, res.StatusCode)

	// AddStudentErr Decoding Error
	errJSON := []byte(`{"}`)
	req, _ = http.NewRequest("POST", ser.URL+"/student", bytes.NewBuffer(errJSON))
	res, _ = http.DefaultClient.Do(req)
	assert.Equalf(t, 400, res.StatusCode, "Expected %d but got %d", 400, res.StatusCode)
}

func TestGetStudentByNameHandler(t *testing.T) {
	// utils.InitConfig()
	//TestGetStudent
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()
	req, _ := http.NewRequest("GET", ser.URL+"/student/Eshan", nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, 200, res.StatusCode, "Expected %d but got %v", 200, res.StatusCode)
	//assert.HTTPSuccessf(t, GetStudentByName(sess, "testing"), "GET", "http://localhost:8081/student/test", nil, "")
	//TestGetStudentErrs
	req, _ = http.NewRequest("GET", ser.URL+"/student/jdfhsdjfhks", nil)
	res, _ = http.DefaultClient.Do(req)
	//assert.HTTPErrorf(t, GetStudentByName(sess, "testing"), "GET", "http://localhost:8081/student/jdfhsdjfhks", nil, "")
	assert.Equal(t, 404, res.StatusCode, "Expected %d but got %v", 404, res.StatusCode)

	// TesGetStudentErr PATCH method
	assert.HTTPErrorf(t, GetStudentByName(sess, "testing"), "PATCH", "http://localhost:8081/student/jdfhsdjfhks", nil, "")
}

func TestGetAllStudentsHandler(t *testing.T) {
	// utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// success get students
	assert.HTTPSuccessf(t, GetAllStudents(sess, "testing"), "GET", "http://localhost:8081/students", nil, "")

	// error request method

	assert.HTTPError(t, GetAllStudents(sess, "testing"), "POST", "http://localhost:8081/students", nil, "")
}

func TestDeleteStudentHandler(t *testing.T) {
	// utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// DeleteStudent success
	//assert.HTTPSuccess(t, DeleteStudent(sess, "testing"), "DELETE", "http://localhost:8081/student/Eshan", nil, "")
	req, _ := http.NewRequest("DELETE", ser.URL+"/student/Eshan", nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, 200, res.StatusCode, "")

	// DeleteStudentErr Invalid Input
	req, _ = http.NewRequest("DELETE", ser.URL+"/student/dhasdjhasj", nil)
	res, _ = http.DefaultClient.Do(req)
	//assert.HTTPErrorf(t, DeleteStudent(sess, "testing"), "DELETE", "http://localhost:8081/student/dhasdjhasj", nil, "")
	assert.Equal(t, 404, res.StatusCode, "")
	// DeleteStudentErr Request Method
	assert.HTTPErrorf(t, DeleteStudent(sess, "testing"), "GET", "http://localhost:8081/student/dhasdjhasj", nil, "")

}

func TestUpdateStudentHandler(t *testing.T) {
	// utils.InitConfig()
	sess, _ := utils.GetDataBaseSessionWithURI("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()
	s := model.Student{
		StudentName:  "Eshan",
		StudentAge:   25,
		StudentMarks: 88,
	}
	_ = dao.AddStudent(s, sess, "testing")

	//Successfull Update
	json1 := []byte(`{"studentName":"Eshan", "studentAge":"24","studentMarks": "99"}`)
	req, _ := http.NewRequest("PUT", ser.URL+"/student/Eshan", bytes.NewBuffer(json1))
	res, _ := http.DefaultClient.Do(req)
	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	//Invalid Name to Update
	json1 = []byte(`{"studentName":"Eshan", "studentAge":"24","studentMarks": "99"}`)
	req, _ = http.NewRequest("PUT", ser.URL+"/student/dfsdf", bytes.NewBuffer(json1))
	res, _ = http.DefaultClient.Do(req)
	assert.Equalf(t, 404, http.StatusNotFound, "Expected %d but got %d ", 404, http.StatusNotFound)

	// UpdateStuden Bad Request
	assert.HTTPErrorf(t, UpdateStud(sess, "testing"), "GET", "http://localhost:8081/student/rewr", nil, "")
}
