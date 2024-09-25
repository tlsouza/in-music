package log

import (
	"go.uber.org/zap"
)

type TestLogger struct{}

func (l TestLogger) Info(msg string, fields ...zap.Field)   {}
func (l TestLogger) Warn(msg string, fields ...zap.Field)   {}
func (l TestLogger) Error(msg string, fields ...zap.Field)  {}
func (l TestLogger) DPanic(msg string, fields ...zap.Field) {}
func (l TestLogger) Panic(msg string, fields ...zap.Field)  {}
