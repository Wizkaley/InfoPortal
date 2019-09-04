package controller

import (
	"RestApp/model"
	"net/http/httptest"
	"RestApp/dao"
	"net/http"
	"testing"
	"encoding/json"
	"bytes"
)

func TestRemoveStudent(t * testing.T){
	tst := []struct{
		m 		string
		uri 	string
		status 	int
	}{
		{
			m 		: "DELETE",
			uri 	:"/remStud/Akash",
			status	:200,
		},
		{
			m		:"DELETE",
			uri 	:"/remStud/Eshan",
			status	:200,
		},
	}

	dao.Init()

	srv := httptest.NewServer(Handlers())


	for i := range(tst){
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil)
		if err!= nil{
			t.Fatalf("New Request cound not be generated : %v\n", err)
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


func TestConnect(t *testing.T){
	dao.Init()
}

func TestUpdateStudent(t *testing.T){
	tst:=[] struct{
		m 		string
		uri 	string
		stat	int
	}{
		{
			m : "PUT",
			uri : "/upStud/Aman Shukla,Aman",
			stat : 200,
		},
		{
			m : "PUT",
			uri : "/upStud/Gaurav Gavanakar,Gaurav",
			stat : 200,
		},
	}
	dao.Init()

	srv := httptest.NewServer(Handlers())

	for i := range(tst){
		req, err := http.NewRequest(tst[i].m,srv.URL+tst[i].uri,nil)

		if err!=nil{
			t.Errorf("Bad Request Constructed! : %v",err)
		}

		res, err := http.DefaultClient.Do(req)

		if res.StatusCode != tst[i].stat{
			t.Errorf("Server Returned Bad Status Code! Got : %v \t Want : %v",res.StatusCode,tst[i].stat)
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


	// expected := `[{"studentName":"Harry","studentAge":22,"studentMarks":70},{"studentName":"Eshan","studentAge":18,"studentMarks":99},{"studentName":"Gaurav","studentAge":25,"studentMarks":90}]`

	// if rr.Body.String() != expected{
	// 	t.Errorf("Error bad result got : %v \t expected : %v",rr.Body.String(),expected )
	// }
}

//TestConnect(t *testing.T)

// func TestGetByName(t *testing.T){

// 	dao.Init()
// 	//srv := httptest.NewServer(Handlers())
// 	//defer srv.Close()
// 	//str := []byte(`{"nm":"Eshan"}`)
// 	req := httptest.NewRequest("GET","/getStud/Eshan",nil)

// 	//q := req.URL.Query()
// 	//fmt.Println("Parsing Noww.........",req.URL.Parse())
// 	///q := req.FormValue("nm")
// 	//fmt.Println("................................value of q",q)
// 	//uri := req.RequestURI
// 	//fmt.Println("Requested URI..........................",uri)
// 	//fmt.Println(req.URL.Path)

// 	//q.Add("nm","Eshan")
// 	//q.Set("nm","Eshan")
// 	//q.Add("/nm","Eshan")
// 	//req.URL.RawQuery = q.Encode()
// 	//req.URL.
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetByName)
// 	fmt.Println("............................1")
// 	//fmt.Println(srv.URL)
// 	//fmt.Print(req.URL.RawPath)
	
// 	handler.ServeHTTP(rr,req)
// 	fmt.Print("............................2")

// 	status := rr.Code; if status!= http.StatusOK{
// 		t.Errorf("Wrong Request! got : %v \t want : %v",status,http.StatusOK)
// 	}
// 	fmt.Print("............................3")
// 	expected := `{"studentName":"Eshan","studentAge":18,"studentMarks":99}`
// 	if rr.Body.String() != expected{
// 		t.Errorf("Didn't receive exected Results! got : %v, expected : %v",rr.Body.String(),expected)
// 	}




// }

func TestGetByName(t *testing.T){
	tst := []struct{
		Method	string
		uri 	string
		stat 	int
	}{
		{
		Method	:"GET",
		uri	:"/getStud/Devansh",
		stat:	200,
		},	
	}
	dao.Init()
	srv := httptest.NewServer(Handlers())
	defer srv.Close()

	req, err := http.NewRequest(tst[0].Method,srv.URL+tst[0].uri,nil)

	if err!=nil{
		t.Errorf("Bad Request : %v",err)
	}

	res ,err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Request wasn't served!: %v",err)
	}

	if res.StatusCode != tst[0].stat{
		t.Error("Bad Status returned")
	}
	
	
}
// func TestGetByName(t * testing.T){

// 	tstSuite := [] struct{
// 		method string
// 		uri string
// 		status int
// 	}{
// 		{
// 			method: "GET",
// 			uri:"/getStud/Eshan",
// 			status: 200,
// 		},
// 		{
// 			method:"GET",
// 			uri:"/getStud/Harry",	
// 			status:200,
// 		},
// 	}
// 		for i:= range tstSuite{
// 		dao.Init()
// 		srv := httptest.NewServer(Handlers())

// 		defer srv.Close()

// 		req, err := http.NewRequest(tstSuite[i].method,srv.URL+tstSuite[i].uri,nil)
// 		fmt.Println(srv.URL+tstSuite[i].uri)
// 		if err != nil{
// 			t.Errorf("New Req could not be complete: %v",err)
// 		}

// 		res,_:=http.DefaultClient.Do(req)
// 		if res.StatusCode != tstSuite[i].status{
// 			t.Errorf("Expected status ok, got different status got: %v \t want: %v",res.StatusCode,tstSuite[i].status)
// 		}

// 		fmt.Println(res.Body)
// 	}
// }


// func TestAddStudent(t *testing.T){


// 	pt := []byte(`{"studentName":"Rushikesh Kinhalkar","studentAge":01,"studentMarks":80}`)

// 	req := httptest.NewRequest("POST","/stud",bytes.NewBuffer(pt))

// 	rr := httptest.NewRecorder()

// 	//q := req.URL.Query()

// 	//q.Add("name":"")

// 	req.Header.Set("Content-type","application/json")
// 	handler := http.HandlerFunc(AddStudent)
// 	handler.ServeHTTP(rr,req)

// 	if rr.Code != 200{
// 		t.Errorf("Bad Request, Student Not Added")
// 	}



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
				StudentName:"Apoorva Joshi",
				StudentAge:30,
				StudentMarks:98,
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
		t.Errorf("Server returned wrong code!check your Request! Got : %v \t Want: %v",res.Status,http.StatusOK)
	}


	
}





	

	
