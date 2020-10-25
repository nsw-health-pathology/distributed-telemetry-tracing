package requestcorrelation

import "strings"

// RequestID models the legacy Request-Id
type RequestID struct {
	TraceID string
	SpanID  string
}

// GenerateNewRequestID returns a new request-id
func GenerateNewRequestID() RequestID {
	rid := RequestID{
		TraceID: GenerateRandomHexString(TraceIDLength),
		SpanID:  GenerateRandomHexString(SpanIDLength),
	}
	return rid
}

// NewRequestID creates a new Request-Id struct
func NewRequestID(traceID string, spanID string) (*RequestID, error) {
	if err := validateRequestIDData(traceID, spanID); err != nil {
		return nil, err
	}

	return &RequestID{
		TraceID: traceID,
		SpanID:  spanID,
	}, nil
}

func validateRequestIDData(traceID string, spanID string) error {

	if err := IsValidTraceIDString(traceID); err != nil {
		return err
	}

	if err := IsValidSpanIDString(spanID); err != nil {
		return err
	}

	return nil
}

// ToString returns the legacy Request-Id header string
func (rid RequestID) ToString() string {
	return "|" + rid.TraceID + "." + rid.SpanID + "."
}

// ToTraceParent converts the legacy Request-Id to a W3C TraceParent
func (rid RequestID) ToTraceParent() (*TraceParent, error) {
	return NewTraceParent(
		0,
		rid.TraceID,
		rid.SpanID,
		0,
	)
}

// RequestIDFromString parses the legacy Request-ID value and converts it to
// a W3C RequestID model
func RequestIDFromString(requestID string) (*RequestID, error) {
	// Legacy Request-ID header is of the form
	// "|" + tp.TraceID + "." + tp.SpanID + "."

	// Extract TraceID
	// Extract SpanID
	traceID := extractTraceIDFromRequestID(requestID)
	spanID := extractSpanIDFromRequestID(requestID)

	return NewRequestID(
		traceID,
		spanID,
	)

}

func extractTraceIDFromRequestID(requestID string) string {
	// Legacy Request-ID header is of the form
	// "|" + tp.TraceID + "." + tp.SpanID + "."

	rootEnd := strings.Index(requestID, ".")
	if rootEnd < 0 {
		rootEnd = len(requestID)
	}

	// rootStart = requestID[0] == "|" ? 1 : 0;
	var rootStart = 0
	if string(requestID[0]) == "|" {
		rootStart = 1
	}

	return requestID[rootStart:rootEnd]
}

func extractSpanIDFromRequestID(requestID string) string {
	// Legacy Request-ID header is of the form
	// "|" + tp.TraceID + "." + tp.SpanID + "."
	substr := requestID[0 : len(requestID)-1]
	endOfRootIDIndex := strings.LastIndex(substr, ".")
	spanID := requestID[1+endOfRootIDIndex : len(requestID)-1]
	return spanID
}
