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
	"RESTApp/dao"
	"RESTApp/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	validator "gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
)

var validate *validator.Validate

//Handlers ...
func Handlers(ds *mgo.Session, trial string) http.Handler {
	server := mux.NewRouter() //create a new Server and attach handlers to it
	dir, _ := homedir.Dir()   // get the home directory of the user to host swagger ui
	server.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir(dir+"/go/src/InfoPortal/public"))))

	server.HandleFunc("/", redir).Methods("GET")
	server.HandleFunc("/swagger", GetSwagger).Methods("GET")
	server.HandleFunc("/plane", AddPlane(ds, trial)).Methods("POST")
	server.HandleFunc("/planes", GetPlanesHandler(ds, trial)).Methods("GET")
	server.HandleFunc("/plane/{name}", RemovePlaneByName(ds, trial)).Methods("DELETE")
	server.HandleFunc("/plane/{id}", RemovePlaneByID(ds, trial)).Methods("DELETE")
	server.HandleFunc("/studentAggregates", StudentAggregates(ds, trial)).Methods("GET")

	//Student Handlers
	server.HandleFunc("/student/{name}", DeleteStudent(ds, trial)).Methods("DELETE") //done
	server.HandleFunc("/students", GetAllStudents(ds, trial)).Methods("GET")         //done
	server.HandleFunc("/student/{name}", GetStudentByName(ds, trial)).Methods("GET") //done
	server.HandleFunc("/student/{name}", UpdateStud(ds, trial)).Methods("PUT")
	server.HandleFunc("/student", AddStudent(ds, trial)).Methods("POST") // done        //done

	//Book Handlers
	//server.HandleFunc("/book", book.GetBookSession).Methods("GET")

	return server
}

// GetSwagger ...
func GetSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/swagger.json")
}

func redir(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8081/public/dist/#/", http.StatusFound)
}

//GetPlanesHandler ...
func GetPlanesHandler(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation GET /planes GET getPlanes
	//
	// Get Planes
	//
	// Get Catalog of planes
	// ---
	// produces:
	// - application/json
	// responses:
	//  '200':
	//    description: Found Results
	//    schema:
	//     type: array
	//     items:
	//      "$ref": "#/definitions/GetPlanesAPIResponse"
	//  '401':
	//    description: Unauthorized, Likely Invalid or Missing Token
	//  '403':
	//    description: Forbidden, you are not allowed to undertake this operation
	//  '404':
	//    description: Not found
	//  '500':
	//    description: Error occurred while processing the request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Error Bad Request", http.StatusMethodNotAllowed)

		}
		allPlanes, err := dao.GetAllPlanes(ds, trial)
		if err != nil {
			http.Error(w, "Error while Fetching Planes ", http.StatusInternalServerError)
			return
		}

		res, _ := json.Marshal(allPlanes)
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
func AddPlane(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation POST /plane POST putPlane
	//
	//
	// Put Plane in the Plane catalog
	//
	//
	// Add a New Plane
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Plane
	//   type: object
	//   description: Student to be Added in the Catalog
	//   required: true
	//   in: body
	//   schema:
	//    $ref: '#/definitions/Plane'
	// responses:
	//  '200':
	//    description: Added Plane To the Catalog Successfully
	//    schema:
	//     $ref: '#/definitions/Plane'
	//  '401':
	//    description: Unauthorized, Likely Invalid or Missing Token
	//  '403':
	//    description: Forbidden, you are not allowed to undertake this operation
	//  '404':
	//    description: Not found
	//  '500':
	//    description: Error occurred while processing the request
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
			return
		}

		if r.Body != nil {
			var pl model.Plane

			err := json.NewDecoder(r.Body).Decode(&pl)
			if err != nil {
				log.Printf("Error while Decode body : %v", err)
			}
			val := validator.New()
			err = val.Struct(pl)
			// Validation
			//err = commons.SimpleStructValidator(pl, model.Plane{})
			if err != nil {
				//log.Printf("Valdation Error %v", err)
				//if _, ok := err.(*validator.InvalidValidationError); ok {
				http.Error(w, "Validating Errors Please check the Values supplied.\nBad Input for - "+err.Error(), http.StatusBadRequest)
				return
				//}
			}

			dao.PutPlane(pl, ds, trial)
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal("Added Plane Successfully")
			w.Write(res)
			defer r.Body.Close()
		}
	})
}

