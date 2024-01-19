package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewLoggerWithCorrelationId(ctx *gin.Context) *zap.Logger {
	correlationId := ctx.GetHeader("X-Correlation-Id")
	if correlationId == "" {
		correlationId = uuid.New().String()
	}

	return zap.L().With(zap.String("correlationId", correlationId))
}
