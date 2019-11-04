package controller

import (
	"RESTApp/utils/mongo"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	sess, _ := mongo.GetDataBaseSession("localhost:27017")
	defer sess.Close()
	ser := httptest.NewServer(Handlers(sess, "testing"))
	defer ser.Close()
	//TestGetStudent
	//req1, _ := http.NewRequest("GET", ser.URL+"/student/test", nil)
	//req2, _ := http.NewRequest("PATCH", ser.URL+"/student/test", nil)
	req3, _ := http.NewRequest("GET", ser.URL+"/student/dwawsdsasadasdsaa", nil)
	// res1, _ := http.DefaultClient.Do(req1)
	// res2, _ := http.DefaultClient.Do(req2)
	req3.Header.Set("Content-type", "application/json")
	res3, _ := http.DefaultClient.Do(req3)
	// assert.Equalf(t, 200, res1.StatusCode, "Expected %d but got %d ", 200, res1.StatusCode)

	// assert.Equalf(t, http.StatusMethodNotAllowed, res2.StatusCode, "Expected %d but got %d ", http.StatusMethodNotAllowed, res2.StatusCode)

	assert.Equalf(t, http.StatusNotFound, res3.StatusCode, "Expected %d but got %d ", http.StatusNotFound, res3.StatusCode)

	assert.HTTPErrorf(t, DeleteStudent(sess, "testing"), "GET", "http://localhost:8081/student/dhasdjhasj", nil, "")
	// stu := model.Student{
	// 	StudentName:  "Eshan",
	// 	StudentAge:   24,
	// 	StudentMarks: 99,
	// }
	json1 := []byte(`{"studentName":"Eshan", "studentAge":"24","studentMarks": "99"}`)
	req, _ := http.NewRequest("POST", ser.URL+"/student", bytes.NewBuffer(json1))
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	// TestRemoveStudent
	// req, _ = http.NewRequest("DELETE", ser.URL+"/student/test", nil)
	// res, _ = http.DefaultClient.Do(req)
	// assert.Equalf(t, 200, res.StatusCode, "Expected %d but got %d ", 200, res.StatusCode)

	//stu := model.Student{
	// 	StudentName:  "Eshan",
	// 	StudentAge:   24,
	// 	StudentMarks: 99,
	// }
	//var students map[string]model.Student

	//assert.HTTPSuccess(t, DeleteStudent(sess, "testing"), "DELETE", "http://localhost:8081/student/test", nil)

	// Test Remove Student
	//assert.HTTPSuccess(t, AddStudent(sess, "testing"), "POST", "http://localhost:8081/student", url.Values(stu), "")
	//assert.
}

// func TestRemoveStudent(t *testing.T) {
// 	tst := []struct {
// 		m      string
// 		uri    string
// 		status int
// 	}{
// 		{
// 			m:      "DELETE",
// 			uri:    "/remStud/Pretty",
// 			status: 200,
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	for i := range tst {
// 		req, err := http.NewRequest(tst[i].m, srv.URL+tst[i].uri, nil)
// 		if err != nil {
// 			t.Errorf("New Request cound not be generated : %v\n", err)
// 			t.Errorf("Error Creating a Test Request : %v", err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Errorf("Error Requesting the Test Server : %v", err)
// 		}

// 		if res.StatusCode != tst[i].status {
// 			t.Errorf("Did not receive expected status. Got : %v \t Want: %v", res.StatusCode, tst[i].status)
// 		}

// 	}
// }

// func TestRemoveStudentErr(t *testing.T) {
// 	tst := []struct {
// 		m      string
// 		uri    string
// 		status int
// 	}{
// 		{
// 			m:      "DELETE",
// 			uri:    "/remStud/Light gYaami",
// 			status: 404,
// 		},
// 	}

// 	s, _ := mongo.GetDataBaseSession("localhost:27017")
// 	defer s.Close()
// 	srv := httptest.NewServer(Handlers(s))

// 	for i := range tst {
// 		req, err := http.NewRequest(tst[i].m, srv.URL+tst[i].uri, nil)
// 		if err != nil {
// 			t.Errorf("New Request cound not be generated : %v\n", err)
// 			t.Errorf("Error Creating a Test Request : %v", err)
// 		}

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			t.Errorf("Error Requesting the Test Server : %v", err)
// 		}

// 		if res.StatusCode != tst[i].status {
// 			fmt.Printf("Received Expected Wrong status expected status. Got : %v \t Want: %v", res.StatusCode, tst[i].status)
// 		}

// 	}
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
