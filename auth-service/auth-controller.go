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
		return IHttpResponse{
			body:       IError{Message: "Invalid login attempt"},
			statusCode: 401,
		}
	}

	token := a.userSvc.makeTokenForUser(pUser)
	fmt.Println("Token String", token)
	return IHttpResponse{
		body:       user,
		statusCode: 200,
		headers: map[string]string{
			"Authorization": token,
		},
	}

}
