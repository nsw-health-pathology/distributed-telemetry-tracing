package main

import (
	"fmt"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func main() {

	client := appinsights.NewTelemetryClient(getEnv("APPINSIGHTS_INSTRUMENTATIONKEY", ""))
	appInsightsService := AppInsightsService{Client: client}

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

	app := App{appInsights: appInsightsService}
	app.initialise(config, authController)
	app.run()
}
