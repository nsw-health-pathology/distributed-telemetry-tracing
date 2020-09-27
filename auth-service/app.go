package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Following
// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// as an example

// AppConfiguration holds app configuration state to be passed into our live app
type AppConfiguration struct {
	userServiceBaseURL string
	jwtSigningKey      string
	port               string
}

// App will hold our application
type App struct {
	config         AppConfiguration
	authController IAuthController
	Router         *mux.Router
}

// Initialise the app with the incoming configuration properties
func (a *App) initialise(c AppConfiguration, authController IAuthController) {
	a.config = c
	a.authController = authController

	a.Router = mux.NewRouter()
	a.initialiseRoutes()
}

func (a *App) initialiseRoutes() {
	a.Router.HandleFunc("/login", a.handleLogin).Methods("POST")
}

// Starts the web api
func (a *App) run() {
	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":"+a.config.port, a.Router))
}

// Handler for /login request
func (a *App) handleLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling /login")

	// Read request
	var u UserLoginDetails
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, IError{Message: "Invalid request payload"})
		return
	}

	defer r.Body.Close()

	fmt.Println(u)

	response := a.authController.login(u.UserName, u.Password)
	respondWithJSON(w, response)
}
