package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	gormlogger "gorm.io/gorm/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

type logger struct {
	config *Config
	logger *zap.SugaredLogger
}

var ZapLogger *zap.Logger

var (
	defaultConfig = NewDefaultConfig()
	l             = newLogger(defaultConfig)
	gormLogger    = NewGormLogger(context.Background(), l, 5*time.Second)
)

func SetConfig(config *Config) {
	l = newLogger(config)
	gormLogger = NewGormLogger(context.Background(), l, 5*time.Second)
}

func SetLevel(level Level) {
	l.SetLevel(level)
}

func GetLevel() Level {
	return l.Level()
}

func GetLogger() Logger {
	return l
}

func GetGormLogger() gormlogger.Interface {
	return gormLogger
}

func newLogger(config *Config) *logger {
	encoderConfig := NewCustomEncoderConfig()
	zapConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(config.Level)),
		Development:       config.Development,
		DisableCaller:     config.DisableCaller,
		DisableStacktrace: config.DisableStacktrace,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          config.Encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       config.OutputPaths,
		InitialFields:     config.InitialFields,
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on build zap logger (%s)", err)
		return nil
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(config.CallerSkip))
	config.zapConfig = zapConfig
	ZapLogger = zapLogger

	return &logger{
		logger: zapLogger.Sugar(),
		config: config,
	}
}

func NewCustomEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
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

func (l *logger) Sync() error {
	return l.logger.Sync()
}
