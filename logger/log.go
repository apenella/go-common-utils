package log

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Constants definitions
const ERROR int = 0
const WARN int = 1
const INFO int = 2
const DEBUG int = 3

// The Message object used by sigleton pattern
var log *Log = nil

type Log struct {
	writer io.Writer
	level  int
	//layout = "2006-01-02 15:04:00"
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
			fmt.Fprintln(log.writer, msg)
		case <-log.quitChan:
			log = nil
			return
		}
	}
}

func messageFormat(level string, message interface{}) string {
	return fmt.Sprintf("%s %v", level, message)
	//	return fmt.Sprintf("%s %s %v", time.Now().Format(layout), level, message)
}

func GetLog() *Log {
	if log == nil {
		//fmt.Println(">>>>Create")
		log = &Log{
			writer:      os.Stdout,
			level:       INFO,
			messageChan: make(chan interface{}),
			quitChan:    make(chan bool),
		}

		go printMachine()
	}
	return log
}

func StopLog() {
	log.quitChan <- true
}

func (l *Log) SetLevel(level int) error {
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
		l.messageChan <- messageFormat("[INFO]", m)
	}
}

func (l *Log) Warn(m ...interface{}) {
	if l.level >= WARN {
		l.messageChan <- messageFormat("[WARN]", m)
	}
}

func (l *Log) Error(m ...interface{}) {
	if l.level >= ERROR {
		l.messageChan <- messageFormat("[ERROR]", m)
	}
}

func (l *Log) Debug(m ...interface{}) {
	if l.level >= DEBUG {
		l.messageChan <- messageFormat("[DEBUG]", m)
	}
}
