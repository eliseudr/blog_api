package middleware

import (
	"log"
	"net/http"

	"github.com/eliseudr/blog_api/response"
)

type errorHandler struct {
	handler http.Handler
}

type responseTracker struct {
	http.ResponseWriter
	written    bool
	statusCode int
}

// WriteHeader tracks when a response header is written
func (rt *responseTracker) WriteHeader(code int) {
	rt.statusCode = code
	if code == http.StatusNotFound {
		rt.written = false
		return
	}
	rt.written = true
	rt.ResponseWriter.WriteHeader(code)
}

// Write tracks when response body is written
func (rt *responseTracker) Write(b []byte) (int, error) {
	if rt.statusCode == http.StatusNotFound {
		return len(b), nil
	}
	if !rt.written {
		rt.statusCode = http.StatusOK
		rt.written = true
	}
	return rt.ResponseWriter.Write(b)
}

// ServeHTTP handles HTTP requests and catches panics or 404 errors
func (eh *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tracker := &responseTracker{ResponseWriter: w, written: false, statusCode: 0}

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic recovered: %v", err)
			if !tracker.written {
				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
		} else if tracker.statusCode == http.StatusNotFound || !tracker.written {
			response.Error(w, http.StatusNotFound, "Route not found")
		}
	}()

	eh.handler.ServeHTTP(tracker, r)
}

// ErrorHandler wraps a handler to catch panics and handle 404 errors
func ErrorHandler(handler http.Handler) http.Handler {
	return &errorHandler{handler: handler}
}
