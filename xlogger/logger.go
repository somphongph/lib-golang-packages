package xlogger

import (
	"fmt"
	"os"

	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level zapcore.Level
	Name  string // application name
}

// MustInit initializes and replaces zap's global logger.
func Init(level zapcore.Level, name string, pretty bool) {
	cfg := Config{
		Level: level,
		Name:  name,
	}

	if !pretty {
		logger := defaultConfig(cfg)
		zap.ReplaceGlobals(logger)
	} else {
		log, err := createPrettyLogger(cfg)
		if err != nil {
			fmt.Printf("cannot init pretty xlogger: %v; fallback to default", err)
		} else {
			zap.ReplaceGlobals(log)
		}
	}
}

func Debugf(template string, args ...any) {
	zap.S().Debugf(template, args...)
}
func Infof(template string, args ...any) {
	zap.S().Infof(template, args...)
}
func Warnf(template string, args ...any) {
	zap.S().Warnf(template, args...)
}
func Errorf(template string, args ...any) {
	zap.S().Errorf(template, args...)
}

func createPrettyLogger(inputCfg Config) (*zap.Logger, error) {
	// Retrieve log level from configuration
	logLevel := inputCfg.Level

	// Initialize Zap production config
	cfg := zap.NewProductionConfig()
	cfg.Development = true
	cfg.Encoding = "pretty_console" // Use pretty console encoder

	// Set the log level in the config
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	// Customize the encoder configuration for better readability
	encoderConfig := prettyconsole.NewEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	cfg.EncoderConfig = encoderConfig

	// Build and return the logger
	return cfg.Build()
}

func defaultConfig(inputCfg Config) *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= inputCfg.Level
	})

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), highPriority),
	)

	logger := zap.New(core, zap.AddCaller())
	logger.Named(inputCfg.Name)

	return logger
}
