package commons

import (
	"errors"
	"strings"

	v "gopkg.in/go-playground/validator.v9"
)

var validate *v.Validate

// SimpleStructValidator is a Utility to validate POST request bodies
// It takes in the Value POSTed, a validator function which contain rules,
// the actual model structure which contain validate tags
// Remember to tag your model with 'validate' before using this utility
// SimpleStructValidator Returs a string array of errors and a custom Error
// NOTE : a Struct level fuction has to be defined (fn), describing the rules to validate that Struct
func SimpleStructValidator(toBeValidated interface{}, types ...interface{}) (string, error) {
	// string array to catch the validation Errors
	var validationErrs []string
	// new instance of Validator
	validate = v.New()
	// // register the function against the struct to be validated
	// validate.RegisterStructValidation(fn, types)

	//validate struct passed in against the ValidationFunction
	err := validate.Struct(toBeValidated)
	if err != nil {
		for _, err := range err.(v.ValidationErrors) {
			validationErrs = append(validationErrs, err.StructField())
		}
		sErr := strings.Join(validationErrs, ",")
		//http.Error(w, "Validation Errors please check the supplied values for Test Status.\nBad Input Provided for - "+sErr, http.StatusUnprocessableEntity)
		return sErr, errors.New("Validation Error, throw a 400")
	}
	return "", nil
}
