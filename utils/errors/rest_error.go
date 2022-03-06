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
		Error:   "bad request",
	}
	return &customErr
}
