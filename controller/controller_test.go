package controller

import (
	"RestApp/model"
	"net/http/httptest"
	"RestApp/dao"
	"net/http"
	"testing"
	"encoding/json"
	"bytes"
	"github.ibm.com/dash/dash_utils/dashtest"
	"fmt"
)

func TestRemoveStudent(t * testing.T){
	tst := []struct{
		m 		string
		uri 	string
		status 	int
	}{
		{
			m 		: "DELETE",
			uri 	:"/remStud/Sagar",
			status	:200,
		},
		{
			m		:"DELETE",
			uri 	:"/remStud/Dio",
			status	:200,
		},
	}

	dao.Init()

	srv := httptest.NewServer(Handlers())


	for i := range(tst){
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil)
		if err!= nil{
			t.Errorf("New Request cound not be generated : %v\n", err)
			 	t.Errorf("Error Creating a Test Request : %v",err)
		}

		res,err := http.DefaultClient.Do(req)
		if err!= nil{
			t.Errorf("Error Requesting the Test Server : %v",err)
		}
		
		if res.StatusCode != tst[i].status{
			t.Errorf("Did not receive expected status. Got : %v \t Want: %v",res.StatusCode,tst[i].status)
		}
		   


		
	}
}

func TestRemoveStudentErr(t * testing.T){
	tst := []struct{
		m 		string
		uri 	string
		status 	int
	}{
		{
			m 		: "GET",
			uri 	:"/remStud/sad",
			status	:200,
		},
		{
			m		:"DELETE",
			uri 	:"/remStud/Light gYaami",
			status	:300,
		},
	}

	dao.Init()

	srv := httptest.NewServer(Handlers())


	for i := range(tst){
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil)
		if err!= nil{
			t.Errorf("New Request cound not be generated : %v\n", err)
			 	t.Errorf("Error Creating a Test Request : %v",err)
		}

		res,err := http.DefaultClient.Do(req)
		if err!= nil{
			t.Errorf("Error Requesting the Test Server : %v",err)
		}
		
		if res.StatusCode != tst[i].status{
			fmt.Printf("Received Expected Wrong status expected status. Got : %v \t Want: %v",res.StatusCode,tst[i].status)
		}
		   


		
	}
}

func TestConnect(t *testing.T){
	dao.Init()
}

func TestUpdateStudentErr(t *testing.T){
	tst := [] struct{
		m 		string
		uri 	string
		stat 	int
		stu  	model.Student
	}{
		{
			m 		: "PUT",
			uri 	: "/upStud/Rushikesh Kin",
			stat 	: 300,
			stu 	: model.Student{
				StudentName: "Rushikesh Kinhalkar",
				StudentAge: 20,
				StudentMarks:89,
			},
		},{
			m 		: "GET",
			uri 	: "/upStud/Devansh",
			stat 	: 300,
			stu 	: model.Student{
				StudentName: "Devansh",
				StudentAge: 10,
				StudentMarks:10,
			},
		},
	}

	srv := httptest.NewServer(Handlers())

	defer srv.Close()
	dao.Init()

	for i := range(tst){
		body := &bytes.Buffer{}
		json.NewEncoder(body).Encode(&tst[i].stu)
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil); if err!=nil{
			t.Errorf("Could not create request : %v",err)
			t.Fatal(err)
		}
		
		res,err := http.DefaultClient.Do(req); if err!=nil{
			t.Errorf("Could not Satisfy request: %v",err)
		}

		if res.StatusCode != tst[i].stat{
			fmt.Printf("Bad Code was expected! got : %v \t want: %v",res.StatusCode,tst[i].stat)
		}
	}
}

func TestUpdateStudent(t *testing.T){
	tst := [] struct{
		m 		string
		uri 	string
		stat 	int
		stu  	model.Student
	}{
		{
			m 		: "PUT",
			uri 	: "/upStud/Rushikesh Kinhalkar",
			stat 	: 200,
			stu 	: model.Student{
				StudentName: "Rushikesh Kinhalkar",
				StudentAge: 20,
				StudentMarks:89,
			},
		},{
			m 		: "PUT",
			uri 	: "/upStud/Devansh",
			stat 	: 200,
			stu 	: model.Student{
				StudentName: "Devansh",
				StudentAge: 10,
				StudentMarks:10,
			},
		},
	}

	srv := httptest.NewServer(Handlers())

	defer srv.Close()
	dao.Init()

	for i := range(tst){
		body := &bytes.Buffer{}
		json.NewEncoder(body).Encode(&tst[i].stu)
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil); if err!=nil{
			t.Errorf("Could not create request : %v",err)
			t.Fatal(err)
		}
		
		res,err := http.DefaultClient.Do(req); if err!=nil{
			t.Errorf("Could not Satisfy request: %v",err)
		}

		if res.StatusCode != tst[i].stat{
			t.Errorf("Code was not expected! got : %v \t want: %v",res.StatusCode,tst[i].stat)
		}
	}
}

