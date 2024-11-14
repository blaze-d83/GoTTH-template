package internal

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware logs the HTTP requests and responses with status and duration
func LoggingMiddleware(logger *logrus.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &StatusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		logRequest(logger, r)

		next.ServeHTTP(rec, r)

		logResponse(logger, rec, r, time.Since(start))
	})
}

// logRequest logs the incoming HTTP request details
func logRequest(logger *logrus.Logger, r *http.Request) {
	logger.WithFields(logrus.Fields{
		"method":  r.Method,
		"url":     r.URL.Path,
		"headers": r.Header,
	}).Info("HTTP REQUEST")
}

// logResponse logs the outgoing HTTP response details
func logResponse(logger *logrus.Logger, rec *StatusRecorder, r *http.Request, duration time.Duration) {
	logger.WithFields(logrus.Fields{
		"status":     rec.statusCode,
		"duration":   duration,
		"method":     r.Method,
		"url":        r.URL.Path,
		"user_agent": r.UserAgent(),
	}).Info("HTTP RESPONSE")
}

// StatusRecorder is a wrapper to capture HTTP status codes for response logging
type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before writing the response header
func (rec *StatusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
