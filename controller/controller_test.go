package controller

import (
	"RESTApp/utils"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

}

func TestAddPlaneHandler(t *testing.T) {
	sess, _ := utils.GetDataBaseSession("localhost:27017")
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

// func TestRemovePlaneByIDHandler(t *testing.T) {
// 	sess, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer sess.Close()
// 	ser := httptest.NewServer(Handlers(sess, "testing"))
// 	defer ser.Close()

// 	// success
// 	req, _ := http.NewRequest("DELETE", ser.URL+"/plane/2", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	res, _ := http.DefaultClient.Do(req)
// 	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)
// }

func TestRemovePlaneByName(t *testing.T) {

	sess, _ := utils.GetDataBaseSession("localhost:27017")
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
	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	//success test
	assert.HTTPSuccessf(t, GetPlanesHandler(sess, "testing"), "GET", "http://localhost:8081/planes", nil, "")
	// bad request test
	assert.HTTPErrorf(t, GetPlanesHandler(sess, "testing"), "PATCH", "http://localhost:8081/planes", nil, "")
}

func TestAddStudentHandelr(t *testing.T) {

	sess, _ := utils.GetDataBaseSession("localhost:27017")
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
	//TestGetStudent
	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()
	req, _ := http.NewRequest("GET", ser.URL+"/student/Eshan", nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, 200, res.StatusCode, "Expected %d but got %", 200, res.StatusCode)
	//assert.HTTPSuccessf(t, GetStudentByName(sess, "testing"), "GET", "http://localhost:8081/student/test", nil, "")
	//TestGetStudentErrs
	assert.HTTPErrorf(t, GetStudentByName(sess, "testing"), "GET", "http://localhost:8081/student/jdfhsdjfhks", nil, "")

	// TesGetStudentErr PATCH method
	assert.HTTPErrorf(t, GetStudentByName(sess, "testing"), "PATCH", "http://localhost:8081/student/jdfhsdjfhks", nil, "")
}

func TestGetAllStudentsHandler(t *testing.T) {
	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// success get students
	assert.HTTPSuccessf(t, GetAllStudents(sess, "testing"), "GET", "http://localhost:8081/students", nil, "")

	// error request method

	assert.HTTPError(t, GetAllStudents(sess, "testing"), "POST", "http://localhost:8081/students", nil, "")
}

func TestDeleteStudentHandler(t *testing.T) {

	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

	// DeleteStudent success
	//assert.HTTPSuccess(t, DeleteStudent(sess, "testing"), "DELETE", "http://localhost:8081/student/Eshan", nil, "")
	req, _ := http.NewRequest("DELETE", ser.URL+"/student/Eshan", nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, 200, res.StatusCode, "")

	// DeleteStudentErr Invalid Input
	assert.HTTPErrorf(t, DeleteStudent(sess, "testing"), "DELETE", "http://localhost:8081/student/dhasdjhasj", nil, "")
	// DeleteStudentErr Request Method
	assert.HTTPErrorf(t, DeleteStudent(sess, "testing"), "GET", "http://localhost:8081/student/dhasdjhasj", nil, "")

}

func TestUpdateStudentHandler(t *testing.T) {
	sess, _ := utils.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()

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

// func TestGetSwagger(t *testing.T) {
// 	sess, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer sess.Close()
// 	ser := httptest.NewServer(Handlers(sess, "testing"))
// 	defer ser.Close()

// 	req, _ := http.NewRequest("GET", ser.URL+"/swagger", nil)
// 	res, _ := http.DefaultClient.Do(req)
// 	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)
// 	//assert.HTTPSuccess(t, GetSwagger, "GET", "/swagger", nil, nil)
// }

// func TestUpdateStudentErr(t *testing.T) {
// 	tst := []struct {
// 		m    string
// 		uri  string
// 		stat int
// 		stu  model.Student
// 	}{
// 		{
// 			m:    "PUT",
// 			uri:  "/upStud/Rushikesh Kin",
// 			stat: 404,
// 			stu: model.Student{
// 				StudentName:  "Rushikesh Kinhalkar",
// 				StudentAge:   20,
// 				StudentMarks: 89,
// 			},
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()

// 	for i := range tst {
// 		body := &bytes.Buffer{}
// 		json.NewEncoder(body).Encode(&tst[i].stu)
// 		req, err := http.NewRequest(tst[i].m, srv.URL+tst[i].uri, nil)
// 		if err != nil {
// 			t.Errorf("Could not create request : %v", err)
// 			t.Fatal(err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Errorf("Could not Satisfy request: %v", err)
// 		}

// 		if res.StatusCode != tst[i].stat {
// 			fmt.Printf("Bad Code was expected! got : %v \t want: %v", res.StatusCode, tst[i].stat)
// 		}
// 	}
// }

// func TestUpdateStudent(t *testing.T) {
// 	tst := []struct {
// 		m    string
// 		uri  string
// 		stat int
// 		stu  model.Student
// 	}{
// 		{
// 			m:    "PUT",
// 			uri:  "/upStud/Rushikesh Kinhalkar",
// 			stat: 200,
// 			stu: model.Student{
// 				StudentName:  "Rushikesh Kinhalkar",
// 				StudentAge:   20,
// 				StudentMarks: 89,
// 			},
// 		}, {
// 			m:    "PUT",
// 			uri:  "/upStud/Devansh",
// 			stat: 200,
// 			stu: model.Student{
// 				StudentName:  "Devansh",
// 				StudentAge:   10,
// 				StudentMarks: 10,
// 			},
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()

// 	for i := range tst {
// 		body := &bytes.Buffer{}
// 		json.NewEncoder(body).Encode(&tst[i].stu)
// 		req, err := http.NewRequest(tst[i].m, srv.URL+tst[i].uri, nil)
// 		if err != nil {
// 			t.Errorf("Could not create request : %v", err)
// 			t.Fatal(err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Errorf("Could not Satisfy request: %v", err)
// 		}

// 		if res.StatusCode != tst[i].stat {
// 			t.Errorf("Code was not expected! got : %v \t want: %v", res.StatusCode, tst[i].stat)
// 		}
// 	}
// }

// func TestGetAllStudents(t *testing.T) {

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()

// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()
// 	req, err := http.NewRequest("GET", "/getAll", nil) // create new request

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	//rr := httptest.NewRecorder()
// 	// recoder to record the response
// 	//	handler := http.HandlerFunc(GetAllStudents(s))      // specify handlerFunc function
// 	//handler.ServeHTTP(rr, req)                          //serv the request to the Handler Func

// 	status := res.StatusCode
// 	if status != http.StatusOK { //check for status 200
// 		t.Errorf("Error Bad Status! got : %v \t want : %v", status, http.StatusOK)
// 	}

// }

// func TestGetAllStudentsErr(t *testing.T) {

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()

// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()
// 	req, err := http.NewRequest("POST", "/getAll", nil) // create new request

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	//rr := httptest.NewRecorder()                // recoder to record the response
// 	//handler := http.HandlerFunc(GetAllStudents) // specify handlerFunc function
// 	//handler.ServeHTTP(rr, req)                  //serv the request to the Handler Func

// 	res, err := http.DefaultClient.Do(req)
// 	status := res.StatusCode
// 	if status != http.StatusOK { //check for status 200
// 		fmt.Printf("Expected Bad Status! got : %v \t want : %v", status, http.StatusOK)
// 	}
// }

// func TestGetByName(t *testing.T) {
// 	tst := []struct {
// 		Method string
// 		uri    string
// 		stat   int
// 	}{
// 		{
// 			Method: "GET",
// 			uri:    "/getStud/Meera",
// 			stat:   200,
// 		},
// 		{
// 			Method: "GET",
// 			uri:    "/getStud/Hatchiko",
// 			stat:   200,
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))
// 	defer srv.Close()
// 	for i := range tst {
// 		req, err := http.NewRequest(tst[i].Method, srv.URL+tst[i].uri, nil)

