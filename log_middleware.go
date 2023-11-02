package gomiddlewares

import (
	"fmt"
	"net/http"
)

type logHandler struct {
	http.ResponseWriter
	statusCode int
	contentLen int
}

func LogHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lh := &logHandler{w, http.StatusOK, 0}
		next.ServeHTTP(lh, r)
		logRequest(lh.statusCode, lh.contentLen, r)
	})
}

func logRequest(status, len int, r *http.Request) {
	fmt.Fprintf(output, "%s %s %s %s %d %d %s", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, status, len, r.UserAgent())
}

func (l *logHandler) WriteHeader(statusCode int) {
	l.statusCode = statusCode
	l.ResponseWriter.WriteHeader(statusCode)
}

func (l *logHandler) Write(b []byte) (int, error) {
	n, err := l.ResponseWriter.Write(b)
	l.contentLen += n
	return n, err
}
