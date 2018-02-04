package log4e

import (
	"github.com/sirupsen/logrus"
	"github.com/labstack/gommon/log"
)

type Entry4e struct {
	*logrus.Entry
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


