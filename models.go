package gomiddlewares

// errorResponse is a struct to represent JSON error responses.
type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ErrorWithStatus is an error type that includes a custom status code.
type ErrorWithStatus struct {
	Message string
	Status  int
}

// Error returns an error description.
func (e ErrorWithStatus) Error() string {
	return e.Message
}

// StatusCode returns an error status code.
func (e ErrorWithStatus) StatusCode() int {
	return e.Status
}
