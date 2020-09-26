package main

import (
	"fmt"
)

// IAuthController defines the controller interface
type IAuthController interface {
	login(username string, password string) IHttpResponse
}

// AuthController is the concrete implementation for IAuthController
type AuthController struct {
	userSvc IUserService
}

func (a AuthController) login(username string, password string) IHttpResponse {

	user, err := a.userSvc.getUser(username)

	if err != nil {
		fmt.Println(err)
	}

	if user.Password != password {
		fmt.Println("Password Mismatch")
	}

	return IHttpResponse{
		body:       user,
		statusCode: 200,
	}

}
