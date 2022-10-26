package src

import (
	"fmt"
	"io"
)

const (
	BOLD = "\033[1m"
	WARN = "\033[33m"
	INFO = "\033[32m"
	FAIL = "\033[31m"
	SUCCEED = "\033[32m"
	END = "\033[0m"
)

type Printer struct {
	Writer io.Writer
}

func NewPrinter(w *io.Writer) *Printer {
	return &Printer{*w}
}

func (p *Printer) Print(s string) {
	p.Writer.Write([]byte(s))
}

func (p *Printer) Printf(s string, args ...interface{}) {
	p.Writer.Write([]byte(fmt.Sprintf(s, args...)))
}

func (p *Printer) Println(s string) {
	p.Writer.Write([]byte(s + "\n"))
}

func (p *Printer) Printlnf(s string, args ...interface{}) {
	p.Writer.Write([]byte(fmt.Sprintf(s, args...) + "\n"))
}

func (p *Printer) PrintlnWarn(s string) {
	p.Writer.Write([]byte(WARN + s + END + "\n"))
}

func (p *Printer) PrintlnWarnf(s string, args ...interface{}) {
	p.Writer.Write([]byte(WARN + fmt.Sprintf(s, args...) + END + "\n"))
}

func (p *Printer) PrintlnInfo(s string) {
	p.Writer.Write([]byte(INFO + s + END + "\n"))
}

func (p *Printer) PrintlnInfof(s string, args ...interface{}) {
	p.Writer.Write([]byte(INFO + fmt.Sprintf(s, args...) + END + "\n"))
}

func (p *Printer) PrintlnFail(s string) {
	p.Writer.Write([]byte(FAIL + s + END + "\n"))
}

func (p *Printer) PrintlnFailf(s string, args ...interface{}) {
	p.Writer.Write([]byte(FAIL + fmt.Sprintf(s, args...) + END + "\n"))
}

func (p *Printer) PrintlnSucceed(s string) {
	p.Writer.Write([]byte(SUCCEED + s + END + "\n"))
}

func (p *Printer) PrintlnSucceedf(s string, args ...interface{}) {
	p.Writer.Write([]byte(SUCCEED + fmt.Sprintf(s, args...) + END + "\n"))
}
