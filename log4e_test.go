package log4e

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T)  {
	buf := bytes.NewBuffer([]byte{})
	logger.Out = buf
	Info("hello")
	assert.Contains(t, buf.String(), "msg=hello", "log should contain hello")
	assert.Contains(t, buf.String(), "level=info", "log should be info level")
}