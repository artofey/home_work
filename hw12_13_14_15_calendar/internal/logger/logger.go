package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel = map[string]zapcore.Level{
		"DEBUG":   zap.DebugLevel,
		"INFO":    zap.InfoLevel,
		"WARNING": zap.WarnLevel,
		"ERROR":   zap.ErrorLevel,
		"FATAL":   zap.FatalLevel,
	}
)

type Logger struct {
	zp *zap.SugaredLogger
}

func New(level string, file string) (*Logger, error) {
	lv, ok := logLevel[level]
	if !ok {
		lv = zap.InfoLevel
	}
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(lv),
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout", file},
		ErrorOutputPaths: []string{"stderr", file},
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("init logger error: %w", err)
	}
	return &Logger{
		zp: log.Sugar(),
	}, nil
}

func (l Logger) Info(msg string) {
	l.zp.Info(msg)
}

func (l Logger) Error(msg string) {
	l.zp.Error(msg)
}

func (l Logger) Sync() error {
	err := l.zp.Sync()
	if err != nil {
		return fmt.Errorf("sync logger error: %w", err)
	}
	return nil
}
