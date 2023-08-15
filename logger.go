package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

var logLevelNames = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
}

type Logger struct {
	level  LogLevel
	logger *log.Logger
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

func (l *Logger) SetOutput(file string) error {
	absPath, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	logFile, err := os.OpenFile(absPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	l.logger.SetOutput(logFile)
	return nil
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level >= l.level {
		msg := fmt.Sprintf("[%s] %s", logLevelNames[level], fmt.Sprintf(format, args...))
		l.logger.Println(msg)
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(Debug, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(Warning, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}
