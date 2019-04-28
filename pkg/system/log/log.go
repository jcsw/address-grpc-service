package log

import (
	golog "log"
	os "os"
)

func newLog() *golog.Logger {
	return golog.New(os.Stdout, "address-grpc-service ", golog.LstdFlags)
}

var syslog = newLog()

// Logger return the logger
func Logger() *golog.Logger {
	return syslog
}

// Debug - Logging in level DEBUG
func Debug(log string, v ...interface{}) {
	syslog.Printf("DEBUG "+log, v...)
}

// Info - Logging in level INFO
func Info(log string, v ...interface{}) {
	syslog.Printf("INFO  "+log, v...)
}

// Warn - Logging in level WARN
func Warn(log string, v ...interface{}) {
	syslog.Printf("WARN  "+log, v...)
}

// Error - Logging in level ERROR
func Error(log string, v ...interface{}) {
	syslog.Printf("ERROR "+log, v...)
}

// Fatal - Logging in level FATAL
func Fatal(log string, v ...interface{}) {
	syslog.Fatalf("FATAL  "+log, v...)
}
