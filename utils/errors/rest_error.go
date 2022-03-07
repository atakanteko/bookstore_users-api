package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	var customErr RestErr
	customErr = RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	return &customErr
}

func NewNotFoundError(message string) *RestErr {
	var customErr RestErr
	customErr = RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_request",
	}
	return &customErr
}
