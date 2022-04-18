/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-18 15:25:04
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-18 15:39:54
 * @FilePath: /api/pkg/router/middleware/log.go
 * @Description: log middleware
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Log(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()
		latency := time.Since(start)
		logger.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("Host", ctx.GetHeader("Host")),
			zap.String("X-Real-IP", ctx.GetHeader("X-Real-IP")),
			zap.String("X-Forwarded-For", ctx.GetHeader("X-Forwarded-For")),
			zap.String("X-Forwarded-Proto", ctx.GetHeader("X-Forwarded-Proto")),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("latency", latency),
		)
	}
}