// 		if err != nil {
// 			t.Errorf("Bad Request : %v", err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Fatalf("Request wasn't served!: %v", err)
// 		}

// 		if res.StatusCode != tst[i].stat {
// 			t.Error("Bad Status returned")
// 		}

// 	}
// }

// func TestGetByNameErr(t *testing.T) {
// 	tst := []struct {
// 		Method string
// 		uri    string
// 		stat   int
// 	}{
// 		{
// 			Method: "PUT",
// 			uri:    "/getStud/Devansh",
// 			stat:   200,
// 		},
// 		{
// 			Method: "POST",
// 			uri:    "/getStud/sad",
// 			stat:   500,
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))
// 	defer srv.Close()
// 	for i := range tst {
// 		req, err := http.NewRequest(tst[i].Method, srv.URL+tst[i].uri, nil)

// 		if err != nil {
// 			t.Errorf("Bad Request : %v", err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Fatalf("Request wasn't served!: %v", err)
// 		}

// 		if res.StatusCode != tst[i].stat {
// 			fmt.Printf("Bad Status returned ! got :%v\t want: %v", res.StatusCode, tst[i].stat)
// 		}

// 	}
// }

// func TestAddStudent(t *testing.T) {
// 	ts := []struct {
// 		m   string
// 		uri string
// 		st  int
// 		stu model.Student
// 	}{
// 		{
// 			m:   "POST",
// 			uri: "/stud",
// 			st:  200,
// 			stu: model.Student{
// 				StudentName:  "Apoorva",
// 				StudentAge:   10,
// 				StudentMarks: 89,
// 			},
// 		},
// 		{
// 			m:   "POST",
// 			uri: "/stud",
// 			st:  300,
// 			stu: model.Student{
// 				StudentName:  "Nitisha",
// 				StudentAge:   80,
// 				StudentMarks: 37,
// 			},
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()

