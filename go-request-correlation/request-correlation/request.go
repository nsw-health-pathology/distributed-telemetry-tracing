package requestcorrelation

// Headers defines a dictionary of key value pairs (generally strings)
// which are passed on on an HTTP request
type Headers map[string]string

// HTTPMethod attempts to alias the HTTP Verb of the HTTP Request
// It should only use values from the permitted enum
type HTTPMethod string

const (
	// GET Request
	GET = "GET"
	// POST Request
	POST = "POST"
	// PUT Request
	PUT = "PUT"
	// PATCH Request
	PATCH = "PATCH"
	// HEAD Request
	HEAD = "HEAD"
	// OPTIONS Request
	OPTIONS = "OPTIONS"
)

// HTTPRequest defines the model for a general HTTP Request
type HTTPRequest struct {
	Headers Headers
	Method  HTTPMethod
	URL     string
}
