package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	lvl string
	zp  *zap.SugaredLogger
}

func New(level string) (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("init logger error: %w", err)
	}
	return &Logger{
		zp:  log.Sugar(),
		lvl: level,
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
