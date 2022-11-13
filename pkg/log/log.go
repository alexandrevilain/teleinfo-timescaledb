package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateLogger() (logr.Logger, error) {
	zc := zap.NewProductionConfig()
	zc.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(-2))
	z, err := zc.Build()
	if err != nil {
		return logr.Discard(), err
	}

	return zapr.NewLogger(z), nil
}
