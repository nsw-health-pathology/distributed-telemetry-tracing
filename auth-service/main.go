package main

import "os"

func main() {

	config := AppConfiguration{
		userServiceBaseURL: os.Getenv("USER_SERVICE_BASE_URL"),
		jwtSigningKey:      os.Getenv("JWT_SIGNING_KEY"),
		port:               os.Getenv("PORT"),
	}

	svc := UserService{}
	authController := AuthController{userSvc: svc}

	app := App{}
	app.initialise(config, authController)
	app.run()
}
