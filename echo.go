package logger

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// EchoLogger extend infra logger
type EchoLogger struct {
}

// Singleton logger
var singletonLogger = &EchoLogger{}

// GetEchoLogger return singleton logger
func GetEchoLogger() *EchoLogger {
	return singletonLogger
}

func GetEchoLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			errMsg := ""
			if err != nil {
				errMsg = err.Error()
			}

			switch req.RequestURI {
			case "/healthz", "/metrics":
				return
			}
			if errMsg == "" {
				Infow("echo", "remote_ip", c.RealIP(), "host", req.Host, "method", req.Method, "uri", req.RequestURI,
					"user_agent", req.UserAgent(), "status", res.Status, "msg", errMsg, "latecy", int64(stop.Sub(start)),
					"latency_human", stop.Sub(start).String(), "bytes_in", echo.HeaderContentLength, "bytes_out", res.Size)
			} else {
				Errorw("echo", "remote_ip", c.RealIP(), "host", req.Host, "method", req.Method, "uri", req.RequestURI,
					"user_agent", req.UserAgent(), "status", res.Status, "msg", errMsg, "latecy", int64(stop.Sub(start)),
					"latency_human", stop.Sub(start).String(), "bytes_in", echo.HeaderContentLength, "bytes_out", res.Size)
			}
			return
		}
	}
}

// To infra logger Level
func toLoggerLevel(level log.Lvl) Level {
	switch level {
	case log.DEBUG:
		return DebugLevel
	case log.INFO:
		return InfoLevel
	case log.WARN:
		return WarnLevel
	case log.ERROR:
		return ErrorLevel
	}
	return InfoLevel
}

// To Echo.log.Lvl
func toEchoLevel(level Level) log.Lvl {
	switch level {
	case DebugLevel:
		return log.DEBUG
	case InfoLevel:
		return log.INFO
	case WarnLevel:
		return log.WARN
	case ErrorLevel:
		return log.ERROR
	}
	return log.OFF
}

// Output return logger io.Writer
func (l *EchoLogger) Output() io.Writer {
	return os.Stdout
}

// SetOutput logger io.Writer
func (l *EchoLogger) SetOutput(w io.Writer) {
}

// Level return logger level
func (l *EchoLogger) Level() log.Lvl {
	return toEchoLevel(InfoLevel)
}

// SetLevel logger level
func (l *EchoLogger) SetLevel(v log.Lvl) {
	SetLevel(toLoggerLevel(v))
}

// SetHeader logger header
// This function do nothing
func (l *EchoLogger) SetHeader(h string) {
	// do nothing
}

// Prefix return logger prefix
// This function do nothing
func (l *EchoLogger) Prefix() string {
	return ""
}

// SetPrefix logger prefix
// This function do nothing
func (l *EchoLogger) SetPrefix(p string) {
	// do nothing
}

// Print output message of print level
func (l *EchoLogger) Print(i ...interface{}) {
	Info(i...)
}

// Printf output format message of print level
func (l *EchoLogger) Printf(format string, args ...interface{}) {
	Infof(format, args...)
}

// Printj output json of print level
func (l *EchoLogger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Info(string(b))
}

// Debug output message of debug level
func (l *EchoLogger) Debug(i ...interface{}) {
	Debug(i...)
}

// Debugf output format message of debug level
func (l *EchoLogger) Debugf(format string, args ...interface{}) {
	Debugf(format, args...)
}

// Debugj output message of debug level
func (l *EchoLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Debug(string(b))
}

// Info output message of info level
func (l *EchoLogger) Info(i ...interface{}) {
	Info(i...)
}

// Infof output format message of info level
func (l *EchoLogger) Infof(format string, args ...interface{}) {
	Infof(format, args...)
}

// Infoj output json of info level
func (l *EchoLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Info(string(b))
}

// Warn output message of warn level
func (l *EchoLogger) Warn(i ...interface{}) {
	Warn(i...)
}

// Warnf output format message of warn level
func (l *EchoLogger) Warnf(format string, args ...interface{}) {
	Warnf(format, args...)
}

// Warnj output json of warn level
func (l *EchoLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Warn(string(b))
}

// Error output message of error level
func (l *EchoLogger) Error(i ...interface{}) {
	Error(i...)
}

// Errorf output format message of error level
func (l *EchoLogger) Errorf(format string, args ...interface{}) {
	Errorf(format, args...)
}

// Errorj output json of error level
func (l *EchoLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Error(string(b))
}

// Fatal output message of fatal level
func (l *EchoLogger) Fatal(i ...interface{}) {
	Fatal(i...)
}

// Fatalf output format message of fatal level
func (l *EchoLogger) Fatalf(format string, args ...interface{}) {
	Fatalf(format, args...)
}

// Fatalj output json of fatal level
func (l *EchoLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Fatal(string(b))
}

// Panic output message of panic level
func (l *EchoLogger) Panic(i ...interface{}) {
	Panic(i...)
}

// Panicf output format message of panic level
func (l *EchoLogger) Panicf(format string, args ...interface{}) {
	Panicf(format, args...)
}

// Panicj output json of panic level
func (l *EchoLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	Panic(string(b))
}
