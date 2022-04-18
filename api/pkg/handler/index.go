/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-18 05:01:06
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-18 07:00:09
 * @FilePath: /api/pkg/handler/index.go
 * @Description: index handler
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"message":   http.StatusText(http.StatusOK),
		"namespace": viper.GetString("NAMESPACE"),
	})
	return
}
