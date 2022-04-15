/*
 * @Author: kokutas
 * @Email: xxx
 * @Phone: xxx
 * @Date: 2022-04-16 02:43:44
 * @LastEditors: kokutas
 * @LastEditTime: 2022-04-16 06:26:23
 * @FilePath: /area/pkg/service/service.go
 * @Description: TODO
 * Copyright (c) 2022 by kokutas, All Rights Reserved.
 */
package service

import (
	"context"

	"github.com/kokutas/pb/area"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AreaService describes the service.
type AreaService interface {
	Save(ctx context.Context, request *area.SaveRequest) (response *area.SaveReply, err error)
	Get(ctx context.Context, request *area.GetRequest) (response *area.GetReply, err error)
	GetById(ctx context.Context, request *area.GetByIdRequest) (response *area.GetByIdReply, err error)
	GetByCode(ctx context.Context, request *area.GetByCodeRequest) (response *area.GetByCodeReply, err error)
	GetByName(ctx context.Context, request *area.GetByNameRequest) (response *area.GetByNameReply, err error)
	GetByTimeZone(ctx context.Context, request *area.GetByTimeZoneRequest) (response *area.GetByTimeZoneReply, err error)
	GetByLanguage(ctx context.Context, request *area.GetByLanguageRequest) (response *area.GetByLanguageReply, err error)
	GetByState(ctx context.Context, request *area.GetByStateRequest) (response *area.GetByStateReply, err error)
	GetByCreateTime(ctx context.Context, request *area.GetByCreateTimeRequest) (response *area.GetByCreateTimeReply, err error)
	GetByUpdateTime(ctx context.Context, request *area.GetByUpdateTimeRequest) (response *area.GetByUpdateTimeReply, err error)
	GetByDeleteTime(ctx context.Context, request *area.GetByDeleteTimeRequest) (rsponse *area.GetByDeleteTimeReply, err error)
	UpdateCode(ctx context.Context, request *area.UpdateCodeRequest) (response *area.UpdateCodeReply, err error)
	UpdateName(ctx context.Context, request *area.UpdateNameRequest) (response *area.UpdateNameReply, err error)
	UpdateTimeZone(ctx context.Context, request *area.UpdateTimeZoneRequest) (response *area.UpdateTimeZoneReply, err error)
	UpdateLanguage(ctx context.Context, request *area.UpdateLanguageRequest) (response *area.UpdateLanguageReply, err error)
	UpdateState(ctx context.Context, request *area.UpdateStateRequest) (response *area.UpdateStateReply, err error)
	DeleteById(ctx context.Context, request *area.DeleteByIdRequest) (response *area.DeleteByIdReply, err error)
	DeleteByCode(ctx context.Context, request *area.DeleteByCodeRequest) (response *area.DeleteByCodeReply, err error)
	DeleteByName(ctx context.Context, request *area.DeleteByNameRequest) (response *area.DeleteByNameReply, err error)
	DeleteByState(ctx context.Context, request *area.DeleteByStateRequest) (response *area.DeleteByStateReply, err error)
}

type basicAreaService struct{}

func (b *basicAreaService) Save(ctx context.Context, request *area.SaveRequest) (response *area.SaveReply, err error) {
	// TODO implement the business logic of Save
	return response, err
}
func (b *basicAreaService) Get(ctx context.Context, request *area.GetRequest) (response *area.GetReply, err error) {
	// TODO implement the business logic of Get
	response = &area.GetReply{
		Reply: &area.Reply{
			Success: true,
			Code:    int32(codes.OK),
			Reason:  status.New(codes.OK, "").String(),
		},
		Areas: []*area.Area{
			{
				Id: "11",
			},
		},
	}
	return response, err
}
func (b *basicAreaService) GetById(ctx context.Context, request *area.GetByIdRequest) (response *area.GetByIdReply, err error) {
	// TODO implement the business logic of GetById
	return response, err
}
func (b *basicAreaService) GetByCode(ctx context.Context, request *area.GetByCodeRequest) (response *area.GetByCodeReply, err error) {
	// TODO implement the business logic of GetByCode
	return response, err
}
func (b *basicAreaService) GetByName(ctx context.Context, request *area.GetByNameRequest) (response *area.GetByNameReply, err error) {
	// TODO implement the business logic of GetByName
	return response, err
}
func (b *basicAreaService) GetByTimeZone(ctx context.Context, request *area.GetByTimeZoneRequest) (response *area.GetByTimeZoneReply, err error) {
	// TODO implement the business logic of GetByTimeZone
	return response, err
}
func (b *basicAreaService) GetByLanguage(ctx context.Context, request *area.GetByLanguageRequest) (response *area.GetByLanguageReply, err error) {
	// TODO implement the business logic of GetByLanguage
	return response, err
}
func (b *basicAreaService) GetByState(ctx context.Context, request *area.GetByStateRequest) (response *area.GetByStateReply, err error) {
	// TODO implement the business logic of GetByState
	return response, err
}
func (b *basicAreaService) GetByCreateTime(ctx context.Context, request *area.GetByCreateTimeRequest) (response *area.GetByCreateTimeReply, err error) {
	// TODO implement the business logic of GetByCreateTime
	return response, err
}
func (b *basicAreaService) GetByUpdateTime(ctx context.Context, request *area.GetByUpdateTimeRequest) (response *area.GetByUpdateTimeReply, err error) {
	// TODO implement the business logic of GetByUpdateTime
	return response, err
}
func (b *basicAreaService) GetByDeleteTime(ctx context.Context, request *area.GetByDeleteTimeRequest) (rsponse *area.GetByDeleteTimeReply, err error) {
	// TODO implement the business logic of GetByDeleteTime
	return rsponse, err
}
func (b *basicAreaService) UpdateCode(ctx context.Context, request *area.UpdateCodeRequest) (response *area.UpdateCodeReply, err error) {
	// TODO implement the business logic of UpdateCode
	return response, err
}
func (b *basicAreaService) UpdateName(ctx context.Context, request *area.UpdateNameRequest) (response *area.UpdateNameReply, err error) {
	// TODO implement the business logic of UpdateName
	return response, err
}
func (b *basicAreaService) UpdateTimeZone(ctx context.Context, request *area.UpdateTimeZoneRequest) (response *area.UpdateTimeZoneReply, err error) {
	// TODO implement the business logic of UpdateTimeZone
	return response, err
}
func (b *basicAreaService) UpdateLanguage(ctx context.Context, request *area.UpdateLanguageRequest) (response *area.UpdateLanguageReply, err error) {
	// TODO implement the business logic of UpdateLanguage
	return response, err
}
func (b *basicAreaService) UpdateState(ctx context.Context, request *area.UpdateStateRequest) (response *area.UpdateStateReply, err error) {
	// TODO implement the business logic of UpdateState
	return response, err
}
func (b *basicAreaService) DeleteById(ctx context.Context, request *area.DeleteByIdRequest) (response *area.DeleteByIdReply, err error) {
	// TODO implement the business logic of DeleteById
	return response, err
}
func (b *basicAreaService) DeleteByCode(ctx context.Context, request *area.DeleteByCodeRequest) (response *area.DeleteByCodeReply, err error) {
	// TODO implement the business logic of DeleteByCode
	return response, err
}
func (b *basicAreaService) DeleteByName(ctx context.Context, request *area.DeleteByNameRequest) (response *area.DeleteByNameReply, err error) {
	// TODO implement the business logic of DeleteByName
	return response, err
}
func (b *basicAreaService) DeleteByState(ctx context.Context, request *area.DeleteByStateRequest) (response *area.DeleteByStateReply, err error) {
	// TODO implement the business logic of DeleteByState
	return response, err
}

// NewBasicAreaService returns a naive, stateless implementation of AreaService.
func NewBasicAreaService() AreaService {
	return &basicAreaService{}
}

// New returns a AreaService with all of the expected middleware wired in.
func New(middleware []Middleware) AreaService {
	var svc AreaService = NewBasicAreaService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
