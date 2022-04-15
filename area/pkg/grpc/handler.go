package grpc

import (
	"context"
	"errors"
	"net/http"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/kokutas/pb/area"
	endpoint "github.com/kokutas/vs/area/pkg/endpoint"
)

// makeSaveHandler creates the handler logic
func makeSaveHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SaveEndpoint, decodeSaveRequest, encodeSaveResponse, options...)
}

// decodeSaveResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Save request.
// TODO implement the decoder
func decodeSaveRequest(_ context.Context, r interface{}) (interface{}, error) {
	// return nil, errors.New("'Area' Decoder is not impelemented")
	return nil, nil
}

// encodeSaveResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSaveResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented save")
}
func (g *grpcServer) Save(ctx context.Context, req *area.SaveRequest) (*area.SaveReply, error) {
	_, rep, err := g.save.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.SaveReply), nil
}

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Get request.
// TODO implement the decoder
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	// return nil, errors.New("'Area' Decoder is not impelemented")
	req := r.(*area.GetRequest)
	return endpoint.GetRequest{
		Request: req,
	}, nil
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.GetResponse)
	return &area.GetReply{
		Reply: &area.Reply{
			Success: true,
			Code:    http.StatusOK,
			Reason:  "ok",
		},
		Areas: reply.Response.Areas,
	}, nil
}
func (g *grpcServer) Get(ctx context.Context, req *area.GetRequest) (*area.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetReply), nil
}

// makeGetByIdHandler creates the handler logic
func makeGetByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByIdEndpoint, decodeGetByIdRequest, encodeGetByIdResponse, options...)
}

// decodeGetByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetById request.
// TODO implement the decoder
func decodeGetByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented by id")
}
func (g *grpcServer) GetById(ctx context.Context, req *area.GetByIdRequest) (*area.GetByIdReply, error) {
	_, rep, err := g.getById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByIdReply), nil
}

// makeGetByCodeHandler creates the handler logic
func makeGetByCodeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByCodeEndpoint, decodeGetByCodeRequest, encodeGetByCodeResponse, options...)
}

// decodeGetByCodeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByCode request.
// TODO implement the decoder
func decodeGetByCodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*area.GetRequest)
	return endpoint.GetRequest{
		Request: req,
	}, nil
	// return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByCodeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByCodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented by code")
}
func (g *grpcServer) GetByCode(ctx context.Context, req *area.GetByCodeRequest) (*area.GetByCodeReply, error) {
	_, rep, err := g.getByCode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByCodeReply), nil
}

// makeGetByNameHandler creates the handler logic
func makeGetByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByNameEndpoint, decodeGetByNameRequest, encodeGetByNameResponse, options...)
}

// decodeGetByNameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByName request.
// TODO implement the decoder
func decodeGetByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByNameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByName(ctx context.Context, req *area.GetByNameRequest) (*area.GetByNameReply, error) {
	_, rep, err := g.getByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByNameReply), nil
}

// makeGetByTimeZoneHandler creates the handler logic
func makeGetByTimeZoneHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByTimeZoneEndpoint, decodeGetByTimeZoneRequest, encodeGetByTimeZoneResponse, options...)
}

// decodeGetByTimeZoneResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByTimeZone request.
// TODO implement the decoder
func decodeGetByTimeZoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByTimeZoneResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByTimeZoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByTimeZone(ctx context.Context, req *area.GetByTimeZoneRequest) (*area.GetByTimeZoneReply, error) {
	_, rep, err := g.getByTimeZone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByTimeZoneReply), nil
}

// makeGetByLanguageHandler creates the handler logic
func makeGetByLanguageHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByLanguageEndpoint, decodeGetByLanguageRequest, encodeGetByLanguageResponse, options...)
}

// decodeGetByLanguageResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByLanguage request.
// TODO implement the decoder
func decodeGetByLanguageRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByLanguageResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByLanguageResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByLanguage(ctx context.Context, req *area.GetByLanguageRequest) (*area.GetByLanguageReply, error) {
	_, rep, err := g.getByLanguage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByLanguageReply), nil
}

// makeGetByStateHandler creates the handler logic
func makeGetByStateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByStateEndpoint, decodeGetByStateRequest, encodeGetByStateResponse, options...)
}

// decodeGetByStateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByState request.
// TODO implement the decoder
func decodeGetByStateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByStateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByStateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByState(ctx context.Context, req *area.GetByStateRequest) (*area.GetByStateReply, error) {
	_, rep, err := g.getByState.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByStateReply), nil
}

// makeGetByCreateTimeHandler creates the handler logic
func makeGetByCreateTimeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByCreateTimeEndpoint, decodeGetByCreateTimeRequest, encodeGetByCreateTimeResponse, options...)
}

// decodeGetByCreateTimeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByCreateTime request.
// TODO implement the decoder
func decodeGetByCreateTimeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByCreateTimeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByCreateTimeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByCreateTime(ctx context.Context, req *area.GetByCreateTimeRequest) (*area.GetByCreateTimeReply, error) {
	_, rep, err := g.getByCreateTime.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByCreateTimeReply), nil
}

// makeGetByUpdateTimeHandler creates the handler logic
func makeGetByUpdateTimeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByUpdateTimeEndpoint, decodeGetByUpdateTimeRequest, encodeGetByUpdateTimeResponse, options...)
}

// decodeGetByUpdateTimeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByUpdateTime request.
// TODO implement the decoder
func decodeGetByUpdateTimeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByUpdateTimeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByUpdateTimeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByUpdateTime(ctx context.Context, req *area.GetByUpdateTimeRequest) (*area.GetByUpdateTimeReply, error) {
	_, rep, err := g.getByUpdateTime.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByUpdateTimeReply), nil
}

