package fortio

import (
	"fmt"
	"log"
	"os"
)

var (
	// Log used in the client. Initially its set to an empty logger and must be
	// set with a specific logging implementation in order to log messages
	Log = Logger(EmptyLogger{})

	//LogLevels represents the different supported levels for logging messages
	LogLevels = struct {
		Debug LogLevel
		Info  LogLevel
		Warn  LogLevel
		Error LogLevel
		Fatal LogLevel
		Off   LogLevel
	}{
		LogLevel(0),
		LogLevel(1),
		LogLevel(2),
		LogLevel(3),
		LogLevel(4),
		LogLevel(5),
	}
)

// LogLevel is a specific level of logging
type LogLevel int

// Logger is an interface the provides basic logging functionality that can be implemented to wrap
// an existing logging framework
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// EmptyLogger does not log any log messages
type EmptyLogger struct{}

// Debug will log out the args at the debug level if debug level or lower is set
func (l EmptyLogger) Debug(args ...interface{}) {}

// Debugf will log out the args using the format string at the debug level if debug level or lower is set
func (l EmptyLogger) Debugf(format string, args ...interface{}) {}

// Info will log out the args at the info level if info level or lower is set
func (l EmptyLogger) Info(args ...interface{}) {}

// Infof will log out the args using the format string at the info level if info level or lower is set
func (l EmptyLogger) Infof(format string, args ...interface{}) {}

// Warn will log out the args at the warn level if warn level or lower is set
func (l EmptyLogger) Warn(args ...interface{}) {}

// Warnf will log out the args using the format string at the warn level if warn level or lower is set
func (l EmptyLogger) Warnf(format string, args ...interface{}) {}

// Error will log out the args at the error level if error level or lower is set
func (l EmptyLogger) Error(args ...interface{}) {}

// Errorf will log out the args using the format string at the error level if error level or lower is set
func (l EmptyLogger) Errorf(format string, args ...interface{}) {}

// Fatal will log out the args at the fatal level if fatal level or lower is set
func (l EmptyLogger) Fatal(args ...interface{}) {}

// Fatalf will log out the args using the format string at the fatal level if fatal level or lower is set
func (l EmptyLogger) Fatalf(format string, args ...interface{}) {}

// StdLogger will write log messages to either stdout or stderr with an appropriate
// prefix depending on the level of the log message
type StdLogger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
	level    LogLevel
}

// NewStdLogger creates a new StdLogger using the provided level and flags
func NewStdLogger(level LogLevel, flags int) *StdLogger {
	return &StdLogger{
		debugLog: log.New(os.Stdout, "[DEBUG] ", flags),
		infoLog:  log.New(os.Stdout, "[INFO] ", flags),
		warnLog:  log.New(os.Stdout, "[WARN] ", flags),
		errorLog: log.New(os.Stderr, "[ERROR] ", flags),
		fatalLog: log.New(os.Stderr, "[FATAL] ", flags),
		level:    level,
	}
}

// Debug will log out the args at the debug level if debug level or lower is set
func (l *StdLogger) Debug(args ...interface{}) {
	if l.LogLevel() <= LogLevels.Debug {
		l.debugLog.Println(args...)
	}
}

// Debugf will log out the args using the format string at the debug level if debug level or lower is set
func (l *StdLogger) Debugf(format string, args ...interface{}) {
	if l.LogLevel() <= LogLevels.Debug {
		l.debugLog.Println(fmt.Sprintf(format, args...))
	}
}

// Info will log out the args at the info level if info level or lower is set
func (l *StdLogger) Info(args ...interface{}) {
	if l.LogLevel() <= LogLevels.Info {
		l.infoLog.Println(args...)
	}
}

// Infof will log out the args using the format string at the info level if info level or lower is set
func (l *StdLogger) Infof(format string, args ...interface{}) {
	if l.LogLevel() <= LogLevels.Info {
		l.infoLog.Println(fmt.Sprintf(format, args...))
	}
}

// Warn will log out the args at the warn level if warn level or lower is set
func (l *StdLogger) Warn(args ...interface{}) {
	if l.LogLevel() <= LogLevels.Warn {
		l.warnLog.Println(args...)
	}
}

// Warnf will log out the args using the format string at the warn level if warn level or lower is set
func (l *StdLogger) Warnf(format string, args ...interface{}) {
	if l.LogLevel() <= LogLevels.Warn {
		l.warnLog.Println(fmt.Sprintf(format, args...))
	}
}

// Error will log out the args at the error level if error level or lower is set
func (l *StdLogger) Error(args ...interface{}) {
	if l.LogLevel() <= LogLevels.Error {
		l.errorLog.Println(args...)
	}
}

// Errorf will log out the args using the format string at the error level if error level or lower is set
func (l *StdLogger) Errorf(format string, args ...interface{}) {
	if l.LogLevel() <= LogLevels.Error {
		l.errorLog.Println(fmt.Sprintf(format, args...))
	}
}

// Fatal will log out the args at the fatal level if fatal level or lower is set
func (l *StdLogger) Fatal(args ...interface{}) {
	if l.LogLevel() <= LogLevels.Fatal {
		l.fatalLog.Println(args...)
	}
	os.Exit(1)
}

// Fatalf will log out the args using the format string at the fatal level if fatal level or lower is set
func (l *StdLogger) Fatalf(format string, args ...interface{}) {
	if l.LogLevel() <= LogLevels.Error {
		l.fatalLog.Println(fmt.Sprintf(format, args...))
	}
	os.Exit(1)
}

// LogLevel returns LogLevel of the Logger
func (l *StdLogger) LogLevel() LogLevel {
	return l.level
}

// SetLogLevel will set the LogLevel for the logger to the specified LogLevel
func (l *StdLogger) SetLogLevel(level LogLevel) {
	l.level = level
}
