package gomiddlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponseWithDefaultErrorDescription(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := ErrorHandlerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(errors.New("error description"))
	}))
	handler.ServeHTTP(rr, req)

	assert.Contains(t, rr.Body.String(), "{\"status\":500,\"message\":\"Internal Server Error\"}\n")
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestErrorResponseWithCustomErrorDescription(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := ErrorHandlerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(ErrorWithStatus{Message: "status bad gateway", Status: http.StatusBadGateway})
	}))
	handler.ServeHTTP(rr, req)

	assert.Contains(t, rr.Body.String(), "{\"status\":502,\"message\":\"status bad gateway\"}\n")
	assert.Equal(t, http.StatusBadGateway, rr.Code)
}
