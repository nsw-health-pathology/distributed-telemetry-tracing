package requestcorrelation

// https://blog.golang.org/error-handling-and-go

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
