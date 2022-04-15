/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-16 05:22:57
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-16 06:21:34
 * @FilePath: /area/client/main.go
 * @Description: TODO
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	grpc1 "github.com/go-kit/kit/transport/grpc"
	"github.com/kokutas/pb/area"
	grpc2 "github.com/kokutas/vs/area/client/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	logrus.SetReportCaller(true)
}
func main() {
	// grpc.WithInsecure() 禁用 安全传输
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	options := make(map[string][]grpc1.ClientOption)
	s, err := grpc2.New(conn, options)
	if err != nil {
		fmt.Println(err)
	}
	response, err := s.Get(context.TODO(), new(area.GetRequest))
	if err != nil {
		logrus.Fatal(err)
		s, ok := status.FromError(err)
		if !ok {
			s = status.New(codes.Unknown, err.Error())
		}
		fmt.Println(s.Code(), s.Message())
	}
	fmt.Println(response)
}
