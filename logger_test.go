package logger

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func show() {
	Debug("debug message")
	Debugf("debugf message: %s", "test...")
	Debugw("test...", "key", "value")

	Info("info message")
	Infof("infof message: %s", "test...")
	Infow("test...", "key", "value")

	Warn("warn message")
	Warnf("warnf message: %s", "test...")
	Warnw("test...", "key", "value")

	Error("Error message")
	Errorf("Errorf message: %s", "test...")
	Errorw("test...", "key", "value")
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
	l.Set("debug")
	SetLevel(l)

	With("trace_id", "274ac2bbf9d5")
	With("span_id", "383d60f1")

	show()
}

func TestDevFileLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.OutputPaths = []string{"stdout", "dev.log", "test.log"}
	config.DisableStacktrace = true
	config.DisableCaller = false
	config.ShortTime = true
	config.EnableColor = false
	SetConfig(config)

	show()
}

func TestColorLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.DisableStacktrace = true
	config.DisableCaller = true
	SetConfig(config)

	show()
}

func TestLoggerInterface(t *testing.T) {
	l := GetLogger()
	l.Info("get logger...")
}

func TestLogger_With(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = InfoLevel
	config.DisableStacktrace = true
	config.DisableCaller = false
	config.EnableColor = true
	config.ShortTime = true
	config.InitialFields = map[string]interface{}{"initKey1": "initVal1"}
	SetConfig(config)

	myCtxLogger := With("key1", "val1").
		With("key2", "value2").
		With("key3", "value3").
		With("key4", "value4").
		With("key5", "value5").
		With("key6", "value6").
		With("key7", "value7").
		With("key8", "value8")
	myCtxLogger.Infow("Infow: hello myCtxLogger", "aaa", "bbb")
	// myCtxLogger.With("this should not shown up", "wtf")
	myCtxLogger.Debugw("Debugw: this should not shown", "nickname", "Dr. Janes")
	myCtxLogger.Info("Info: this", " is", " just", " a", " demo")
	myCtxLogger.Warnf("Warnf: this is just a %v", "simple demo")

	fmt.Println("-----------------------------------------------------")

	ninthLogger := myCtxLogger.With("key9", "v9")
	ninthLogger.Infow("Infow: this is the ninthLogger context", "second_other_key", "second_other_value")

	fmt.Println("-----------------------------------------------------")

	tenthLogger := ninthLogger.With("key10", "v10")
	tenthLogger.Infow("Infow: this is the tenthLogger context", "third_other_key", "third_other_value")

	fmt.Println("-----------------------------------------------------")
	// global logger does not impacted
	show()
}

func Benchmark_With(b *testing.B) {
	zapConfig := &zap.Config{
		Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
		DisableCaller: false,
		Sampling:      &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:      "console",
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		OutputPaths:   []string{"/dev/null"},
		InitialFields: map[string]interface{}{"initKey1": "initVal1"},
	}
	zapLogger, _ := zapConfig.Build()
	sugar := zapLogger.Sugar()

	config := NewDevelopmentConfig()
	config.Level = InfoLevel
	config.DisableStacktrace = true
	config.DisableCaller = false
	config.EnableColor = false
	config.ShortTime = true
	config.InitialFields = map[string]interface{}{"initKey1": "initVal1"}
	config.OutputPaths = []string{"/dev/null"}
	SetConfig(config)

	b.Run("ZapSugarWith", func(b *testing.B) {
		myCtxLogger := sugar.With(fmt.Sprintf("with_key_%d", b.N), fmt.Sprintf("with_val_%d", b.N)).
			With("key2", "value2").
			With("key3", "value3").
			With("key4", "value4").
			With("key5", "value5").
			With("key6", "value6").
			With("key7", "value7").
			With("key8", "value8")
		myCtxLogger.Infow("Infow: hello myCtxLogger", "aaa", "bbb")
		// myCtxLogger.With("this should not shown up", "wtf")
		myCtxLogger.Debugw("Debugw: this should not shown", "nickname", "Dr. Janes")
		myCtxLogger.Info("Info: this", " is", " just", " a", " demo")
		myCtxLogger.Warnf("Warnf: this is just a %v", "simple demo")
	})

	b.Run("LoggerWith", func(b *testing.B) {
		myCtxLogger := With(fmt.Sprintf("with_key_%d", b.N), fmt.Sprintf("with_val_%d", b.N)).
			With("key2", "value2").
			With("key3", "value3").
			With("key4", "value4").
			With("key5", "value5").
			With("key6", "value6").
			With("key7", "value7").
			With("key8", "value8")
		myCtxLogger.Infow("Infow: hello myCtxLogger", "aaa", "bbb")
		// myCtxLogger.With("this should not shown up", "wtf")
		myCtxLogger.Debugw("Debugw: this should not shown", "nickname", "Dr. Janes")
		myCtxLogger.Info("Info: this", " is", " just", " a", " demo")
		myCtxLogger.Warnf("Warnf: this is just a %v", "simple demo")
	})
}

func TestLogger_With_Invalid(t *testing.T) {
	config := NewProductionConfig()
	config.InitialFields = map[string]interface{}{"initKey1": "initVal1"}
	// config.Level = DebugLevel
	// config.DisableStacktrace = true
	config.DisableCaller = false
	// config.EnableColor = true
	// config.ShortTime = true
	SetConfig(config)

	myCtxLogger := With("hello", "world", "not_paired_key")

	// this will panic if with Development config
	Infow("xxx", "ooo")
	// this will panic if with Development config
	myCtxLogger.Infow("xxx", "ooo")

	fmt.Println("-----------------------------------------------------")

	myCtxLogger.Info("xxx", "ooo")
	myCtxLogger.Error("xxx", 4399)
	myCtxLogger.Errorf("traffic %d", 122)
}

func TestConfigClone(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = InfoLevel
	config.DisableStacktrace = true
	config.DisableCaller = false
	config.EnableColor = false
	config.ShortTime = true
	config.InitialFields = map[string]interface{}{"initKey1": "initVal1"}
	config.OutputPaths = []string{"/dev/null"}
	config.buildZapConfig()

	cloned := config.clone()
	cloned.EnableColor = true
	cloned.OutputPaths[0] = "/dev/stderr"
	cloned.OutputPaths = append(cloned.OutputPaths, "/dev/stdout")
	cloned.InitialFields["hello"] = "world"
	cloned.zapConfig.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	cloned.zapConfig.Sampling.Thereafter = 300

	if config.EnableColor == cloned.EnableColor {
		t.Fatalf("cloned.EnableColor fail")
	}
	if len(config.OutputPaths) == len(cloned.OutputPaths) {
		t.Fatalf("cloned.OutputPaths fail")
	}
	if len(config.InitialFields) == len(cloned.InitialFields) {
		t.Fatalf("cloned.InitialFields fail")
	}
	if config.zapConfig.Level == cloned.zapConfig.Level {
		t.Fatalf("cloned.zapConfig.Level fail")
	}
	if config.zapConfig.Sampling.Thereafter == cloned.zapConfig.Sampling.Thereafter {
		t.Fatalf("cloned.Sampling.Thereafter fail")
	}
}
