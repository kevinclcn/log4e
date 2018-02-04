package log4e

import (
	"github.com/sirupsen/logrus"
	"io"
	"github.com/labstack/gommon/log"
)

type Logger4e struct {
	*logrus.Logger
}

func (l *Logger4e) Output() io.Writer {
	return l.Out
}

func (l *Logger4e) SetOutput(w io.Writer) {
	l.Out = w
}

func (l *Logger4e) Prefix() string {
	return ""
}

func (l *Logger4e) SetPrefix(p string) {
	// do nothing
}

func (l *Logger4e) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
	case logrus.FatalLevel:
		return log.ERROR
	}
	return log.OFF
}

func (l *Logger4e) SetLevel(v log.Lvl) {
	switch v {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	}
	l.Logger.SetLevel(logrus.InfoLevel)
}

func (l *Logger4e) Printj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Print()
}

func (l *Logger4e) Debugj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Debug()
}

func (l *Logger4e) Infoj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Info()
}

func (l *Logger4e) Warnj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Warn()
}

func (l *Logger4e) Errorj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Error()
}

func (l *Logger4e) Fatalj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Fatal()
}

func (l *Logger4e) Panicj(j log.JSON) {
	l.Logger.WithFields(logrus.Fields(j)).Panic()
}
