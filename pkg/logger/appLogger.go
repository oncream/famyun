package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// appLogger is the implemention for logrus.
type appLogger struct {
	// name is the name of logger that is published to log as a scope
	name string
	// loger is the instance of logrus logger
	logger *logrus.Entry
}

var _ Logger = (*appLogger)(nil)

type Option func(logger *appLogger)

type New func(name string) Logger

func newAppLogger(name string, options ...Option) *appLogger {
	newLogger := logrus.New()
	newLogger.SetOutput(os.Stdout)
	dl := &appLogger{
		name: name,
		logger: newLogger.WithFields(logrus.Fields{
			logFieldScope: name,
		}),
	}
	dl.EnableJSONOutput(true)
	for _, option := range options {
		option(dl)
	}
	return dl
}

func Level(outputLevel LogLevel) Option {
	return func(l *appLogger) {
		l.logger.Logger.SetLevel(toLogrusLevel(outputLevel))
	}
}

func AddHook(hook logrus.Hook) Option {
	return func(l *appLogger) {
		l.logger.Logger.AddHook(hook)
	}
}

// EnableJSONOutput enables JSON formatted output log.
func (l *appLogger) EnableJSONOutput(enabled bool) {
	var formatter logrus.Formatter

	fieldMap := logrus.FieldMap{
		// If time field name is conflicted, logrus adds "fields." prefix.
		// So rename to unused field @time to avoid the confliction.
		logrus.FieldKeyTime:  logFieldTimeStamp,
		logrus.FieldKeyLevel: logFieldLevel,
		logrus.FieldKeyMsg:   logFieldMessage,
	}

	hostname, _ := os.Hostname()
	l.logger.Data = logrus.Fields{
		logFieldScope:    l.logger.Data[logFieldScope],
		logFieldInstance: hostname,
	}

	if enabled {
		formatter = &logrus.JSONFormatter{ //nolint: exhaustruct
			TimestampFormat: time.RFC3339Nano,
			FieldMap:        fieldMap,
		}
	} else {
		formatter = &logrus.TextFormatter{ //nolint: exhaustruct
			TimestampFormat: time.RFC3339Nano,
			FieldMap:        fieldMap,
		}
	}

	l.logger.Logger.SetFormatter(formatter)
}

func toLogrusLevel(lvl LogLevel) logrus.Level {
	// ignore error because it will never happen
	l, _ := logrus.ParseLevel(string(lvl))
	return l
}

// SetOutputLevel sets log output level.
func (l *appLogger) SetOutputLevel(outputLevel LogLevel) {
	l.logger.Logger.SetLevel(toLogrusLevel(outputLevel))
}

// IsOutputLevelEnabled returns true if the logger will output this LogLevel.
func (l *appLogger) IsOutputLevelEnabled(level LogLevel) bool {
	return l.logger.Logger.IsLevelEnabled(toLogrusLevel(level))
}

// SetOutput sets the destination for the logs.
func (l *appLogger) SetOutput(dst io.Writer) {
	l.logger.Logger.SetOutput(dst)
}

// WithFields returns a logger with the added structured fields.
func (l *appLogger) WithFields(fields map[string]any) Logger {
	return &appLogger{
		name:   l.name,
		logger: l.logger.WithFields(fields),
	}
}

func (l *appLogger) GetLevel() logrus.Level {
	return l.logger.Logger.GetLevel()
}

// Info logs a message at level Info.
func (l *appLogger) Info(args ...interface{}) {
	l.logger.Log(logrus.InfoLevel, args...)
}

// Infof logs a message at level Info.
func (l *appLogger) Infof(format string, args ...interface{}) {
	l.logger.Logf(logrus.InfoLevel, format, args...)
}

// Debug logs a message at level Debug.
func (l *appLogger) Debug(args ...interface{}) {
	l.logger.Log(logrus.DebugLevel, args...)
}

// Debugf logs a message at level Debug.
func (l *appLogger) Debugf(format string, args ...interface{}) {
	l.logger.Logf(logrus.DebugLevel, format, args...)
}

// Warn logs a message at level Warn.
func (l *appLogger) Warn(args ...interface{}) {
	l.logger.Log(logrus.WarnLevel, args...)
}

// Warnf logs a message at level Warn.
func (l *appLogger) Warnf(format string, args ...interface{}) {
	l.logger.Logf(logrus.WarnLevel, format, args...)
}

// Error logs a message at level Error.
func (l *appLogger) Error(args ...interface{}) {
	l.logger.Log(logrus.ErrorLevel, args...)
}

// Errorf logs a message at level Error.
func (l *appLogger) Errorf(format string, args ...interface{}) {
	l.logger.Logf(logrus.ErrorLevel, format, args...)
}

// Fatal logs a message at level Fatal then the process will exit with status set to 1.
func (l *appLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalf logs a message at level Fatal then the process will exit with status set to 1.
func (l *appLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *appLogger) AddHook(hook logrus.Hook) {
	l.logger.Logger.AddHook(hook)
}