// 	for i := range ts {
// 		body := &bytes.Buffer{}
// 		json.NewEncoder(body).Encode(ts[i].stu)
// 		req, err := http.NewRequest(ts[i].m, srv.URL+ts[i].uri, body)
// 		if err != nil {
// 			t.Errorf("Error Occured creating a request! : %v", err)
// 		}

// 		res, err := http.DefaultClient.Do(req)

// 		if err != nil {
// 			t.Errorf("Error while Request! : %v ", err)
// 		}

// 		if res.StatusCode != http.StatusOK {
// 			t.Errorf("Server returned wrong code!check your Request! Got : %v \t Want: %v", res.Status, ts[i].st)
// 		}

// 	}
// }

// func TestAddStudentErr(t *testing.T) {
// 	ts := []struct {
// 		m   string
// 		uri string
// 		st  int
// 		stu model.Student
// 	}{
// 		{
// 			m:   "DELETE",
// 			uri: "/stud",
// 			st:  200,
// 			stu: model.Student{
// 				StudentName:  "Light Yagami",
// 				StudentAge:   9,
// 				StudentMarks: 99,
// 			},
// 		},
// 		{
// 			m:   "GET",
// 			uri: "/stud",
// 			st:  300,
// 			stu: model.Student{
// 				StudentName:  "Hatchiko",
// 				StudentAge:   70,
// 				StudentMarks: 67,
// 			},
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	defer srv.Close()

// 	body := &bytes.Buffer{}
// 	json.NewEncoder(body).Encode(ts[0].stu)
// 	req, err := http.NewRequest(ts[0].m, srv.URL+ts[0].uri, body)
// 	if err != nil {
// 		t.Errorf("Error Occured creating a request! : %v", err)
// 	}

// 	res, err := http.DefaultClient.Do(req)

// 	if err != nil {
// 		t.Errorf("Error while Request! : %v ", err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("Expected Bad code ! Got : %v \t Want: %v", res.Status, http.StatusOK)
// 	}

// }
