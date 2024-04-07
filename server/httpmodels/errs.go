package httpmodels

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s. Message: %s", e.Code, e.Message)
}

var IncorrectReqBodyErr = Error{
	Code:    "INCORRECT_REQUEST_BODY",
	Message: "description....",
}

var ValidationFailedErr = Error{
	Code:    "VALIDATION_FAILED",
	Message: "description....",
}

var IncorrectParameter = Error{
	Code:    "INCAORRECT_PARAMETER",
	Message: "description....",
}
