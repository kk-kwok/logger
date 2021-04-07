package logger

import (
	"testing"
)

func show() {
	Debug("debug message")
	Debugf("debug message: %s", "test...")
	Info("info message")
	Infof("infof message: %s", "test...")
	Infow("infow", "message", "test...")
	Warn("warn message")
	Warnf("warnf message: %s", "test...")
	Error("Error message")
	Errorf("Errorf message: %s", "test...")
}

func TestDefultLogger(t *testing.T) {
	show()
	Info("------------------")
	var l Level
	l.Set("error")
	SetLevel(l)
	show()
}

func TestLevel(t *testing.T) {
	show()
}

func TestProdLogger(t *testing.T) {
	config := NewProductionConfig(FieldPair{"service", "client_string"}, FieldPair{"version", "v0.1.0-5309251"})
	SetConfig(config)
	var l Level
	l.Set("info")
	SetLevel(l)

	show()
}

func TestDevFileLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.OutputPaths = []string{"dev.log", "test.log"}
	config.DisableStacktrace = true
	config.DisableCaller = false
	SetConfig(config)

	show()
}

func TestColorLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.DisableStacktrace = true
	config.DisableCaller = true
	config.EnableColor = true
	config.ShortTime = true
	SetConfig(config)

	show()
}

func TestLoggerInterface(t *testing.T) {
	l := GetLogger()
	l.Info("get logger...")
}
