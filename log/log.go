package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var defaultLogger *logrus.Logger

func init() {
	defaultLogger = NewLogrusLogger()
}

func NewLogrusLogger() *logrus.Logger {
	l := logrus.New()

	l.Formatter = new(logrus.JSONFormatter)
	l.Out = os.Stderr
	l.Level = logrus.InfoLevel
	return l
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorf(foramt string, args ...interface{}) {
	defaultLogger.Errorf(foramt,args )
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Fatalf(foramt string, args ...interface{}) {
	defaultLogger.Fatalf(foramt,args )
}