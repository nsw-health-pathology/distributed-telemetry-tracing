package main

// IHttpResponse models a http response message payload
type IHttpResponse struct {
	statusCode int
	headers    map[string]string
	body       interface{}
}

// IError models a generic error message
type IError struct {
	message string
}
