/*
-- Logger Package --

This package provides a simple logging interface with support for both synchronous and asynchronous logging.

Uses slog package for structure logging in various formats.
*/
package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/blaze-d83/go-GoTTH/pkg/config"
)

// LoggerStrategy defines an interface for logging strategies
type Logger interface {
	LogRequests(ctx context.Context, method, path, remoteAddr, requestID string)                          // Logs an HTTP Request
	LogResponses(ctx context.Context, status int, duration time.Duration, method, path, requestID string) // Logs an HTTP Response
	LogError(ctx context.Context, err error, method, path, requestID string)                              // Logs an error that occured during an HTTP request
	LogEvent(ctx context.Context, message string, fields ...slog.Attr)                                    // Logs a general event with custom attributes
}

// SyncLogger provides a logger that writes log entries synchronously
type SyncLogger struct {
	// logger is the underlying slog.Logger instance for logging
	logger *slog.Logger
}

// AsyncLogger provides a logger that writes log entries asynchronously
type AsyncLogger struct {
	// logger is the underlying slog.Logger instance for logging
	logger *slog.Logger
	// logChan is a buffered channel that holds the log entries for async processing
	logChan chan LogEntry
	// wg initializes sync.WaitGroup to track collection of goroutines
	wg sync.WaitGroup
	// stopChan is used to signal the logger to stop processing
	stopChan chan struct{}
}

// LogEntry represents a log entry consisting of a message, log level, and custom attributes
type LogEntry struct {
	message string      // Log message
	level   slog.Level  // Log level (eg. Info, Error)
	fields  []slog.Attr // Additional attributes associated with log entry
}

func InitializeLogger(cfg config.LoggerConfig) (Logger, error) {
	output, err := logOutput(cfg.LogOutput)
	if err != nil {
		return nil, err
	}

	handler := createHandler(cfg, output)

	logger := slog.New(handler)

	if strings.ToLower(cfg.LogType) == "async" {
		return newAsyncLogger(logger, 100), nil
	}
	return newSyncLogger(logger), nil
}

func createHandler(cfg config.LoggerConfig, output *os.File) slog.Handler {
	level := parseLogLevel(cfg.LogLevel)

	if strings.ToLower(cfg.LogFormat) == "json" {
		return slog.NewJSONHandler(output, &slog.HandlerOptions{Level: level})
	}
	return slog.NewTextHandler(output, &slog.HandlerOptions{Level: level})
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func logOutput(output string) (*os.File, error) {
	if strings.ToLower(output) == "stdout" {
		return os.Stdout, nil
	}
	return os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// newSyncLogger creates a new instance of SyncLogger
func newSyncLogger(logger *slog.Logger) *SyncLogger {
	return &SyncLogger{
		logger: logger,
	}
}

// newAsyncLogger creates a new instance of AsyncLogger
func newAsyncLogger(logger *slog.Logger, buffersize int) *AsyncLogger {
	asl := &AsyncLogger{
		logger:   logger,
		logChan:  make(chan LogEntry, buffersize),
		stopChan: make(chan struct{}),
	}

	asl.wg.Add(1)
	go asl.processLogs()
	return asl

}

/*
 -- Synchronous logging methods --
*/

func (l *SyncLogger) LogRequests(ctx context.Context, method, path, remoteAddr, requestID string) {
	l.logger.Info("HTTP request",
		slog.String("method", method),
		slog.String("path", path),
		slog.String("remote_addr", remoteAddr),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogResponses(ctx context.Context, status int, duration time.Duration, method, path, requestID string) {
	l.logger.Info("HTTP response",
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status", status),
		slog.Duration("duration", duration),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogError(ctx context.Context, err error, method, path, requestID string) {
	l.logger.Error("ERROR",
		slog.String("method", method),
		slog.String("path", path),
		slog.Any("error", err),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogEvent(ctx context.Context, message string, fields ...slog.Attr) {
	args := make([]any, len(fields)*2)
	for i, field := range fields {
		args[i*2] = field.Key
		args[i*2+1] = field.Value
	}
	l.logger.Info(message, args...)
}

/*
 -- Asynchronous logging methods --
*/

func (l *AsyncLogger) LogRequests(ctx context.Context, method, path, remote_addr, requestID string) {
	entry := LogEntry{
		message: "HTTP request",
		level:   slog.LevelInfo,
		fields: []slog.Attr{
			slog.String("method", method),
			slog.String("path", path),
			slog.String("remote_addr", remote_addr),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogResponses(ctx context.Context, status int, duration time.Duration, method, path, requestID string) {
	entry := LogEntry{
		message: "HTTP response",
		level:   slog.LevelInfo,
		fields: []slog.Attr{
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", status),
			slog.Duration("duration", duration),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogError(ctx context.Context, err error, method, path, requestID string) {
	entry := LogEntry{
		message: "ERROR",
		level:   slog.LevelError,
		fields: []slog.Attr{
			slog.String("method", method),
			slog.String("path", path),
			slog.Any("error", err),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogEvent(ctx context.Context, message string, fields ...slog.Attr) {
	entry := LogEntry{
		message: message,
		level:   slog.LevelInfo,
		fields:  fields,
	}
	l.logChan <- entry
}

func (l *AsyncLogger) processLogs() {
	defer l.wg.Done()
	for {
		select {
		case entry := <-l.logChan:
			l.logger.LogAttrs(context.Background(), entry.level, entry.message, entry.fields...)
		case <-l.stopChan:
			for len(l.logChan) > 0 {
				entry := <-l.logChan
				l.logger.LogAttrs(context.Background(), entry.level, entry.message, entry.fields...)
			}
			return
		}
	}

}

func (l *AsyncLogger) StopLogger() {
	close(l.stopChan)
	l.wg.Wait()

}
