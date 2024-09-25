package log

import (
	"api/pkg/configs"
	"context"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type Adapter interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
}

type Logger struct {
	adapter Adapter
	Ctx     context.Context
}

func (logger *Logger) Info(msg string, fields ...zap.Field) {
	logger.adapter.Info(msg, fields...)
}

func (logger *Logger) Warn(msg string, fields ...zap.Field) {
	logger.adapter.Warn(msg, fields...)
}

func (logger *Logger) Error(err error, msg string, fields ...zap.Field) {
	fields = append(fields, Error(err))
	logger.adapter.Error(msg, fields...)

}

func (logger *Logger) DPanic(msg string, fields ...zap.Field) {
	logger.adapter.DPanic(msg, fields...)
}

func (logger *Logger) Panic(msg string, fields ...zap.Field) {
	logger.adapter.Panic(msg, fields...)
}

func NewWithPortOut(ctx context.Context, portOut string) (logger *Logger) {
	ctxWithPortOut := context.WithValue(ctx, "portOut", portOut)
	logger = New(ctxWithPortOut)
	return
}

func New(ctx context.Context) *Logger {
	var logger *Logger

	if configs.ENV == "testing" {
		logger = &Logger{
			adapter: TestLogger{},
			Ctx:     ctx,
		}
		return logger
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig = ecszap.ECSCompatibleEncoderConfig(config.EncoderConfig)
	config.OutputPaths = []string{"stdout"}

	config.Level.UnmarshalText([]byte(configs.LOG_LEVEL))

	zapLogger, _ := config.Build(ecszap.WrapCoreOption(), zap.AddCaller(), zap.AddCallerSkip(1))
	zapLogger = zapLogger.Named(configs.APP_NAME)

	portIn := ctx.Value("portIn")
	if portIn != nil {
		zapLogger = zapLogger.With(zap.String("application.hexagon.portin", portIn.(string)))
	}

	portOut := ctx.Value("portOut")
	if portOut != nil {
		zapLogger = zapLogger.With(zap.String("application.hexagon.portout", portOut.(string)))
	}

	traceId := ctx.Value("traceId")
	if traceId != nil {
		zapLogger = zapLogger.With(zap.String("trace.id", traceId.(string)))
	}
	requestId := ctx.Value("requestId")
	if requestId != nil {
		zapLogger = zapLogger.With(zap.String("request.id", requestId.(string)))
	}

	return &Logger{adapter: zapLogger, Ctx: ctx}
}
