package grpc

import (
	"context"
	"errors"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	"github.com/kokutas/pb/area"
	endpoint1 "github.com/kokutas/vs/area/pkg/endpoint"
	service "github.com/kokutas/vs/area/pkg/service"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.AreaService, error) {
	var saveEndpoint endpoint.Endpoint
	{
		saveEndpoint = grpc1.NewClient(conn, "area.AreaService", "Save", encodeSaveRequest, decodeSaveResponse, area.SaveReply{}, options["Save"]...).Endpoint()
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpc1.NewClient(conn, "area.AreaService", "Get", encodeGetRequest, decodeGetResponse, area.GetReply{}, options["Get"]...).Endpoint()
	}

	var getByIdEndpoint endpoint.Endpoint
	{
		getByIdEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetById", encodeGetByIdRequest, decodeGetByIdResponse, area.GetByIdReply{}, options["GetById"]...).Endpoint()
	}

	var getByCodeEndpoint endpoint.Endpoint
	{
		getByCodeEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByCode", encodeGetByCodeRequest, decodeGetByCodeResponse, area.GetByCodeReply{}, options["GetByCode"]...).Endpoint()
	}

	var getByNameEndpoint endpoint.Endpoint
	{
		getByNameEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByName", encodeGetByNameRequest, decodeGetByNameResponse, area.GetByNameReply{}, options["GetByName"]...).Endpoint()
	}

	var getByTimeZoneEndpoint endpoint.Endpoint
	{
		getByTimeZoneEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByTimeZone", encodeGetByTimeZoneRequest, decodeGetByTimeZoneResponse, area.GetByTimeZoneReply{}, options["GetByTimeZone"]...).Endpoint()
	}

	var getByLanguageEndpoint endpoint.Endpoint
	{
		getByLanguageEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByLanguage", encodeGetByLanguageRequest, decodeGetByLanguageResponse, area.GetByLanguageReply{}, options["GetByLanguage"]...).Endpoint()
	}

	var getByStateEndpoint endpoint.Endpoint
	{
		getByStateEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByState", encodeGetByStateRequest, decodeGetByStateResponse, area.GetByStateReply{}, options["GetByState"]...).Endpoint()
	}

	var getByCreateTimeEndpoint endpoint.Endpoint
	{
		getByCreateTimeEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByCreateTime", encodeGetByCreateTimeRequest, decodeGetByCreateTimeResponse, area.GetByCreateTimeReply{}, options["GetByCreateTime"]...).Endpoint()
	}

	var getByUpdateTimeEndpoint endpoint.Endpoint
	{
		getByUpdateTimeEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByUpdateTime", encodeGetByUpdateTimeRequest, decodeGetByUpdateTimeResponse, area.GetByUpdateTimeReply{}, options["GetByUpdateTime"]...).Endpoint()
	}

	var getByDeleteTimeEndpoint endpoint.Endpoint
	{
		getByDeleteTimeEndpoint = grpc1.NewClient(conn, "area.AreaService", "GetByDeleteTime", encodeGetByDeleteTimeRequest, decodeGetByDeleteTimeResponse, area.GetByDeleteTimeReply{}, options["GetByDeleteTime"]...).Endpoint()
	}

	var updateCodeEndpoint endpoint.Endpoint
	{
		updateCodeEndpoint = grpc1.NewClient(conn, "area.AreaService", "UpdateCode", encodeUpdateCodeRequest, decodeUpdateCodeResponse, area.UpdateCodeReply{}, options["UpdateCode"]...).Endpoint()
	}

	var updateNameEndpoint endpoint.Endpoint
	{
		updateNameEndpoint = grpc1.NewClient(conn, "area.AreaService", "UpdateName", encodeUpdateNameRequest, decodeUpdateNameResponse, area.UpdateNameReply{}, options["UpdateName"]...).Endpoint()
	}

	var updateTimeZoneEndpoint endpoint.Endpoint
	{
		updateTimeZoneEndpoint = grpc1.NewClient(conn, "area.AreaService", "UpdateTimeZone", encodeUpdateTimeZoneRequest, decodeUpdateTimeZoneResponse, area.UpdateTimeZoneReply{}, options["UpdateTimeZone"]...).Endpoint()
	}

	var updateLanguageEndpoint endpoint.Endpoint
	{
		updateLanguageEndpoint = grpc1.NewClient(conn, "area.AreaService", "UpdateLanguage", encodeUpdateLanguageRequest, decodeUpdateLanguageResponse, area.UpdateLanguageReply{}, options["UpdateLanguage"]...).Endpoint()
	}

	var updateStateEndpoint endpoint.Endpoint
	{
		updateStateEndpoint = grpc1.NewClient(conn, "area.AreaService", "UpdateState", encodeUpdateStateRequest, decodeUpdateStateResponse, area.UpdateStateReply{}, options["UpdateState"]...).Endpoint()
	}

	var deleteByIdEndpoint endpoint.Endpoint
	{
		deleteByIdEndpoint = grpc1.NewClient(conn, "area.AreaService", "DeleteById", encodeDeleteByIdRequest, decodeDeleteByIdResponse, area.DeleteByIdReply{}, options["DeleteById"]...).Endpoint()
	}

	var deleteByCodeEndpoint endpoint.Endpoint
	{
		deleteByCodeEndpoint = grpc1.NewClient(conn, "area.AreaService", "DeleteByCode", encodeDeleteByCodeRequest, decodeDeleteByCodeResponse, area.DeleteByCodeReply{}, options["DeleteByCode"]...).Endpoint()
	}

	var deleteByNameEndpoint endpoint.Endpoint
	{
		deleteByNameEndpoint = grpc1.NewClient(conn, "area.AreaService", "DeleteByName", encodeDeleteByNameRequest, decodeDeleteByNameResponse, area.DeleteByNameReply{}, options["DeleteByName"]...).Endpoint()
	}

	var deleteByStateEndpoint endpoint.Endpoint
	{
		deleteByStateEndpoint = grpc1.NewClient(conn, "area.AreaService", "DeleteByState", encodeDeleteByStateRequest, decodeDeleteByStateResponse, area.DeleteByStateReply{}, options["DeleteByState"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		DeleteByCodeEndpoint:    deleteByCodeEndpoint,
		DeleteByIdEndpoint:      deleteByIdEndpoint,
		DeleteByNameEndpoint:    deleteByNameEndpoint,
		DeleteByStateEndpoint:   deleteByStateEndpoint,
		GetByCodeEndpoint:       getByCodeEndpoint,
		GetByCreateTimeEndpoint: getByCreateTimeEndpoint,
		GetByDeleteTimeEndpoint: getByDeleteTimeEndpoint,
		GetByIdEndpoint:         getByIdEndpoint,
		GetByLanguageEndpoint:   getByLanguageEndpoint,
		GetByNameEndpoint:       getByNameEndpoint,
		GetByStateEndpoint:      getByStateEndpoint,
		GetByTimeZoneEndpoint:   getByTimeZoneEndpoint,
		GetByUpdateTimeEndpoint: getByUpdateTimeEndpoint,
		GetEndpoint:             getEndpoint,
		SaveEndpoint:            saveEndpoint,
		UpdateCodeEndpoint:      updateCodeEndpoint,
		UpdateLanguageEndpoint:  updateLanguageEndpoint,
		UpdateNameEndpoint:      updateNameEndpoint,
		UpdateStateEndpoint:     updateStateEndpoint,
		UpdateTimeZoneEndpoint:  updateTimeZoneEndpoint,
	}, nil
}

// encodeSaveRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Save request to a gRPC request.
func encodeSaveRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeSaveResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSaveResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Get request to a gRPC request.
func encodeGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &area.GetRequest{}, nil
	// return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetResponse(_ context.Context, reply interface{}) (interface{}, error) {
	response := reply.(*area.GetReply)
	return endpoint1.GetResponse{
		Response: &area.GetReply{
			Reply: &area.Reply{
				Code:   response.Reply.Code,
				Reason: response.Reply.Reason,
			},
			Areas: response.Areas,
		},
	}, nil
}

// encodeGetByIdRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetById request to a gRPC request.
func encodeGetByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByIdResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByIdResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByCode request to a gRPC request.
func encodeGetByCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByCodeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByCodeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByNameRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByName request to a gRPC request.
func encodeGetByNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByNameResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByNameResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByTimeZoneRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByTimeZone request to a gRPC request.
func encodeGetByTimeZoneRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByTimeZoneResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByTimeZoneResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByLanguageRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByLanguage request to a gRPC request.
func encodeGetByLanguageRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByLanguageResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByLanguageResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByStateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByState request to a gRPC request.
func encodeGetByStateRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByStateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByStateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByCreateTimeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByCreateTime request to a gRPC request.
func encodeGetByCreateTimeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByCreateTimeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByCreateTimeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByUpdateTimeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByUpdateTime request to a gRPC request.
func encodeGetByUpdateTimeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByUpdateTimeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByUpdateTimeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByDeleteTimeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetByDeleteTime request to a gRPC request.
func encodeGetByDeleteTimeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeGetByDeleteTimeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetByDeleteTimeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateCode request to a gRPC request.
func encodeUpdateCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeUpdateCodeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateCodeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateNameRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateName request to a gRPC request.
func encodeUpdateNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeUpdateNameResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateNameResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateTimeZoneRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateTimeZone request to a gRPC request.
func encodeUpdateTimeZoneRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeUpdateTimeZoneResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateTimeZoneResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateLanguageRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateLanguage request to a gRPC request.
func encodeUpdateLanguageRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeUpdateLanguageResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateLanguageResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateStateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateState request to a gRPC request.
func encodeUpdateStateRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeUpdateStateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateStateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByIdRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain DeleteById request to a gRPC request.
func encodeDeleteByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeDeleteByIdResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDeleteByIdResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain DeleteByCode request to a gRPC request.
func encodeDeleteByCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeDeleteByCodeResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDeleteByCodeResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByNameRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain DeleteByName request to a gRPC request.
func encodeDeleteByNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeDeleteByNameResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDeleteByNameResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByStateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain DeleteByState request to a gRPC request.
func encodeDeleteByStateRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}

// decodeDeleteByStateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDeleteByStateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}
