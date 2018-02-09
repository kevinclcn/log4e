package log4e

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func TestLogger(t *testing.T)  {
	buf := bytes.NewBuffer([]byte{})
	logger.Out = buf
	Info("hello")
	assert.Contains(t, buf.String(), "msg=hello", "log")
	assert.Contains(t, buf.String(), "level=info", "log")
}

func TestWithFields(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	logger.Out = buf

	WithFields(Fields{
		"requestId": "123",
	}).Info()

	assert.Contains(t, buf.String(), "requestId=123", "")
}

func TestEchoLogger(t *testing.T)  {
	var v echo.Logger
	v = logger
	buf := bytes.NewBuffer([]byte{})
	v.SetOutput(buf)

	json := log.JSON{"msg": "I'm a test message"}
	v.Infoj(json)

	assert.Contains(t, buf.String(), "msg=\"I'm a test message\"", "log")

	v = WithFields(Fields{"field1":"field1"})
	v.Infoj(json)
	assert.Contains(t, buf.String(), "field1=field1", "")
}