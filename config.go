package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FieldPair []string

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level Level `json:"level" yaml:"level"`

	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development" yaml:"development"`

	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`

	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`

	// Encoding sets the logger's encoding. Valid values are "json" and "console"
	Encoding string `json:"encoding" yaml:"encoding"`

	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`

	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`

	EnableColor bool
	ShortTime   bool

	CallerSkip int
	zapConfig  *zap.Config
}

func NewProductionConfig(fields ...FieldPair) *Config {
	return &Config{
		Level:             InfoLevel,
		Development:       false,
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		CallerSkip:        2,
		DisableStacktrace: false,
		InitialFields:     genInitialFields(fields),
	}
}

func NewDevelopmentConfig(fields ...FieldPair) *Config {
	return &Config{
		Level:             DebugLevel,
		Development:       true,
		ShortTime:         true,
		EnableColor:       true,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
		CallerSkip:        2,
		DisableStacktrace: true,
		InitialFields:     genInitialFields(fields),
	}
}

func NewDefaultConfig(fields ...FieldPair) *Config {
	return NewProductionConfig(fields...)
}

func genInitialFields(args []FieldPair) map[string]interface{} {
	fields := make(map[string]interface{})
	for _, f := range args {
		fields[f[0]] = f[1]
	}
	return fields
}

func (c *Config) buildZapConfig() {
	encoderConfig := c.newCustomEncoderConfig()

	zapConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(c.Level)),
		Development:       c.Development,
		DisableCaller:     c.DisableCaller,
		DisableStacktrace: c.DisableStacktrace,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          c.Encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       c.OutputPaths,
		InitialFields:     c.InitialFields,
	}
	c.zapConfig = zapConfig
}

func (c *Config) newCustomEncoderConfig() zapcore.EncoderConfig {
	encodeLevel := zapcore.LowercaseLevelEncoder
	if c.EnableColor {
		encodeLevel = zapcore.LowercaseColorLevelEncoder
	}
	encodeTime := zapcore.ISO8601TimeEncoder
	if c.ShortTime {
		encodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			type appendTimeEncoder interface {
				AppendTimeLayout(time.Time, string)
			}
			layout := "2006-01-02 15:04:05"
			if enc, ok := enc.(appendTimeEncoder); ok {
				enc.AppendTimeLayout(t, layout)
				return
			}
			enc.AppendString(t.Format(layout))
		}
	}
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     encodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (c *Config) clone() *Config {
	cloned := *c
	cloned.OutputPaths = make([]string, len(c.OutputPaths))
	copy(cloned.OutputPaths, c.OutputPaths)
	cloned.InitialFields = make(map[string]interface{})
	for k, v := range c.InitialFields {
		cloned.InitialFields[k] = v
	}
	if cloned.zapConfig != nil {
		cloned.buildZapConfig()
	}
	return &cloned
}
