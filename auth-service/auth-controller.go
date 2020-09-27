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

	pUser, pErr := a.userSvc.getUser(username)

	if pErr != nil {
		fmt.Println(*pErr)
		return IHttpResponse{
			body:       *pErr,
			statusCode: 500,
		}
	}

	user := *pUser
	if user.Password != password {
		fmt.Println("Password Mismatch")
	}

	return IHttpResponse{
		body:       user,
		statusCode: 200,
	}

}
