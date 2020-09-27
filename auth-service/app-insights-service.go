package main

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

// AppInsightsService wraps the App Insights Telemetry Client
type AppInsightsService struct {
	Client appinsights.TelemetryClient
}