// RemovePlaneByName ...
func RemovePlaneByName(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation DELETE /plane/{name} DELETE removePlane
	//
	// Delete Plane
	//
	// Delete a Plane from Plane Catalog
	// ---
	// produces:
	// - application/json
	// - application/xml
	// parameters:
	// - name: name
	//   in: path
	//   required: true
	//   description: The name of the Plane to be removed
	// responses:
	//  '200':
	//    description: Plane Removed Successfully
	//  '401':
	//    description: Unauthorized, Likely Invalid or Missing Token
	//  '403':
	//    description: Forbidden, you are not allowed to undertake this operation
	//  '404':
	//    description: Not found
	//  '500':
	//    description: Error occurred while processing the request
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			res, _ := json.Marshal("Bad Request")
			sendErr(w, http.StatusMethodNotAllowed, res)
		}
		params := mux.Vars(r)
		del := params["name"]
		fmt.Println("Name to be Deleted:", del)
		ok := dao.DeletePlane(del, ds, trial)
		if !ok {
			res, _ := json.Marshal("Could Not Delete Server Error")
			sendErr(w, http.StatusInternalServerError, res)
		} else {
			res, _ := json.Marshal("Deleted Successfully")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	})
}

// RemovePlaneByID ...
func RemovePlaneByID(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation DELETE /plane/{id} DELETE deletePlane
	//
	// Delete Plane
	//
	// Delete A Plane From Plane Catalog
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id
	//   type: string
	//   description: ID of the Plane to Delete
	//   in: path
	//   required: true
	// responses:
	//  200:
	//   description: Removed Plane from the Catalog
	//  500:
	//   description: "Invalid Plane ID Specified"
	//  404:
	//   description: "Plane Not Found"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			res, _ := json.Marshal("Bad Request")
			sendErr(w, http.StatusMethodNotAllowed, res)
			return
		}

		countCheckSession := ds.Clone()
		defer countCheckSession.Close()
		count, err := countCheckSession.DB("trial").C("Student").Count()
		if err != nil {
			log.Println("Error while getting count from DB", err)
		}

		params := mux.Vars(r)
		idIns := params["id"]
		fmt.Println("IN PATH :", params)
		log.Println(params)
		id, err := strconv.Atoi(idIns)
		fmt.Println("IN INTEGER", id)
		log.Print(id)
		if err != nil {
			res, _ := json.Marshal("Server Error")
			sendErr(w, http.StatusInternalServerError, res)
			return
		}
		if id > count {
			res, _ := json.Marshal("ID with the given value Doesn't Exist")
			sendErr(w, http.StatusBadRequest, res)
			return
		}

		ok := dao.DeletePlaneByID(id, ds, trial)
		if !ok {
			res, _ := json.Marshal("Could Not Delete Server Error")
			sendErr(w, http.StatusInternalServerError, res)
			return
		}
		res, _ := json.Marshal("Deleted Successfully")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	})
}

//UpdateStud Update Student Info ByName
func UpdateStud(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation PUT /student/{name} UPDATE updateStudent
	//
	// Update a Students Information in The Student Catalog
	// ---
	// description: "Update Student Details"
	// summary: "Update Student Details in the Catalog"
	// parameters:
	// - name: name
	//   in: path
	//   description: "Student Name to Update Details"
	//   required: true
	//   type: string
	// - name : Student
	//   in: body
	//   required: true
	//   schema:
	//    $ref: '#/definitions/Student'
	// responses:
	//  '200':
	//   description: "Student Updated Successfully"
	//  '400':
	//   description: "Invalid Student Name Specified"
	//  '404':
	//   description: "Student Not Found"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check for method PUT, if anything else, respond with appropriate status
		if r.Method != "PUT" {
			res, _ := json.Marshal("Bad Request")
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(res)
		}

		if r.Body != nil {
			var stuNew model.Student //student object to store updated student

			//extract name from path
			params := mux.Vars(r)
			nm := params["name"]
			fmt.Println(nm)
			defer r.Body.Close()

			//get a studentObject from GetByName using the extracted name
			stuToChange, err := dao.GetByName(nm, ds, trial)
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

			//update the values from the body to The object got from the Database
			//stuToChange.StudentName = stuNew.StudentName
			stuToChange.StudentAge = stuNew.StudentAge     //update age
			stuToChange.StudentMarks = stuNew.StudentMarks //update marks

			//respond with appropriate message after calling Data Access Layer
			err = dao.UpdateStudent(stuToChange, ds, trial)
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

//GetStudentByName ...
func GetStudentByName(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation GET /student/{name} GET getStudent
	//
	// ---
	// description: Get Student Details by name
	// summary: "Get Student By Name"
	// parameters:
	// - name: name
	//   type: string
	//   description: Name of the Student t
	//   in: path
	//   required: true
	// responses:
	//  '200':
	//   description: Details Fetched Status Ok
	//   schema:
	//    type: object
	//    $ref: '#/definitions/Student'
	//    example:
	//     studentName: Eshan
	//     studentAge: 25
	//     studentMarks: 70
	//  '400':
	//     description: "No Entry Found By that Name"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check for method GET, if anything else, respond with appropriate status
		if r.Method != "GET" {

			res, _ := json.Marshal("Bad Request")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			//http.Error(w, err.Error(), 200)
			w.Write(res)
		}

		params := mux.Vars(r) //extract name from URL path
		var s model.Student
		s, err := dao.GetByName(params["name"], ds, trial) //call data access layer

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
func AddStudent(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation POST /student POST AddStudent
	//
	// Add Student
	//
	// Add a Student to the Student
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Student
	//   required: true
	//   in: body
	//   schema:
	//    "$ref": '#/definitions/Student'
	// responses:
	//  '200':
	//   description: Added Student Successfully to the Catalog
	//  '401':
	//   description: Unauthorized, Likely Invalid or Missing Token
	//  '403':
	//   description: Forbidden, you are not allowed to undertake this operation
	//  '404':
	//   description: Not found
	//  '411':
	//   description: Error while Decoding POSTED body
	//  '422':
	//   description: Unprocessable Entity in POST body
	//  '500':
	//   description: Error occurred while processing the request
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

			var errs []string
			validate = validator.New()
			// register validation for 'Student'
			// NOTE: only have to register a non-pointer type for 'User', validator
			// internally dereferences during it's type checks
			validate.RegisterStructValidation(validateStudentStruct, model.Student{})
			//decode the body for student details
			err := json.NewDecoder(r.Body).Decode(&stu)
			if err != nil {
				http.Error(w, "Error while Decoding Body", 400)
			}

			//Return validation errors
			//valErrs := []validator.FieldError{}
			err = validate.Struct(stu)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					errs = append(errs, err.StructField())

				}
				sErr := strings.Join(errs, ",")
				http.Error(w, "Validation Errors please check the supplied values for Test Status.\nBad Input Provided for - "+sErr, http.StatusUnprocessableEntity)
				return
			}
			dao.AddStudent(stu, ds, trial)
			response, _ := json.Marshal("Added Successfully")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}
	})
}

