package middleware

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// WriteHeader captures the HTTP status code for logging
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write tracks response size for logging
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

// Hijack implements http.Hijacker for WebSocket support
func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := rw.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, fmt.Errorf("ResponseWriter does not implement http.Hijacker")
}

// Logging middleware logs request method, path, status code, response time and size
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		now := time.Now()
		// BR format: "02-01-2006 15:04:05 -07:00" | US format: "2006/01/02 15:04:05"
		timestamp := now.Format("2006/01/02 15:04:05")
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path = path + "?" + r.URL.RawQuery
		}

		logger.Printf("%s: %s %s %d %.3f ms - %d",
			timestamp,
			r.Method,
			path,
			wrapped.statusCode,
			float64(duration.Nanoseconds())/1e6,
			wrapped.size,
		)
	})
}
