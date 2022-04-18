/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-18 04:32:49
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-18 14:49:18
 * @FilePath: /api/cmd/main.go
 * @Description: main
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kokutas/vs/api/pkg/log"
	"github.com/kokutas/vs/api/pkg/router"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("NAMESPACE", "api gateway")
}
func main() {
	// zap logger
	logger := log.NewZapLogger().With(
		zap.Namespace(viper.GetString("NAMESPACE")),
	)
	router := router.SetupRouter(logger)
	srv := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", viper.GetInt("PORT")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		logger.Sugar().Infof("http server lister : %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.With(
				zap.Error(err),
			).Error("http server closed.")
			logger.Fatal(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warn("shutting down http server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.With(
			zap.Error(err),
		).Error("http server forced to shutdown.")
		logger.Fatal(err.Error())
	}
	logger.Info("http server exiting")
}
