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

	"github.com/blaze-d83/go-GoTTH/config"
	"github.com/labstack/echo/v4"
)

// LoggerStrategy defines an interface for logging strategies
type LoggerStrategy interface {
	LogRequests(ctx echo.Context)                                      // Logs an HTTP Request
	LogResponses(ctx echo.Context, status int, duration time.Duration) // Logs an HTTP Response
	LogError(ctx echo.Context, err error)                              // Logs an error that occured during an HTTP request
	LogEvent(message string, fields ...slog.Attr)                      // Logs a general event with custom attributes
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

// NewLogger initializes the new LoggerStrategy based on the configured logger type (sync or async)
func NewLogger(loggerType bool, handler slog.Handler) LoggerStrategy {
	logger := slog.New(handler) // Creates a new slogger instance

	// Returns the appropriate logger type
	if loggerType {
		return newAsyncLogger(logger, 100)
	}
	return newSyncLogger(logger)
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

func (l *SyncLogger) LogRequests(ctx echo.Context) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	l.logger.Info("HTTP request",
		slog.String("method", ctx.Request().Method),
		slog.String("path", ctx.Path()),
		slog.String("remote_addr", ctx.RealIP()),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogResponses(ctx echo.Context, status int, duration time.Duration) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	l.logger.Info("HTTP response",
		slog.String("method", ctx.Request().Method),
		slog.String("path", ctx.Path()),
		slog.Int("status", status),
		slog.Duration("duration", duration),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogError(ctx echo.Context, err error) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	l.logger.Error("ERROR",
		slog.String("method", ctx.Request().Method),
		slog.String("path", ctx.Path()),
		slog.Any("error", err),
		slog.String("request_id", requestID),
	)
}

func (l *SyncLogger) LogEvent(message string, fields ...slog.Attr) {
	args := make([]any, len(fields))
	for i, field := range fields {
		args[i] = field
	}
	l.logger.Info(message, args)
}

/*
 -- Asynchronous logging methods --
*/

func (l *AsyncLogger) LogRequests(ctx echo.Context) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	entry := LogEntry{
		message: "HTTP request",
		level:   slog.LevelInfo,
		fields: []slog.Attr{
			slog.String("method", ctx.Request().Method),
			slog.String("path", ctx.Path()),
			slog.String("remote_addr", ctx.RealIP()),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogResponses(ctx echo.Context, status int, duration time.Duration) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	entry := LogEntry{
		message: "HTTP response",
		level:   slog.LevelInfo,
		fields: []slog.Attr{
			slog.String("method", ctx.Request().Method),
			slog.String("path", ctx.Path()),
			slog.Int("status", status),
			slog.Duration("duration", duration),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogError(ctx echo.Context, err error) {
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	entry := LogEntry{
		message: "ERROR",
		level:   slog.LevelError,
		fields: []slog.Attr{
			slog.String("method", ctx.Request().Method),
			slog.String("path", ctx.Path()),
			slog.Any("error", err),
			slog.String("request_id", requestID),
		},
	}
	l.logChan <- entry
}

func (l *AsyncLogger) LogEvent(message string, fields ...slog.Attr) {
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

/*
	-- Local utilities --
*/

func setLogFormat(cfg config.LoggerConfig) slog.Handler {
	var handler slog.Handler
	level := setLogLevel(cfg)
	output, err := logOutput(cfg)
	if err != nil {
		return nil
	}
	if strings.ToLower(cfg.LogFormat) == "json" {
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{Level: level})
	} else {
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{Level: level})
	}
	return handler
}

func setLoggerType(cfg config.LoggerConfig) bool {
	if strings.ToLower(cfg.LogType) == "async" {
		return true
	}
	return false
}

func setLogLevel(cfg config.LoggerConfig) slog.Level {
	switch strings.ToLower(cfg.LogLevel) {
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

func logOutput(cfg config.LoggerConfig) (*os.File, error) {
	var output *os.File
	var err error

	if strings.ToLower(cfg.LogOutput) == "stdout" {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(cfg.LogOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("faile to open log file: " + err.Error())
		}
	}

	return output, err

}
