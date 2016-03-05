package logs

import (
	"io"
	"os"
	"fmt"
)

const (
	DEBUG Level = iota
	INFO
	ERROR
)

type Appender struct {
	level  Level
	writer io.Writer
}

func newAppender(level Level, writer io.Writer) *Appender {
	appender := new(Appender)
	appender.level = level
	appender.writer = writer
	return appender

}

func NewConsoleAppender(level Level) *Appender {
	return newAppender(level, os.Stdout)
}

// todo test what's append with many struct (how many line for exemple...)
// maybe wrapping the writer ?...
func (a *Appender) write(level Level, structs []interface{}) {
	if a.level <= level {
		fmt.Fprint(a.writer, structs...)
	}
}

func (a *Appender) writef(level Level, template string, structs []interface{}) {
	if a.level <= level {
		fmt.Fprintf(a.writer,template, structs...)
	}
}