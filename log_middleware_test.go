package gomiddlewares

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogFormat(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	req.Header.Set("User-Agent", "application/json")
	req.Header.Set("Referer", "testing")
	if err != nil {
		log.Fatal(err)
	}
	status := http.StatusOK
	conLen := 23
	var b bytes.Buffer
	SetOutput(&b)
	logRequest(status, conLen, req)
	rval := b.String()
	assert.Equal(t, rval, " GET  HTTP/1.1 200 23 application/json")
}

func TestShadowResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	l := logHandler{rr, http.StatusOK, 0}
	l.WriteHeader(http.StatusBadGateway)
	sval := "this is a body"
	l.Write([]byte(sval))
	assert.Equal(t, http.StatusBadGateway, l.statusCode, "should be equal")
	assert.Equal(t, len(sval), l.contentLen, "should be equal")
}

func TestErrorResponseWithLogMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)

	if err != nil {
		t.Fatal(err)
	}

	var b bytes.Buffer
	SetOutput(&b)

	rr := httptest.NewRecorder()
	handler := LogHandlerMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("empty response"))
	}))
	handler.ServeHTTP(rr, req)

	rval := b.String()
	assert.Equal(t, rval, " GET  HTTP/1.1 200 14 ")
}
