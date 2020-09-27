package main

import (
	"fmt"
	"os"
)

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func main() {

	config := AppConfiguration{
		userServiceBaseURL: getEnv("USER_SERVICE_BASE_URL", "http://localhost:8000"),
		jwtSigningKey:      getEnv("JWT_SIGNING_KEY", "MockSigningKey"),
		port:               getEnv("PORT", "8010"),
	}

	fmt.Println(config)

	svc := UserService{
		userServiceBaseURL: config.userServiceBaseURL,
	}
	authController := AuthController{userSvc: svc}

	app := App{}
	app.initialise(config, authController)
	app.run()
}
