package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &StatusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		logger.WithFields(logrus.Fields{
			"method":  r.Method,
			"url":     r.URL.Path,
			"headers": r.Header,
		}).Info("HTTP REQUEST")

		next.ServeHTTP(rec, r)

		logger.WithFields(logrus.Fields{
			"status":     rec.statusCode,
			"duration":   time.Since(start),
			"method":     r.Method,
			"url":        r.URL.Path,
			"user_agent": r.UserAgent(),
		}).Info("HTTP RESPONSE")

	})
}

type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *StatusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
