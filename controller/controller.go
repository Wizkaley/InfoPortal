// Package controller Test Project API's
//
// The purpose of this application is to store and retrieve test records for published test results
//
//
//
//     BasePath: /
//     Version: 1.0.0
//     License: bleh
//
//     Contact: Eshan Kaley<eshkaley@in.ibm.com>
//
//     Consumes:
//       - application/json
//
//     Produces:
//       - application/json
//
//     Security:
//       - token:
//
//     SecurityDefinitions:
//       token:
//         type: apiKey
//         in: header
//         name: Authorization
//
//
// swagger:meta
package controller

import (

	//"fmt"
	"RESTApp/dao"
	"RESTApp/model"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	//"errors"
	"encoding/json"
	"log"
)

//Handlers ...
func Handlers(ds *mgo.Session) http.Handler {
	server := mux.NewRouter() //create a new Server and attach handlers to it
	//meths := []string{"POST","OPTIONS",}

	//Base
	//fs := http.FileServer(http.Dir("/home/wiz/go/src/RESTApp/public"))

	//server.Handle("/public/", http.StripPrefix("/public", fs))

	server.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("/home/wiz/go/src/RESTApp/public/"))))

	//swagger := http.FileServer(http.Dir("static/dist/files"))
	// server.Handle("/dist/files", swagger)
	//server.HandleFunc("/", GetDist).Methods("GET")
	server.HandleFunc("/", redir).Methods("GET")
	server.HandleFunc("/swagger", GetSwagger).Methods("GET")
	server.HandleFunc("/plane", AddPlane(ds)).Methods("POST")
	server.HandleFunc("/planes", GetPlanesHandler(ds)).Methods("GET")
	//server.HandleFunc("/stud",PreflightAddStudent).Methods("OPTIONS")

	//Student Handlers
	server.HandleFunc("/remStud/{name}", DeleteStudent(ds)).Methods("DELETE") //done
	server.HandleFunc("/getAll", GetAllStudents(ds)).Methods("GET")           //done
	server.HandleFunc("/getStud/{nm}", GetByName(ds)).Methods("GET")          //done
	server.HandleFunc("/upStud/{nm1}", UpdateStud(ds)).Methods("PUT")
	server.HandleFunc("/stud", AddStudent(ds)).Methods("POST") // done        //done

	return server
}

// // GetDist ...
// func GetDist(w http.ResponseWriter, r *http.Request) {
// 	s := mux.NewRouter()

// 	http.ServeFile(w, r, )
// }

// GetSwagger ...
func GetSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/wiz/go/src/RESTApp/public/swagger.json")
}

func redir(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8081/public/dist/#/", http.StatusFound)
}

//GetPlanesHandler ...
func GetPlanesHandler(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Error while Marshalling: %v", err)
				sendErr(w, http.StatusMethodNotAllowed, res)
			}
		}

		allPlanes, err := dao.GetAllPlanes(ds)
		if err != nil {
			log.Printf("Error while Fetching Planes : %v ", err)
		}

		res, err := json.Marshal(allPlanes)
		if err != nil {
			log.Printf("Error while Marshalling to send Result")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}

// sendErr helper function
func sendErr(w http.ResponseWriter, stat int, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stat)
	w.Write(res)
}

// AddPlane ...
func AddPlane(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Error while Marshalling : %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(res)
		}

		if r.Body != nil {
			var plane model.Plane
			err := json.NewDecoder(r.Body).Decode(&plane)
			if err != nil {
				log.Printf("Error while Decode body : %v", err)
			}
			dao.PutPlane(plane, ds)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			res, err := json.Marshal("Added Plane Successfully")
			if err != nil {
				log.Printf("Error while Marshalling : %v", err)
			}
			w.Write(res)
			defer r.Body.Close()
		}
	})
}

// RemovePlaneByName ...
func RemovePlaneByName(name string, ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Error while encoding error message : %v", err)
			}
			sendErr(w, http.StatusMethodNotAllowed, res)
		}

		// countCheckSession := ds.Clone()
		// count, err := countCheckSession.DB("trial").C("Student").Count()
		// if err != nil {
		// 	log.Println("Error while getting count from DB : %v", err)
		// }

		// if id > count {
		// 	res, err := json.Marshal("ID with the given value Doesn't Exist")
		// 	if err != nil {
		// 		log.Printf("Error while encoding error message : %v", err)
		// 	}
		// 	sendErr(w, http.StatusBadRequest, res)
		// }

		ok := dao.DeletePlane(name, ds)
		if !ok {
			res, err := json.Marshal("Could Not Delete Server Error")
			if err != nil {
				log.Printf("Error while encoding error message : %v", err)
			}
			sendErr(w, http.StatusInternalServerError, res)
		}
		res, _ := json.Marshal("Deleted Successfully")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}

// RemovePlaneByID ...
func RemovePlaneByID(id int, ds *mgo.Session) http.HandlerFunc {
	// swagger:operation DELETE /removePlane DELETE removePlane
	//
	// Remove Plane
	//
	// Removes a Plane from DB
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   type: integer
	//   description: id of the plane to remove
	//   required: true
	//   in: query
	// responses:
	//   '200':
	//     description: Removed Successfully
	//   '401':
	//     description: Unauthorized, Likely Invalid or Missing Token
	//   '403':
	//     description: Forbidden, you are not allowed to undertake this operation
	//   '404':
	//     description: Not found
	//   '500':
	//     description: Error occurred while processing the request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Error while encoding error message : %v", err)
			}
			sendErr(w, http.StatusMethodNotAllowed, res)
		}

		countCheckSession := ds.Clone()
		count, err := countCheckSession.DB("trial").C("Student").Count()
		if err != nil {
			log.Println("Error while getting count from DB : %v", err)
		}

		if id > count {
			res, err := json.Marshal("ID with the given value Doesn't Exist")
			if err != nil {
				log.Printf("Error while encoding error message : %v", err)
			}
			sendErr(w, http.StatusBadRequest, res)
		}

		ok := dao.DeletePlaneByID(id, ds)
		if !ok {
			res, err := json.Marshal("Could Not Delete Server Error")
			if err != nil {
				log.Printf("Error while encoding error message : %v", err)
			}
			sendErr(w, http.StatusInternalServerError, res)
		}
		res, _ := json.Marshal("Deleted Successfully")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	})
}