func TestGetAllStudents( t *testing.T ) {
	req,err := http.NewRequest("GET","/getAll",nil)// create new request

	if err!=nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()// recoder to record the response
	handler := http.HandlerFunc(GetAllStudents)// specify handlerFunc function
	handler.ServeHTTP(rr,req)//serv the request to the Handler Func

	status := rr.Code; if status != http.StatusOK{ //check for status 200
		t.Errorf("Error Bad Status! got : %v \t want : %v",status,http.StatusOK)
	}

}



func TestGetAllStudentsErr( t *testing.T ) {
	req,err := http.NewRequest("POST","/getAll",nil)// create new request

	if err!=nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()// recoder to record the response
	handler := http.HandlerFunc(GetAllStudents)// specify handlerFunc function
	handler.ServeHTTP(rr,req)//serv the request to the Handler Func

	status := rr.Code; if status != http.StatusOK{ //check for status 200
		fmt.Printf("Expected Bad Status! got : %v \t want : %v",status,http.StatusOK)
	}
}



func TestGetByName(t *testing.T){
	tst := []struct{
		Method	string
		uri 	string
		stat 	int
	}{
		{
		Method	:"GET",
		uri	:"/getStud/Meera",
		stat:	200,
		},
		{
		Method	:"GET",
		uri	:"/getStud/Hatchiko",
		stat:	200,
		},	
	}


	dao.Init()
	srv := httptest.NewServer(Handlers())
	defer srv.Close()
	for i := range(tst){
		req, err := http.NewRequest(tst[i].Method,srv.URL+tst[i].uri,nil)

		if err!=nil{
			t.Errorf("Bad Request : %v",err)
		}

		res ,err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Request wasn't served!: %v",err)
		}

		if res.StatusCode != tst[i].stat{
			t.Error("Bad Status returned")
		}
	
	}
}

func TestGetByNameErr(t *testing.T){
	tst := []struct{
		Method	string
		uri 	string
		stat 	int
	}{
		{
		Method	:"PUT",
		uri	:"/getStud/Devansh",
		stat:	200,
		},
		{
		Method	:"POST",
		uri	:"/getStud/sad",
		stat:	500,
		},	
	}


	dao.Init()
	srv := httptest.NewServer(Handlers())
	defer srv.Close()
	for i := range(tst){
		req, err := http.NewRequest(tst[i].Method,srv.URL+tst[i].uri,nil)

		if err!=nil{
			t.Errorf("Bad Request : %v",err)
		}

		res ,err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Request wasn't served!: %v",err)
		}

		if res.StatusCode != tst[i].stat{
			fmt.Printf("Bad Status returned ! got :%v\t want: %v",res.StatusCode,tst[i].stat)
		}
	
	}
}

func TestAddStudent(t *testing.T){
	ts := []struct{
		m	string
		uri string
		st 	int
		stu model.Student
	}{
		{	
			m : "POST",
			uri : "/stud",
			st : 200,
			stu : model.Student {
				StudentName:"Apoorva",
				StudentAge:10,
				StudentMarks:89,
			},
		},
		{
			m 	: "POST",
			uri : "/stud",
			st 	: 300,
			stu : model.Student{
				StudentName: "Nitisha",
				StudentAge : 80,
				StudentMarks: 37,
			},
		},
	}



	dao.Init()
	srv := httptest.NewServer(Handlers())

	defer srv.Close()

	for i := range ts{
		body := &bytes.Buffer{}
		json.NewEncoder(body).Encode(ts[i].stu)
		req, err := http.NewRequest(ts[i].m,srv.URL+ts[i].uri,body)
		if err !=nil{
			t.Errorf("Error Occured creating a request! : %v",err)
		}

		res,err := http.DefaultClient.Do(req)

		if err != nil{
			t.Errorf("Error while Request! : %v ",err)
		}


		if res.StatusCode != http.StatusOK{
			t.Errorf("Server returned wrong code!check your Request! Got : %v \t Want: %v",res.Status,ts[i].st)
		}


	}
}


func TestAddStudentErr(t *testing.T){
	ts := []struct{
		m	string
		uri string
		st 	int
		stu model.Student
	}{
		{	
			m : "DELETE",
			uri : "/stud",
			st : 200,
			stu : model.Student {
				StudentName:"Light Yagami",
				StudentAge:9,
				StudentMarks:99,
			},
		},
		{
			m 	: "GET",
			uri : "/stud",
			st 	: 300,
			stu : model.Student{
				StudentName: "Hatchiko",
				StudentAge : 70,
				StudentMarks: 67,
			},
		},
	}



	dao.Init()
	srv := httptest.NewServer(Handlers())

	defer srv.Close()

	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(ts[0].stu)
	req, err := http.NewRequest(ts[0].m,srv.URL+ts[0].uri,body)
	if err !=nil{
		t.Errorf("Error Occured creating a request! : %v",err)
	}

	res,err := http.DefaultClient.Do(req)

	if err != nil{
		t.Errorf("Error while Request! : %v ",err)
	}


	if res.StatusCode != http.StatusOK{
		fmt.Printf("Expected Bad code ! Got : %v \t Want: %v",res.Status,http.StatusOK)
	}


	
}



func TestMain(m * testing.M){
	dashtest.ControlCoverage(m)
}

	

	
