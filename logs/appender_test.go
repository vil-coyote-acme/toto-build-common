package logs

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

// will test the appender for write operation while
// calling it with a superior level than minimal required
func Test_write_should_append_new_line_on_sup_log_level(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.write(ERROR, []interface {}{"toto"})
	// then
	assert.Equal(t,bufString.String(), "toto")
}

// will test the appender for write operation while
// calling it with the minimal required log level
func Test_write_should_append_new_line_on_minimal_log_level(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.write(INFO, []interface {}{"toto"})
	// then
	assert.Equal(t,bufString.String(), "toto")
}

// will test the appender for write operation while
// calling it for a too low log level for appender configuration
func Test_write_should_not_append_new_line(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.write(DEBUG, []interface {}{"toto"})
	// then
	assert.Equal(t, "", bufString.String())
}


// will test the appender for writef operation while
// calling it with a superior level than minimal required
func Test_writef_should_append_new_line_on_sup_log_level(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.writef(ERROR, "the number is : %d", []interface {}{1})
	// then
	assert.Equal(t,bufString.String(), "the number is : 1")
}

// will test the appender for write foperation while
// calling it with the minimal required log level
func Test_writef_should_append_new_line_on_minimal_log_level(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.writef(INFO, "the number is : %d", []interface {}{1})
	// then
	assert.Equal(t,bufString.String(), "the number is : 1")
}

// will test the appender for writef operation while
// calling it for a too low log level for appender configuration
func Test_writef_should_not_append_new_line(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	// when
	appender.writef(DEBUG, "the number is : %d", []interface {}{1})
	// then
	assert.Equal(t, "", bufString.String())
}
