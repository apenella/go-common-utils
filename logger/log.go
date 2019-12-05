package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// Constants definitions
const (
	ERROR int8 = iota
	WARN
	INFO
	DEBUG

	TAB int8 = iota
	JSON

	TabFormat             = "%s	%s	%s"
	JsonFormat            = "{\"timestamp\": \"%s\", \"level\": \"%s\", \"message\": \"%s\"}"
	TimestampLayoutFormat = "2006-01-02 15:04:00"
)

// The Message object used by sigleton pattern
var log *Log = nil

type Log struct {
	writer      io.Writer
	level       int8
	format      string
	layout      string
	timeLayout  string
	messageChan chan interface{}
	quitChan    chan bool
}

func printMachine() {

	if log == nil {
		return
	}
	defer close(log.messageChan)
	defer close(log.quitChan)

	for {
		select {
		case msg := <-log.messageChan:
			//fmt.Fprintln(log.writer, msg)
			fmt.Fprintf(log.writer, "%v\n", msg)
		case <-log.quitChan:
			log = nil
			return
		}
	}
}

func (l *Log) messageFormat(layout string, level string, message interface{}) string {
	// return fmt.Sprintf(l.format, level, message)
	return fmt.Sprintf(l.format, time.Now().Format(layout), level, message)
}

func new() *Log {
	return &Log{
		writer:      os.Stdout,
		level:       INFO,
		layout:      TimestampLayoutFormat,
		format:      TabFormat,
		messageChan: make(chan interface{}),
		quitChan:    make(chan bool),
	}
}

func GetLog() *Log {
	if log == nil {
		//fmt.Println(">>>>Create")
		log = new()

		go printMachine()
	}
	return log
}

func StopLog() {
	log.quitChan <- true
}

func Reset() error {
	if log == nil {
		return errors.New("Logger has not been initialized")
	}

	StopLog()
	log = new()
	go printMachine()
	return nil
}

func (l *Log) SetFormat(format int8) error {
	if l == nil {
		return errors.New("Logger has not been initialized")
	}

	if format == TAB || format == JSON {
		switch format {
		case TAB:
			l.format = TabFormat
		case JSON:
			l.format = JsonFormat
		}
	} else {
		return errors.New("Invalid format")
	}

	return nil
}

func (l *Log) SetLayout(layout string) error {
	if l == nil {
		return errors.New("Logger has not been initialized")
	}
	l.layout = layout

	return nil
}

func (l *Log) SetLevel(level int8) error {
	if l == nil {
		return errors.New("Logger has not been initialized")
	}
	l.level = level

	return nil
}

func (l *Log) SetWriter(writer io.Writer) error {
	if l == nil {
		return errors.New("Logger has not been initialized")
	}
	l.writer = writer

	return nil
}

func (l *Log) Info(m ...interface{}) {
	if l.level >= INFO {
		l.messageChan <- l.messageFormat(l.layout, "INFO", m)
	}
}

func (l *Log) Warn(m ...interface{}) {
	if l.level >= WARN {
		l.messageChan <- l.messageFormat(l.layout, "WARN", m)
	}
}

func (l *Log) Error(m ...interface{}) {
	if l.level >= ERROR {
		l.messageChan <- l.messageFormat(l.layout, "ERROR", m)
	}
}

func (l *Log) Debug(m ...interface{}) {
	if l.level >= DEBUG {
		l.messageChan <- l.messageFormat(l.layout, "DEBUG", m)
	}
}
