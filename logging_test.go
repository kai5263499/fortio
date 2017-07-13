package fortio

import (
	"log"
	"testing"
)

func TestStdLoggerDebug(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Debug, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Debug("Debug message")
}

func TestStdLoggerDebugWithInfoLevel(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Info, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Debug("No Debug message")
}

func TestStdLoggerDebugf(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Debug, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Debugf("%s", "Debug format message")
}

func TestStdLoggerInfo(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Info, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Info("Info message")
}

func TestStdLoggerInfof(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Info, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Infof("%s", "Info format message")
}

func TestStdLoggerInfoWithWarnLevel(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Warn, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Info("No Info message")
}

func TestStdLoggerWarn(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Warn, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Warn("Warn message")
}

func TestStdLoggerWarnf(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Warn, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Warnf("%s", "Warn format message")
}

func TestStdLoggerWarnWithErrorLevel(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Error, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Warn("No Warn message")
}

func TestStdLoggerError(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Error, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Error("Error message")
}

func TestStdLoggerErrorf(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Error, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Errorf("%s", "Error format message")
}

func TestStdLoggerErrorWithOffLevel(t *testing.T) {
	stdLog := NewStdLogger(LogLevels.Off, log.Ldate|log.Ltime|log.Lshortfile)
	stdLog.Error("Error message")
}