func validateStudentStruct(s validator.StructLevel) {

	stud := s.Current().Interface().(model.Student)

	if len(stud.StudentName) == 0 {
		s.ReportError(stud.StudentName, "StudentName", "studentName", "studentName", "")
	}

	if stud.StudentAge <= 0 || stud.StudentAge >= 110 {
		s.ReportError(stud.StudentAge, "StudentAge", "studentAge", "studentAge", "")
	}

	if stud.StudentMarks < 0 || stud.StudentMarks > 100 {
		s.ReportError(stud.StudentMarks, "StudentMarks", "studentMarks", "studentMarks", "")
	}
}

//DeleteStudent ...
func DeleteStudent(ds *mgo.Session, trial string) http.HandlerFunc {
	// swagger:operation DELETE /student/{name} DELETE deleteStudent
	//
	// Delete Stident
	//
	// Delete A Student From Student Catalog
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: name
	//   type: string
	//   description: Name of the Student to Delete
	//   in: path
	//   required: true
	// responses:
	//  200:
	//   description: Removed Student from the Catalog
	//  400:
	//   description: "Invalid Student Name Specified"
	//  404:
	//   description: "Student Not Found"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if method is DELETE else respond with error
		if r.Method != "DELETE" {
			http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
			return
		}

		//check if body has content
		//if r.Body != nil {

		//defer r.Body.Close()
		params := mux.Vars(r) //extract name of student from URL path

		//err := dao.GetByName(params["name"])
		//Respond to the requeset after calling Data Access Layer
		err := dao.RemoveByName(params["name"], ds, trial)
		if err != nil {
			res, _ := json.Marshal("Could not Find anyone with that name")
			w.Header().Set("Content-Type", "appication/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(res)
			return
		}
		response, _ := json.Marshal("Removed Student")
		w.Header().Set("Content-Type", "appication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		//}
	})
}

//GetAllStudents ...
func GetAllStudents(ds *mgo.Session, trial string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// swagger:operation GET /students GET getAllStudents
		//
		// List of Students
		//
		// Get the Student Catalog in Response
		//
		// ---
		// produces:
		// - application/json
		// responses:
		//  '200':
		//   description: Found Results
		//   schema:
		//    type: array
		//    items:
		//     "$ref": "#/definitions/GetAllStudentsAPIResponse"
		//  '400':
		//   description: "Invalid Student Name Specified"
		//  '404':
		//   description: "Student Not Found"

		//check for method GET, if any other, respond with error with appropriate status
		if r.Method != "GET" {
			http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
			return
		}

		//respond with appropriate message after calling Data Access Layer
		res, err := dao.GetAll(ds, trial)
		if err != nil {
			log.Fatal(err)
		}
		response, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	})
}
