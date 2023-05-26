package logger

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm/log"
)

type XormLogger struct {
	log     Logger
	showSQL bool
}

func NewXormLogger(log Logger) log.Logger {
	return &XormLogger{
		log:     log,
		showSQL: false,
	}
}

var _ log.Logger = (*XormLogger)(nil)

func (x *XormLogger) Debug(v ...interface{}) {
	x.log.Debug(v...)
}

func (x *XormLogger) Debugf(format string, v ...interface{}) {
	x.log.Debugf(format, v...)
}

func (x *XormLogger) Error(v ...interface{}) {
	x.log.Error(v...)
}

func (x *XormLogger) Errorf(format string, v ...interface{}) {
	x.log.Errorf(format, v...)
}

func (x *XormLogger) Info(v ...interface{}) {
	x.log.Info(v...)
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	x.log.Infof(format, v...)
}

func (x *XormLogger) Warn(v ...interface{}) {
	x.log.Warn(v...)
}

func (x *XormLogger) Warnf(format string, v ...interface{}) {
	x.log.Warnf(format, v...)
}

func (x *XormLogger) Level() log.LogLevel {
	return toXormLogLevel(x.log.GetLevel())
}

func (x *XormLogger) SetLevel(l log.LogLevel) {
	x.log.SetOutputLevel(fromXormLogLevel(l))
}

func (x *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		x.showSQL = true
		return
	}
	x.showSQL = show[0]
}

func (x *XormLogger) IsShowSQL() bool {
	return x.showSQL
}

func toXormLogLevel(level logrus.Level) log.LogLevel {
	switch level {
	case logrus.InfoLevel:
		return log.LOG_INFO
	case logrus.DebugLevel:
		return log.LOG_DEBUG
	case logrus.ErrorLevel:
		return log.LOG_ERR
	case logrus.WarnLevel:
		return log.LOG_WARNING
	default:
		return log.LOG_UNKNOWN
	}
}

func fromXormLogLevel(level log.LogLevel) LogLevel {
	switch level {
	case log.LOG_INFO:
		return InfoLevel
	case log.LOG_DEBUG:
		return DebugLevel
	case log.LOG_ERR:
		return ErrorLevel
	case log.LOG_WARNING:
		return WarnLevel
	default:
		return UndefinedLevel
	}
}
