/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-18 04:36:18
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-18 15:33:48
 * @FilePath: /api/pkg/router/router.go
 * @Description: router
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kokutas/vs/api/pkg/handler"
	"github.com/kokutas/vs/api/pkg/router/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger) *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))
	router := gin.New()
	router.Use(gin.Recovery())

	// router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(middleware.Log(logger))
	router.GET("/", handler.Index)
	return router
}
