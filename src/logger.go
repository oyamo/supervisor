package src

import (
	"fmt"
	"io"
	"time"
)

type Logger struct {
	Printer *Printer
}

func NewLogger(w *io.Writer) *Logger {
	return &Logger{NewPrinter(w)}
}

func (l *Logger) getTime() string {
	time := time.Now()
	return time.Format("2006-01-02 15:04:05")
}


func (l *Logger) Println(s string) {
	l.Printer.Printlnf("%s %s", l.getTime(), s)
}

func (l *Logger) Printlnf(s string, args ...interface{}) {
	l.Printer.Printlnf("%s %s", l.getTime(), fmt.Sprintf(s, args...))
}

func (l *Logger) PrintlnWarn(s string) {
	l.Printer.PrintlnWarnf("%s %s", l.getTime(), s)
}

func (l *Logger) PrintlnWarnf(s string, args ...interface{}) {
	l.Printer.PrintlnWarnf("%s %s", l.getTime(), fmt.Sprintf(s, args...))
}

func (l *Logger) PrintlnInfo(s string) {
	l.Printer.PrintlnInfof("%s %s", l.getTime(), s)
}

func (l *Logger) PrintlnInfof(s string, args ...interface{}) {
	l.Printer.PrintlnInfof("%s %s", l.getTime(), fmt.Sprintf(s, args...))
}

func (l *Logger) PrintlnFail(s string) {
	l.Printer.PrintlnFailf("%s %s", l.getTime(), s)
}

func (l *Logger) PrintlnFailf(s string, args ...interface{}) {
	l.Printer.PrintlnFailf("%s %s", l.getTime(), fmt.Sprintf(s, args...))
}

func (l *Logger) PrintlnSucceed(s string) {
	l.Printer.PrintlnSucceedf("%s %s", l.getTime(), s)
}

func (l *Logger) PrintlnSucceedf(s string, args ...interface{}) {
	l.Printer.PrintlnSucceedf("%s %s", l.getTime(), fmt.Sprintf(s, args...))
}