package log4e

import (
	"github.com/sirupsen/logrus"
	"github.com/labstack/gommon/log"
	"io"
)


var logger = &Logger4e{
	Logger: logrus.New(),
}

func Logger() *Logger4e  {
	return logger
}

func Output() io.Writer {
	return logger.Out
}

func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

// SetLevel sets the standard logger level.
func SetLevel(level log.Lvl) {
	logger.SetLevel(level)
}

// GetLevel returns the standard logger level.
func Level() log.Lvl {
	return logger.Level()
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}

func Print(i ...interface{}) {
	logger.Print(i...)
}

func Printf(format string, i ...interface{}) {
	logger.Printf(format, i...)
}

func Printj(j log.JSON) {
	logger.Printj(j)
}

func Debug(i ...interface{}) {
	logger.Debug(i...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debugj(j log.JSON) {
	logger.Debugj(j)
}

func Info(i ...interface{}) {
	logger.Info(i...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Infoj(j log.JSON) {
	logger.Infoj(j)
}

func Warn(i ...interface{}) {
	logger.Warn(i...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warnj(j log.JSON) {
	logger.Warnj(j)
}

func Error(i ...interface{}) {
	logger.Error(i...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Errorj(j log.JSON) {
	logger.Errorj(j)
}

func Fatal(i ...interface{}) {
	logger.Fatal(i...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Fatalj(j log.JSON) {
	logger.Fatalj(j)
}

func Panic(i ...interface{}) {
	logger.Panic(i...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Panicj(j log.JSON) {
	logger.Panicj(j)
}
