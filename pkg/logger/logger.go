package logger

import (
    "log"
    "os"
)

// Logger represents a custom logger instance
type Logger struct {
    logger *log.Logger
}

// New creates a new Logger instance
func New() *Logger {
    return &Logger{
        logger: log.New(os.Stdout, "[E-COMMERCE] ", log.LstdFlags),
    }
}

// Info logs information messages
func (l *Logger) Info(message string) {
    l.logger.Printf("INFO: %s", message)
}

// Error logs error messages
func (l *Logger) Error(message string) {
    l.logger.Printf("ERROR: %s", message)
}

// Debug logs debug messages
func (l *Logger) Debug(message string) {
    l.logger.Printf("DEBUG: %s", message)
}