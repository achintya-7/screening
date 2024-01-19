package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLoggerWithCorrelationId creates a new logger with correlation id
func NewLoggerWithCorrelationId(ctx *gin.Context) *zap.Logger {
	correlationId := ctx.GetHeader("X-Correlation-Id")
	if correlationId == "" {
		correlationId = uuid.New().String()
	}

	config := zap.NewProductionConfig()
    config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    logger, _ := config.Build()

    return logger.With(zap.String("correlationId", correlationId))
}
