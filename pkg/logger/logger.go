// Package logger provides a simple logging interface and implementation.
package logger

import (
	"log"
	"os"
)

// Logger defines the interface for logging operations.
type Logger interface {
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Fatal(msg string, keyvals ...interface{})
}

// simpleLogger implements the Logger interface using the standard log package.
type simpleLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewSimpleLogger creates and returns a new instance of simpleLogger.
func NewSimpleLogger() Logger {
	return &simpleLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message.
func (l *simpleLogger) Info(msg string, keyvals ...interface{}) {
	l.infoLogger.Println(msg, keyvals)
}

// Error logs an error message.
func (l *simpleLogger) Error(msg string, keyvals ...interface{}) {
	l.errorLogger.Println(msg, keyvals)
}

// Fatal logs a fatal error message and then calls os.Exit(1).
func (l *simpleLogger) Fatal(msg string, keyvals ...interface{}) {
	l.errorLogger.Fatal(msg, keyvals)
}