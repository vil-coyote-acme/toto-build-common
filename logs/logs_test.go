package logs

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

// will test the debug function when debug
// level is set
func Test_Debug_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debug("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the debug function when info
// level is set
func Test_Debug_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debug("titi")
	// then
	assert.Equal(t, "", bufString.String())
}

// will test the debug function when error
// level is set
func Test_Debug_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debug("titi")
	// then
	assert.Equal(t, "", bufString.String())
}

// will test the debugf function when debug
// level is set
func Test_Debugf_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debugf("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the debugf function when info
// level is set
func Test_Debugf_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debugf("%s", "titi")
	// then
	assert.Equal(t, "", bufString.String())
}

// will test the debugf function when error
// level is set
func Test_Debugf_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Debugf("%s", "titi")
	// then
	assert.Equal(t, "", bufString.String())
}


// will test the info function when debug
// level is set
func Test_Info_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Info("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the info function when info
// level is set
func Test_Info_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Info("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the info function when error
// level is set
func Test_Info_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Info("titi")
	// then
	assert.Equal(t, "", bufString.String())
}


// will test the infof function when debug
// level is set
func Test_Infof_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Infof("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the infof function when info
// level is set
func Test_Infof_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Infof("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the infof function when error
// level is set
func Test_Infof_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Infof("%s", "titi")
	// then
	assert.Equal(t, "", bufString.String())
}


// will test the error function when debug
// level is set
func Test_Error_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Error("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the error function when info
// level is set
func Test_Error_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Error("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the error function when error
// level is set
func Test_Error_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Error("titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the ierrorf function when debug
// level is set
func Test_Errorf_When_Debug_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(DEBUG, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Errorf("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the errorf function when info
// level is set
func Test_Errorf_When_Info_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(INFO, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Errorf("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}

// will test the infof function when error
// level is set
func Test_Errorf_When_Error_mode(t *testing.T) {
	// given
	bufString := bytes.NewBufferString("")
	appender := newAppender(ERROR, bufString)
	logger := NewLogger("toto : ", appender)
	// when
	logger.Errorf("%s", "titi")
	// then
	assert.Contains(t,bufString.String(), "toto : titi")
}
