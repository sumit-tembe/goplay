package logger

import (
	"fmt"
	"log"
	"os"
)

// Log levels
const (
	TRACE   = "TRACE"
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

//LogLevelsMap represents the string names to integer mappings
var LogLevelsMap = map[string]int{
	"TRACE":   0,
	"DEBUG":   1,
	"INFO":    2,
	"WARNING": 3,
	"ERROR":   4,
}

var logLevel = 2

var (
	traceLogger = log.New(os.Stdout, "[TRACE]: ", log.Ldate|log.Ltime|log.Llongfile)
	debugLogger = log.New(os.Stdout, "[DEBUG]: ", log.Ldate|log.Ltime|log.Llongfile)
	infoLogger  = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime|log.Llongfile)
)

//Trace does trace level logging
func Trace(v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Tracef does trace level logging and formats the message
func Tracef(format string, v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Debug does debug level logging
func Debug(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Debugf does debug level logging
func Debugf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Info does info level logging
func Info(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		infoLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Infof does info level logging and formats the message
func Infof(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		infoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Warn does warn level logging
func Warn(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		warnLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Warnf does warn level logging and formats the message
func Warnf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		warnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Error does error level logging
func Error(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		errorLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Errorf does error level logging and formats the message
func Errorf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//SetLogLevel sets the log level to be used for logging
func SetLogLevel(loggingLevel string) {
	logLevel = LogLevelsMap[loggingLevel]
}
