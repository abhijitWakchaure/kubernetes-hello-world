package main

import (
	"fmt"
	"log"
	"os"
)

var logger *Logger
var logLevel string

// Log levels
const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
)

// Logger is a wrapper over log provider
type Logger struct {
	l     *log.Logger
	level string
}

func initLogger() {
	logger = &Logger{}
	logger.l = log.New(os.Stdout, "example-app    ", log.Lmsgprefix)
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = INFO
	}
	logger.level = level
}

// Info ...
func (l *Logger) Info(msg string) {
	l.l.Println("[INFO ]", msg)
}

// Infof ...
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.l.Println(fmt.Sprintf("[INFO ] "+msg, args...))
}

// Debug ...
func (l *Logger) Debug(msg string) {
	if l.level == DEBUG {
		l.l.Println("[DEBUG]", msg)
	}
}

// Debugf ...
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.level == DEBUG {
		l.l.Println(fmt.Sprintf("[DEBUG] "+msg, args...))
	}
}

// Error ...
func (l *Logger) Error(msg string) {
	l.l.Println("[ERROR]", msg)
}

// Errorf ...
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.l.Println(fmt.Sprintf("[ERROR] "+msg, args...))
}