//UpdateStud ...
func UpdateStud(ds *mgo.Session) http.HandlerFunc {
	// swagger:operation PUT /upStud/{nm1} UPDATE updateStudent
	//
	// Update a Students Information in The Student Catalog
	// ---
	// description: "Update Student Details"
	// summary: "Update Student Details in the Catalog"
	// parameters:
	// - in: path
	//   name: name
	//   description: "Student Name to Update Details"
	//   required: true
	//   schema:
	//    type: string
	// requestBody:
	//  description: Updated Student Details
	//  content:
	//   application/json:
	//    schema:
	//     $ref: '#/definitions/schemas/Student'
	//
	// responses:
	//  200:
	//   description: "Student Updated Successfully"
	//   content:
	//    schema:
	//     $ref: '#/definitions/schemas/Student'
	//  400:
	//   description: "Invalid Student Name Specified"
	//   content:
	//    {}
	//  404:
	//   description: "Student Not Found"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check for method PUT, if anything else, respond with appropriate status
		if r.Method != "PUT" {
			res, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Bad Request : %v", err)
			}

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(res)
		}

		if r.Body != nil {
			var stuNew model.Student //student object to store updated student

			//extract name from path
			params := mux.Vars(r)
			nm := params["nm1"]

			defer r.Body.Close()

			//get a studentObject from GetByName using the extracted name
			stuToChange, err := dao.GetByName(nm, ds)
			if err != nil {
				log.Printf("Error While Fetching Record to Update : %v", err)
				res, _ := json.Marshal("Invalid Student Name specified")
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write(res) //send Message
				return

			}

			//Decode values from body sent from client into a studentObject
			err = json.NewDecoder(r.Body).Decode(&stuNew)
			if err != nil {
				log.Printf("Error While Deconding Body : %v", err)
			}

			//update the values from the body to The object got from the Database
			//stuToChange.StudentName = stuNew.StudentName
			stuToChange.StudentAge = stuNew.StudentAge     //update age
			stuToChange.StudentMarks = stuNew.StudentMarks //update marks

			//respond with appropriate message after calling Data Access Layer
			err = dao.UpdateStudent(stuToChange, ds)
			if err != nil {
				log.Printf("Could not Update student: %v", err)
			}

			res, err := json.Marshal("Updated Successfully")
			if err != nil {
				log.Printf("Error While Marshalling! : %v", err)
			}

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res) //send Message
		}

	})
}

//GetByName ...
func GetByName(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check for method GET, if anything else, respond with appropriate status
		if r.Method != "GET" {

			res, err := json.Marshal("Bad Request")

			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, err.Error(), 200)
			w.Write(res)
		}

		params := mux.Vars(r) //extract name from URL path

		var s model.Student
		s, err := dao.GetByName(params["nm"], ds) //call data access layer

		if err != nil {
			res, _ := json.Marshal("No Entry Found By That Name")
			w.Header().Set("Content Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(res)
			return
		}

		//respond with appropriate message
		mresult, _ := json.Marshal(s)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(mresult)
	})
}

//AddStudent ...
func AddStudent(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check if method is POST else show error
		if r.Method != "POST" {

			response, _ := json.Marshal("Bad Request")
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(response)
		}
		//check if body has content
		if r.Body != nil {
			defer r.Body.Close()
			var stu model.Student

			//decode the body for student details
			err := json.NewDecoder(r.Body).Decode(&stu)
			if err == nil {
				dao.AddStudent(stu, ds)
				//w.Header().Set("Access-Control-Allow-Methods","POST,OPTIONS")
				response, _ := json.Marshal("Added Successfully")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(response)
			}
		}
	})
}

//DeleteStudent ...
func DeleteStudent(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if method is DELETE else respond with error
		if r.Method != "DELETE" {
			response, err := json.Marshal("Bad Request")
			if err != nil {
				log.Printf("Bad Request!: %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(response)

		}

		//check if body has content
		if r.Body != nil {

			defer r.Body.Close()
			params := mux.Vars(r) //extract name of student from URL path

			//err := dao.GetByName(params["name"])
			//Respond to the requeset after calling Data Access Layer
			err := dao.RemoveByName(params["name"], ds)
			if err != nil {
				res, _ := json.Marshal("Could not Find anyone with that name")
				w.Header().Set("Content-Type", "appication/json")
				w.WriteHeader(http.StatusNotFound)
				w.Write(res)
				return
			}
			response, err := json.Marshal("Removed Student")
			if err != nil {
				log.Fatal(err)
				return
			}

			w.Header().Set("Content-Type", "appication/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}
	})
}

//GetAllStudents ...
func GetAllStudents(ds *mgo.Session) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check for method GET, if any other, respond with error with appropriate status
		if r.Method != "GET" {
			response, err := json.Marshal("Bad Request")

			if err != nil {
				log.Printf("Bad Request!: %v ", err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			//w.Header().S
			//w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(response)
		}

		//respond with appropriate message after calling Data Access Layer
		res, err := dao.GetAll(ds)
		if err != nil {
			log.Fatal(err)
		}
		response, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	})
}
