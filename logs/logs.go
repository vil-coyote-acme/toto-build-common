package logs


type Level int

type Logger struct {
	prefix string
	appender *Appender
}

func NewLogger(prefix string, appender *Appender) *Logger {
	l := new(Logger)
	l.prefix = prefix
	l.appender = appender
	return l
}

func (l *Logger) Debug(structs ...interface{}) {
}

func (l *Logger) Debugf(template string, structs ...interface{}) {
}

func (l *Logger) Info(structs ...interface{}) {
}

func (l *Logger) Infof(template string, structs ...interface{}) {
}

func (l *Logger) Error(structs ...interface{}) {
}

func (l *Logger) Errorf(template string, structs ...interface{}) {
}