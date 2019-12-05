package log

import (
	"bytes"
	"errors"
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
	l.SetLayout("")
	l.SetWriter(io.Writer(&w))
	l.Info("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "	INFO	[hi!]\n", w.String())

	t.Log("Testing Warn method")
	l.SetLevel(WARN)
	w.Reset()
	l.Warn("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "	WARN	[hi!]\n", w.String())

	t.Log("Testing Error method")
	l.SetLevel(ERROR)
	w.Reset()
	l.Error("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "	ERROR	[hi!]\n", w.String())

	t.Log("Testing Debug method")
	l.SetLevel(DEBUG)
	w.Reset()
	l.Debug("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "	DEBUG	[hi!]\n", w.String())

	t.Log("Testing info message with debug level")
	l.SetLevel(DEBUG)
	w.Reset()
	l.Info("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "	INFO	[hi!]\n", w.String())

	t.Log("Testing skip message due low level log")
	l.SetLevel(INFO)
	w.Reset()
	l.Debug("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 0, len(w.String()), "")

	t.Log("Testing json message")
	l.SetFormat(JSON)
	l.Info("hi!")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "{\"timestamp\": \"\", \"level\": \"INFO\", \"message\": \"[hi!]\"}\n", w.String())

	StopLog()
}

func TestSetFormat(t *testing.T) {
	tests := []struct {
		desc   string
		log    *Log
		format int8
		res    string
		err    error
	}{
		{
			desc:   "Testing nil log",
			log:    nil,
			format: TAB,
			res:    "",
			err:    errors.New("Logger has not been initialized"),
		},
		{
			desc:   "Testing set invalid format",
			log:    GetLog(),
			format: int8(29),
			res:    "",
			err:    errors.New("Invalid format"),
		},
		{
			desc:   "Testing set tab format",
			log:    GetLog(),
			format: TAB,
			res:    TabFormat,
			err:    nil,
		},
		{
			desc:   "Testing set json format",
			log:    GetLog(),
			format: JSON,
			res:    JsonFormat,
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.log.SetFormat(test.format)
		if err != nil {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res, test.log.format, "Format not expected")
		}
	}
}
