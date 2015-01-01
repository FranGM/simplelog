// Package simplelog provides a dead simple logging system.
package simplelog

import (
	"errors"
	"log"
	"os"
)

// Constants for the different log levels supported by the library
const (
	LevelDebug   int = iota // Threshold for Debug log level
	LevelInfo               // Threshold for Info log level
	LevelWarning            // Threshold for Warning log level
	LevelError              // Threshold for Error log level
)

// LogLevel represents a logger object for a given log level.
type LogLevel struct {
	logger *log.Logger
	prefix string
	level  int
}

// Logger objects that will be used to perform the actual logging.
// Each of them represents a different logging level and can be pointed to a different backend (file, stdout, etc...)
var (
	Error   = &LogLevel{prefix: "ERROR: ", level: LevelError}
	Warning = &LogLevel{prefix: "WARNING: ", level: LevelWarning}
	Info    = &LogLevel{prefix: "INFO: ", level: LevelInfo}
	Debug   = &LogLevel{prefix: "DEBUG: ", level: LevelDebug}
)
var logThreshold = LevelError

// Common errors that can be returned
var (
	ErrInvalidThreshold = errors.New("Invalid Threshold. Need one between LOG_DEBUG and LOG_ERROR") //
)

func init() {
	var levels = []*LogLevel{Error, Warning, Info, Debug}
	for _, level := range levels {
		// TODO: Allow for different handles than stdout
		level.logger = log.New(os.Stdout, level.prefix, log.LstdFlags)
	}
}

// SetThreshold sets the logging threshold level.
// Will return InvalidThreshold if new threshold isn't in the accepted range
func SetThreshold(t int) error {
	if t < LevelDebug || t > LevelError {
		return ErrInvalidThreshold
	}
	logThreshold = t
	return nil
}

// IsDebug will return true if logging threshold is currently set at Debug level
func IsDebug() bool {
	return logThreshold == LevelDebug
}

// LogThreshold will return the current log level
func LogThreshold() int {
	return logThreshold
}

// Printf will use the logger attached to this LogLevel to write a log message.
// Message will only get written if current log level allows it (it won't write INFO messages if we're at ERROR)
func (l *LogLevel) Printf(format string, v ...interface{}) {
	if l.level >= logThreshold {
		l.logger.Printf(format, v...)
	}
}
