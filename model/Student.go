package model

// Student represents students info as a record in this application
// swagger:model Student
type Student struct {
	// Name for the Student
	// required: true
	// min length: 3
	StudentName string `json:"studentName" bson:"studentName" validate:"required"`
	// Age for the Student
	// required: true
	// min: 1
	StudentAge int `json:"studentAge,string" bson:"studentAge" validate:"required"`
	// Marks for the Student
	// required: true
	// min: 1
	StudentMarks int `json:"studentMarks,string" bson:"studentMarks" validate:"required"`
}

// GetAllStudentsAPIResponse holds list of Students
// swagger:model GetAllStudentsAPIResponse
type GetAllStudentsAPIResponse struct {
	Student []Student `json:"Students"`
}
