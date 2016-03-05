package logs

type Level int

type Logger struct {
	prefix   string
	appender *Appender
}

func NewLogger(prefix string, appender *Appender) *Logger {
	l := new(Logger)
	l.prefix = prefix
	l.appender = appender
	return l
}

func (l *Logger) Debug(structs ...interface{}) {
	l.appender.write(DEBUG, l.appendPrefix(structs))
}

func (l *Logger) Debugf(template string, structs ...interface{}) {
	l.appender.writef(DEBUG, l.prefix + template, structs)
}

func (l *Logger) Info(structs ...interface{}) {
	l.appender.write(INFO, l.appendPrefix(structs))
}

func (l *Logger) Infof(template string, structs ...interface{}) {
	l.appender.writef(INFO, l.prefix + template, structs)
}

func (l *Logger) Error(structs ...interface{}) {
	l.appender.write(ERROR, l.appendPrefix(structs))
}

func (l *Logger) Errorf(template string, structs ...interface{}) {
	l.appender.writef(ERROR, l.prefix + template, structs)
}

func (l *Logger) appendPrefix(structs []interface{}) []interface{} {
	obj := make([]interface{}, 0)
	obj = append(obj, l.prefix)
	obj = append(obj, structs...)
	return obj
}