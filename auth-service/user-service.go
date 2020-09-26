package main

// IUserService defines the interface that structs will need to implement in order
// to be considered compatible with communicating with the User Service
type IUserService interface {
	getUser(username string) (User, error)
}

// UserService is a concrete implementation of the IUserService
type UserService struct {
	userServiceBaseURL string
}

// Because the function signature matches the interface defined for IUserService,
// The UserService struct is considered to implement the interface
func (u UserService) getUser(userName string) (User, error) {
	return User{UserName: "mock"}, nil
}
