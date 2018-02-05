package log4e

import (
	"github.com/sirupsen/logrus"
	"github.com/labstack/gommon/log"
	"io"
)

type Entry4e struct {
	*logrus.Entry
}

func (entry *Entry4e) Logger() *Logger4e {
	return &Logger4e{entry.Entry.Logger}
}

func (entry *Entry4e) Output() io.Writer {
	return entry.Logger().Output()
}

func (entry *Entry4e) SetOutput(w io.Writer) {

}

func (entry *Entry4e) Prefix() string {
	return ""
}

func (entry *Entry4e) SetPrefix(p string) {
}

func (entry *Entry4e) Level() log.Lvl {
	return entry.Logger().Level()
}

func (entry *Entry4e) SetLevel(v log.Lvl) {

}

func (entry *Entry4e) WithField(key string, value interface{}) *Entry4e {
	return &Entry4e{entry.Entry.WithFields(logrus.Fields{key: value})}
}

func (entry *Entry4e) Printj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Print()
}

func (entry *Entry4e)  Debugj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Debug()
}

func (entry *Entry4e)  Infoj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Info()
}

func (entry *Entry4e)  Warnj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Warn()
}

func (entry *Entry4e)  Errorj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Error()
}

func (entry *Entry4e)  Fatalj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Fatal()
}

func (entry *Entry4e)  Panicj(j log.JSON) {
	entry.WithFields(logrus.Fields(j)).Panic()
}


