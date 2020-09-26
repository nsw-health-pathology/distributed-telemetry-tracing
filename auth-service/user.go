package main

// User defines the model for a User Object in the User Service
type User struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserLoginDetails defines the model for a /login request
type UserLoginDetails struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
