package model

// Student represents students info as a record in this application
// swagger:model Student
type Student struct {
	// Name for the Student
	// required: true
	// min length: 3
	StudentName string `json:"studentName" bson:"studentName"`
	// Age for the Student
	// required: true
	// min: 1
	StudentAge int `json:"studentAge,string" bson:"studentAge"`
	// Marks for the Student
	// required: true
	// min: 1
	StudentMarks int `json:"studentMarks,string" bson:"studentMarks"`
}
