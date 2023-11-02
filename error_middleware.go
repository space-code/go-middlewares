package gomiddlewares

import (
	"encoding/json"
	"net/http"
)

// ErrorHandlerMiddleware is a middleware function that wraps an HTTP handler
// and handles panics that occur during the execution of the wrapped handler.
// It captures and recovers from panics, then sends an appropriate HTTP
// response with status code and error message in JSON format.
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				statusCode := http.StatusInternalServerError
				errorMessage := "Internal Server Error"

				if err, ok := r.(ErrorWithStatus); ok {
					errorMessage = err.Error()
					statusCode = err.StatusCode()
				}

				response := errorResponse{
					Status:  statusCode,
					Message: errorMessage,
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(statusCode)

				json.NewEncoder(w).Encode(response)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
