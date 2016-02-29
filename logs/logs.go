package logs


type Level int

const (
	DEBUG Level = iota
	INFO
	ERROR
)

type Logger struct {
	prefix string
	Level Level
}

func NewLogger(prefix string, Level Level) *Logger {
	l := new(Logger)
	l.prefix = prefix
	l.Level = Level
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