package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// IUserService defines the interface that structs will need to implement in order
// to be considered compatible with communicating with the User Service
type IUserService interface {
	getUser(username string) (*User, *IError)
}

// UserService is a concrete implementation of the IUserService
type UserService struct {
	userServiceBaseURL string
}

// Because the function signature matches the interface defined for IUserService,
// The UserService struct is considered to implement the interface
func (u UserService) getUser(userName string) (*User, *IError) {

	url := u.userServiceBaseURL + "/user"
	queryParams := map[string]string{
		"username": userName,
	}

	pResponse, err := userServiceGetUser(url, queryParams)

	if err != nil {
		e := IError{Message: err.Error()}
		return nil, &e
	}

	httpResponseBody := (*pResponse).body

	if (*pResponse).statusCode == 200 {
		var u User
		mapstructure.Decode(httpResponseBody, &u)
		fmt.Println("User Object", u)
		return &u, nil
	}

	var e IError
	mapstructure.Decode(httpResponseBody, &e)
	fmt.Println("Error Object", e)
	return nil, &e

}

func userServiceGetUser(url string, queryParams map[string]string) (*IHttpResponse, error) {

	client := &http.Client{}

	// Make initial request object

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create GET Request")
		return nil, err
	}

	// Add query params
	q := req.URL.Query()
	for qpName, qpValue := range queryParams {
		q.Add(qpName, qpValue)
	}
	req.URL.RawQuery = q.Encode()

	fmt.Println("Calling URL", url, "with params", q)

	// Execute Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failed to execute GET Request")
		return nil, err
	}

	// Read Response to User struct
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	if err != nil {
		fmt.Println("Failed to read response from GET Request")
		return nil, err
	}

	var responseObject interface{}
	decoder.Decode(&responseObject)
	fmt.Println("Response Object", responseObject)

	httpResponse := IHttpResponse{
		body:       responseObject,
		statusCode: resp.StatusCode,
	}

	return &httpResponse, nil

}
