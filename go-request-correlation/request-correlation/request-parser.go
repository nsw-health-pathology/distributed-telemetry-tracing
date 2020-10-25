package requestcorrelation

// ParsedHTTPRequest defines a model for a parsed HTTP Request
type ParsedHTTPRequest struct {
	Method      HTTPMethod
	URL         string
	TraceParent TraceParent
}

const traceStateHeader = "tracestate"
const traceParentHeader = "traceparent"
const requestIDHeader = "request-id"

// ParseHTTPRequest parses the inbound HTTP Request and extracts the metadata
// required to build a traceable HTTP Request
func ParseHTTPRequest(req HTTPRequest, enableBackwardsCompatible bool) ParsedHTTPRequest {
	method := req.Method
	url := req.URL
	tp := processTraceHeaders(req, enableBackwardsCompatible)

	return ParsedHTTPRequest{
		Method:      method,
		URL:         url,
		TraceParent: *tp,
	}
}

func processTraceHeaders(req HTTPRequest, enableBackwardsCompatible bool) *TraceParent {
	var tp = extractTraceIDFromHTTPRequest(req)
	if tp == nil && enableBackwardsCompatible {
		tp = extractRequestIDFromHTTPRequest(req)
	}

	if tp == nil {
		newTp := GenerateNewTraceParent()
		tp = &newTp
	}

	return tp
}

func extractTraceIDFromHTTPRequest(req HTTPRequest) *TraceParent {
	if len(req.Headers) == 0 {
		return nil
	}

	tph, exists := req.Headers[traceParentHeader]
	if !exists {
		return nil
	}

	tp, err := TraceParentFromString(tph)
	if err != nil {
		return nil
	}

	newSpan := NewSpanFromParent(*tp)
	return &newSpan

}

func extractRequestIDFromHTTPRequest(req HTTPRequest) *TraceParent {
	// error processing TP from header. Attempt to parse from legacy Request-ID
	ridh, exists := req.Headers[requestIDHeader]
	if !exists {
		return nil
	}

	rid, err := RequestIDFromString(ridh)
	if err != nil {
		return nil
	}

	tp, _ := rid.ToTraceParent()

	newSpan := NewSpanFromParent(*tp)
	return &newSpan

}
