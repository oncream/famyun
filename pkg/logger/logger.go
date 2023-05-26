package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

const (
	// Field names that defines Dapr log schema.
	logFieldTimeStamp = "time"
	logFieldLevel     = "level"
	logFieldScope     = "scope"
	logFieldMessage   = "msg"
	logFieldInstance  = "instance"
)

// LogLevel is App Logger Level type.
type LogLevel string

const (
	// DebugLevel has verbose message.
	DebugLevel LogLevel = "debug"
	// InfoLevel is default log level.
	InfoLevel LogLevel = "info"
	// WarnLevel is for logging messages about possible issues.
	WarnLevel LogLevel = "warn"
	// ErrorLevel is for logging errors.
	ErrorLevel LogLevel = "error"
	// FatalLevel is for logging fatal messages. The system shuts down after logging the message.
	FatalLevel LogLevel = "fatal"

	// UndefinedLevel is for undefined log level.
	UndefinedLevel LogLevel = "undefined"
)

var (
	globalLoggers     = map[string]Logger{}
	globalLoggersLock = sync.RWMutex{}
)

// Logger includes the logging api sets.
type Logger interface { //nolint: interfacebloat
	// EnableJSONOutput enables JSON formatted output log
	EnableJSONOutput(enabled bool)

	// SetOutputLevel sets the log output level
	SetOutputLevel(outputLevel LogLevel)
	// SetOutput sets the destination for the logs
	SetOutput(dst io.Writer)

	// IsOutputLevelEnabled returns true if the logger will output this LogLevel.
	IsOutputLevelEnabled(level LogLevel) bool

	// WithFields returns a logger with the added structured fields.
	WithFields(fields map[string]any) Logger

	// Info logs a message at level Info.
	Info(args ...interface{})
	// Infof logs a message at level Info.
	Infof(format string, args ...interface{})
	// Debug logs a message at level Debug.
	Debug(args ...interface{})
	// Debugf logs a message at level Debug.
	Debugf(format string, args ...interface{})
	// Warn logs a message at level Warn.
	Warn(args ...interface{})
	// Warnf logs a message at level Warn.
	Warnf(format string, args ...interface{})
	// Error logs a message at level Error.
	Error(args ...interface{})
	// Errorf logs a message at level Error.
	Errorf(format string, args ...interface{})
	// Fatal logs a message at level Fatal then the process will exit with status set to 1.
	Fatal(args ...interface{})
	// Fatalf logs a message at level Fatal then the process will exit with status set to 1.
	Fatalf(format string, args ...interface{})
	// AddHook add hook
	AddHook(hook logrus.Hook)
	// GetLevel get log level
	GetLevel() logrus.Level
}

func NewLogger(name string, options ...Option) Logger {
	globalLoggersLock.Lock()
	defer globalLoggersLock.Unlock()

	logger, ok := globalLoggers[name]
	if !ok {
		logger = newAppLogger(name)
		globalLoggers[name] = logger
	}
	return logger

}
