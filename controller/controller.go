package controller

import(
	//"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"RestApp/model"
	"RestApp/dao"
	//"errors"
	"encoding/json"
	"log"
)

//Handlers ...
func Handlers()http.Handler{
	server := mux.NewRouter()//create a new Server and attach handlers to it
	//meths := []string{"POST","OPTIONS",}
	server.HandleFunc("/stud",AddStudent).Methods("POST")// done
	//server.HandleFunc("/stud",PreflightAddStudent).Methods("OPTIONS")
	server.HandleFunc("/remStud/{name}",DeleteStudent).Methods("DELETE")//done
	server.HandleFunc("/getAll",GetAllStudents).Methods("GET")//done
	server.HandleFunc("/getStud/{nm}",GetByName).Methods("GET")//done
	server.HandleFunc("/upStud/{nm1}",UpdateStud).Methods("PUT")//done

	return server
}




//UpdateStud ...
func UpdateStud(w http.ResponseWriter, r * http.Request){

	//check for method PUT, if anything else, respond with appropriate status
	if r.Method != "PUT"{
		res, err := json.Marshal("Bad Request"); if err!= nil{
			log.Printf("Bad Request : %v",err)
		}

		
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(res)
	}

	if r.Body!= nil{
		var stuNew model.Student //student object to store updated student

		//extract name from path
		params := mux.Vars(r)
		nm := params["nm1"]
	
		defer r.Body.Close()

		//get a studentObject from GetByName using the extracted name
		stuToChange,err:= dao.GetByName(nm); if err!=nil{
			log.Printf("Error While Fetching Record to Update : %v", err)
			res,_ := json.Marshal("Invalid Student Name specified")
			w.Header().Set("Content-type","application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(res) //send Message
			return

		}

		//Decode values from body sent from client into a studentObject 
		err = json.NewDecoder(r.Body).Decode(&stuNew); if err!=nil{
			log.Printf("Error While Deconding Body : %v", err)
		}

		//update the values from the body to The object got from the Database
		//stuToChange.StudentName = stuNew.StudentName
		stuToChange.StudentAge = stuNew.StudentAge //update age
		stuToChange.StudentMarks = stuNew.StudentMarks//update marks


		//respond with appropriate message after calling Data Access Layer
		err = dao.UpdateStudent(stuToChange); if err!=nil{
			log.Printf("Could not Update student: %v",err)
		}

		res,err := json.Marshal("Updated Successfully"); if err!=nil{
			log.Printf("Error While Marshalling! : %v",err)
		}

		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res) //send Message
	}



}
//GetByName ...
func GetByName(w http.ResponseWriter,r *http.Request){

	//check for method GET, if anything else, respond with appropriate status
	if r.Method != "GET"{
		
		res,err := json.Marshal("Bad Request")

		if err!= nil{
			log.Fatal(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, err.Error(), 200)
		w.Write(res)
	}

	params := mux.Vars(r) //extract name from URL path 
	
	var s model.Student
	s,err := dao.GetByName(params["nm"])//call data access layer
	
	if err!=nil{
		res,_ := json.Marshal("No Entry Found By That Name")
		w.Header().Set("Content Type","application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	//respond with appropriate message
	mresult,_ := json.Marshal(s)
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mresult)
}

//AddStudent ...
func AddStudent(w http.ResponseWriter,r *http.Request){
	
	//check if method is POST else show error
		if r.Method != "POST"{
			
			response,_:= json.Marshal("Bad Request")
			w.Header().Set("Content-type","application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(response)
		}
		//check if body has content
		if r.Body != nil{
			defer r.Body.Close()
			var stu model.Student

			//decode the body for student details
			err := json.NewDecoder(r.Body).Decode(&stu)
			if err == nil{
				dao.AddStudent(stu)
				//w.Header().Set("Access-Control-Allow-Methods","POST,OPTIONS")
				response,_ := json.Marshal("Added Successfully")
				w.Header().Set("Content-Type","application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(response)
			}
		}
}

//DeleteStudent ...
func DeleteStudent(w http.ResponseWriter,r * http.Request){
	//check if method is DELETE else respond with error
	if r.Method != "DELETE"{
		response, err  := json.Marshal("Bad Request")
		if err != nil{
			log.Printf("Bad Request!: %v",err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response)
		
	}
	

	//check if body has content
	if r.Body != nil{


		defer r.Body.Close()
		params := mux.Vars(r)//extract name of student from URL path
		
		//err := dao.GetByName(params["name"])
		//Respond to the requeset after calling Data Access Layer
		err := dao.RemoveByName(params["name"])
		if err!=nil{
			res,_ := json.Marshal("Could not Find anyone with that name")
			w.Header().Set("Content-Type","appication/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(res)
			return
		}
		response,err := json.Marshal("Removed Student")
		if err!= nil{
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type","appication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

//GetAllStudents ...
func GetAllStudents(w http.ResponseWriter, r * http.Request){
	
	//check for method GET, if any other, respond with error with appropriate status 
	if r.Method != "GET"{
		response, err := json.Marshal("Bad Request")

		if err!= nil{
			log.Printf("Bad Request!: %v ",err)
		}

		w.Header().Set("Content-Type","application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().S
		//w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response)
	}

	//respond with appropriate message after calling Data Access Layer
	res, err := dao.GetAll(); if err != nil{
	 	log.Fatal(err)
	}
	response, _ := json.Marshal(res)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)	

}
