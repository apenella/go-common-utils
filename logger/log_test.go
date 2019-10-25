package log

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLog(t *testing.T) {
	t.Log("Testing GetLog method")
	expected := &Log{
		writer:      os.Stdout,
		level:       INFO,
		messageChan: make(chan interface{}),
		quitChan:    make(chan bool),
	}

	log := GetLog()
	StopLog()

	assert.Equal(t, expected.writer, log.writer)
	assert.Equal(t, expected.level, log.level)
}

func TestMessages(t *testing.T) {
	t.Log("Testing Info method")
	var w bytes.Buffer

	l := GetLog()
	l.SetWriter(io.Writer(&w))
	l.Info("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "[INFO] [hi!]\n", w.String())

	t.Log("Testing Warn method")
	l.SetLevel(WARN)
	w.Reset()
	l.Warn("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "[WARN] [hi!]\n", w.String())

	t.Log("Testing Error method")
	l.SetLevel(ERROR)
	w.Reset()
	l.Error("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "[ERROR] [hi!]\n", w.String())

	t.Log("Testing Debug method")
	l.SetLevel(DEBUG)
	w.Reset()
	l.Debug("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "[DEBUG] [hi!]\n", w.String())

	t.Log("Testing info message with debug level")
	l.SetLevel(DEBUG)
	w.Reset()
	l.Info("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "[INFO] [hi!]\n", w.String())

	t.Log("Testing skip message due low level log")
	l.SetLevel(INFO)
	w.Reset()
	l.Debug("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, len(w.String()), "")

	StopLog()
}

