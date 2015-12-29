// log provides a small set of functions for extended logging.
// Log messages have different levels (Debug, Info, Warning, Error),
// are colorized per level and are prepended with an optional timestamp.
// Only messages greater or equal to the set log level will be logged to std out/error.
package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mitchellh/colorstring"
)

const (
	// Timestamp format.
	tsf = "2006/01/02 15:04:05"
)

// Level holds the current log level.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levelPrefix = [...]string{
	LevelDebug: "DEBUG: ",
	LevelInfo:  "INFO: ",
	LevelWarn:  "WARNING: ",
	LevelError: "ERROR: ",
}

var levelColor = [...]string{
	LevelDebug: "[blue]",
	LevelInfo:  "[cyan]",
	LevelWarn:  "[yellow]",
	LevelError: "[red]",
}

var level = LevelDebug

// SetLevel changes the current log level to l.
func SetLevel(l Level) {
	level = l
}

// DisableColor, if true, disables colorized output.
var DisableColor = false

// DisableTime, if true, disables prepended timestamp.
var DisableTime = false

// Debug logs a debug message to stdout.
func Debug(v ...interface{}) error {
	return output(os.Stdout, LevelDebug, v...)
}

// Debugf logs a formatted debug message to stdout.
func Debugf(format string, v ...interface{}) error {
	return outputf(os.Stdout, LevelDebug, format, v...)
}

// Info logs an informational message to stdout.
func Info(v ...interface{}) error {
	return output(os.Stdout, LevelInfo, v...)
}

// Infof logs a formatted informational message to stdout.
func Infof(format string, v ...interface{}) error {
	return outputf(os.Stdout, LevelInfo, format, v...)
}

// Warn logs a warning message to stdout.
func Warn(v ...interface{}) error {
	return output(os.Stdout, LevelWarn, v...)
}

// Warnf logs a formatted warning message to stdout.
func Warnf(format string, v ...interface{}) error {
	return outputf(os.Stdout, LevelWarn, format, v...)
}

// Error logs an error message to stderr.
func Error(v ...interface{}) error {
	return output(os.Stderr, LevelError, v...)
}

// Errorf logs a formatted error message to stderr.
func Errorf(format string, v ...interface{}) error {
	return outputf(os.Stderr, LevelError, format, v...)
}

// colorize colorizes the output.
func colorize(l Level, s string) string {
	if DisableColor {
		return s
	}
	return colorstring.Color(levelColor[l] + s)
}

// timestamp prepends a timestamp.
func timestamp(s string) string {
	if DisableTime {
		return s
	}
	t := time.Now().Format(tsf)
	return t + " - " + s
}

// output writes the log message controlled by log level.
func output(w io.Writer, l Level, v ...interface{}) error {
	if l < level {
		return nil
	}
	//t := time.Now().Format(timeFormat)
	_, err := fmt.Fprint(w, colorize(l, timestamp(levelPrefix[l]+fmt.Sprintln(v...))))
	return err
}

// outputf writes the formatted log message controlled by log level.
func outputf(w io.Writer, l Level, format string, v ...interface{}) error {
	if l < level {
		return nil
	}
	_, err := fmt.Fprintf(w, colorize(l, timestamp(levelPrefix[l]+fmt.Sprintln(format))), v...)
	return err
}
