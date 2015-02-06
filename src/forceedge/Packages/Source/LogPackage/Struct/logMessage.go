package logPackage

import (
	"fmt"
	"log"
)

type Log struct {
	text string
}

func NewLog() *Log {
	return new(Log)
}

func (l *Log) SetText(text string) *Log {
	l.text = fmt.Sprintf("Log entry: " + text)

	return l
}

func (l *Log) PrintLog() {
	if l.text == "" {
		log.Fatal("text is not set for the log")
	}

	fmt.Println(l.text)
}
