package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	// nolint: gofumpt
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	// nolint: gofumpt
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	// nolint: gofumpt
	Log(level Level, args ...interface{})
	Logf(level Level, template string, args ...interface{})
	Logw(level Level, msg string, keysAndValues ...interface{})
	With(keyValues ...interface{}) Logger
	SetLevel(level Level)
}

type logger struct {
	config    *Config
	logger    *zap.SugaredLogger
	zapLogger *zap.Logger
}

var (
	defaultConfig = NewDefaultConfig()
	l             = newLogger(defaultConfig)
)

func SetConfig(config *Config) {
	l = newLogger(config)
}

func SetLevel(level Level) {
	l.SetLevel(level)
}

func GetLevel() Level {
	return l.Level()
}

func SetOutputPaths(outputPaths []string) {
	l.config.OutputPaths = outputPaths
	l = newLogger(l.config)
}

func NewLogger(config *Config) Logger {
	l := newLogger(config)
	return l
}

func GetLogger() Logger {
	return l
}

func newLogger(config *Config) *logger {
	config.buildZapConfig()

	zapLogger, err := config.zapConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on build zap logger (%s)", err)
		return nil
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(config.CallerSkip))

	return &logger{
		logger:    zapLogger.Sugar(),
		zapLogger: zapLogger,
		config:    config,
	}
}

func (l *logger) SetLevel(level Level) {
	l.config.Level = level
	l.config.zapConfig.Level.SetLevel(zapcore.Level(level))
}

func (l *logger) Level() Level {
	return l.config.Level
}

func (l *logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

func (l *logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

func (l *logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *logger) Log(level Level, args ...interface{}) {
	// nolint: exhaustive
	switch level {
	case DebugLevel:
		l.logger.Debug(args...)
	case InfoLevel:
		l.logger.Info(args...)
	case WarnLevel:
		l.logger.Warn(args...)
	case ErrorLevel:
		l.logger.Error(args...)
	default:
		l.logger.Info(args...)
	}
}

func (l *logger) Logf(level Level, template string, args ...interface{}) {
	// nolint: exhaustive
	switch level {
	case DebugLevel:
		l.logger.Debugf(template, args...)
	case InfoLevel:
		l.logger.Infof(template, args...)
	case WarnLevel:
		l.logger.Warnf(template, args...)
	case ErrorLevel:
		l.logger.Errorf(template, args...)
	default:
		l.logger.Infof(template, args...)
	}
}

func (l *logger) Logw(level Level, msg string, keysAndValues ...interface{}) {
	// nolint: exhaustive
	switch level {
	case DebugLevel:
		l.logger.Debugw(msg, keysAndValues...)
	case InfoLevel:
		l.logger.Infow(msg, keysAndValues...)
	case WarnLevel:
		l.logger.Warnw(msg, keysAndValues...)
	case ErrorLevel:
		l.logger.Errorw(msg, keysAndValues...)
	default:
		l.logger.Infow(msg, keysAndValues...)
	}
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}

func (l *logger) With(keyValues ...interface{}) Logger {
	if len(keyValues) == 0 {
		return l
	}

	newLogger := &logger{
		config:    l.config.clone(),
		logger:    l.logger.With(keyValues...),
		zapLogger: l.zapLogger,
	}
	return newLogger
}

func (l *logger) GetZapLogger() *zap.Logger {
	return l.zapLogger
}
