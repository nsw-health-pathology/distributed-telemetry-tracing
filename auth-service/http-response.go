package main

import "fmt"

// IHttpResponse models a http response message payload
type IHttpResponse struct {
	statusCode int
	headers    map[string]string
	body       interface{}
}

// IError models a generic error message
type IError struct {
	Message string `json:"message"`
}

// Need to implement the Error() method for IError
// to be considered to implement the error interface
func (e IError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