// makeGetByDeleteTimeHandler creates the handler logic
func makeGetByDeleteTimeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetByDeleteTimeEndpoint, decodeGetByDeleteTimeRequest, encodeGetByDeleteTimeResponse, options...)
}

// decodeGetByDeleteTimeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetByDeleteTime request.
// TODO implement the decoder
func decodeGetByDeleteTimeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeGetByDeleteTimeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetByDeleteTimeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) GetByDeleteTime(ctx context.Context, req *area.GetByDeleteTimeRequest) (*area.GetByDeleteTimeReply, error) {
	_, rep, err := g.getByDeleteTime.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.GetByDeleteTimeReply), nil
}

// makeUpdateCodeHandler creates the handler logic
func makeUpdateCodeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateCodeEndpoint, decodeUpdateCodeRequest, encodeUpdateCodeResponse, options...)
}

// decodeUpdateCodeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateCode request.
// TODO implement the decoder
func decodeUpdateCodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateCodeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateCodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) UpdateCode(ctx context.Context, req *area.UpdateCodeRequest) (*area.UpdateCodeReply, error) {
	_, rep, err := g.updateCode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.UpdateCodeReply), nil
}

// makeUpdateNameHandler creates the handler logic
func makeUpdateNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateNameEndpoint, decodeUpdateNameRequest, encodeUpdateNameResponse, options...)
}

// decodeUpdateNameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateName request.
// TODO implement the decoder
func decodeUpdateNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateNameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) UpdateName(ctx context.Context, req *area.UpdateNameRequest) (*area.UpdateNameReply, error) {
	_, rep, err := g.updateName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.UpdateNameReply), nil
}

// makeUpdateTimeZoneHandler creates the handler logic
func makeUpdateTimeZoneHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTimeZoneEndpoint, decodeUpdateTimeZoneRequest, encodeUpdateTimeZoneResponse, options...)
}

// decodeUpdateTimeZoneResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTimeZone request.
// TODO implement the decoder
func decodeUpdateTimeZoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateTimeZoneResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTimeZoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTimeZone(ctx context.Context, req *area.UpdateTimeZoneRequest) (*area.UpdateTimeZoneReply, error) {
	_, rep, err := g.updateTimeZone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.UpdateTimeZoneReply), nil
}

// makeUpdateLanguageHandler creates the handler logic
func makeUpdateLanguageHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateLanguageEndpoint, decodeUpdateLanguageRequest, encodeUpdateLanguageResponse, options...)
}

// decodeUpdateLanguageResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateLanguage request.
// TODO implement the decoder
func decodeUpdateLanguageRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateLanguageResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateLanguageResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) UpdateLanguage(ctx context.Context, req *area.UpdateLanguageRequest) (*area.UpdateLanguageReply, error) {
	_, rep, err := g.updateLanguage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.UpdateLanguageReply), nil
}

// makeUpdateStateHandler creates the handler logic
func makeUpdateStateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateStateEndpoint, decodeUpdateStateRequest, encodeUpdateStateResponse, options...)
}

// decodeUpdateStateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateState request.
// TODO implement the decoder
func decodeUpdateStateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeUpdateStateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateStateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) UpdateState(ctx context.Context, req *area.UpdateStateRequest) (*area.UpdateStateReply, error) {
	_, rep, err := g.updateState.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.UpdateStateReply), nil
}

// makeDeleteByIdHandler creates the handler logic
func makeDeleteByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteByIdEndpoint, decodeDeleteByIdRequest, encodeDeleteByIdResponse, options...)
}

// decodeDeleteByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteById request.
// TODO implement the decoder
func decodeDeleteByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) DeleteById(ctx context.Context, req *area.DeleteByIdRequest) (*area.DeleteByIdReply, error) {
	_, rep, err := g.deleteById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.DeleteByIdReply), nil
}

// makeDeleteByCodeHandler creates the handler logic
func makeDeleteByCodeHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteByCodeEndpoint, decodeDeleteByCodeRequest, encodeDeleteByCodeResponse, options...)
}

// decodeDeleteByCodeResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteByCode request.
// TODO implement the decoder
func decodeDeleteByCodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByCodeResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteByCodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) DeleteByCode(ctx context.Context, req *area.DeleteByCodeRequest) (*area.DeleteByCodeReply, error) {
	_, rep, err := g.deleteByCode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.DeleteByCodeReply), nil
}

// makeDeleteByNameHandler creates the handler logic
func makeDeleteByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteByNameEndpoint, decodeDeleteByNameRequest, encodeDeleteByNameResponse, options...)
}

// decodeDeleteByNameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteByName request.
// TODO implement the decoder
func decodeDeleteByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByNameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) DeleteByName(ctx context.Context, req *area.DeleteByNameRequest) (*area.DeleteByNameReply, error) {
	_, rep, err := g.deleteByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.DeleteByNameReply), nil
}

// makeDeleteByStateHandler creates the handler logic
func makeDeleteByStateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteByStateEndpoint, decodeDeleteByStateRequest, encodeDeleteByStateResponse, options...)
}

// decodeDeleteByStateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteByState request.
// TODO implement the decoder
func decodeDeleteByStateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Decoder is not impelemented")
}

// encodeDeleteByStateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteByStateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Area' Encoder is not impelemented")
}
func (g *grpcServer) DeleteByState(ctx context.Context, req *area.DeleteByStateRequest) (*area.DeleteByStateReply, error) {
	_, rep, err := g.deleteByState.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*area.DeleteByStateReply), nil
}
