package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/kokutas/pb/area"
	service "github.com/kokutas/vs/area/pkg/service"
)

// SaveRequest collects the request parameters for the Save method.
type SaveRequest struct {
	Request *area.SaveRequest `json:"request"`
}

// SaveResponse collects the response parameters for the Save method.
type SaveResponse struct {
	Response *area.SaveReply `json:"response"`
	Err      error           `json:"err"`
}

// MakeSaveEndpoint returns an endpoint that invokes Save on the service.
func MakeSaveEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveRequest)
		response, err := s.Save(ctx, req.Request)
		return SaveResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r SaveResponse) Failed() error {
	return r.Err
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Request *area.GetRequest `json:"request"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Response *area.GetReply `json:"response"`
	Err      error          `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		response, err := s.Get(ctx, req.Request)
		return GetResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// GetByIdRequest collects the request parameters for the GetById method.
type GetByIdRequest struct {
	Request *area.GetByIdRequest `json:"request"`
}

// GetByIdResponse collects the response parameters for the GetById method.
type GetByIdResponse struct {
	Response *area.GetByIdReply `json:"response"`
	Err      error              `json:"err"`
}

// MakeGetByIdEndpoint returns an endpoint that invokes GetById on the service.
func MakeGetByIdEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIdRequest)
		response, err := s.GetById(ctx, req.Request)
		return GetByIdResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIdResponse) Failed() error {
	return r.Err
}

// GetByCodeRequest collects the request parameters for the GetByCode method.
type GetByCodeRequest struct {
	Request *area.GetByCodeRequest `json:"request"`
}

// GetByCodeResponse collects the response parameters for the GetByCode method.
type GetByCodeResponse struct {
	Response *area.GetByCodeReply `json:"response"`
	Err      error                `json:"err"`
}

// MakeGetByCodeEndpoint returns an endpoint that invokes GetByCode on the service.
func MakeGetByCodeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByCodeRequest)
		response, err := s.GetByCode(ctx, req.Request)
		return GetByCodeResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByCodeResponse) Failed() error {
	return r.Err
}

// GetByNameRequest collects the request parameters for the GetByName method.
type GetByNameRequest struct {
	Request *area.GetByNameRequest `json:"request"`
}

// GetByNameResponse collects the response parameters for the GetByName method.
type GetByNameResponse struct {
	Response *area.GetByNameReply `json:"response"`
	Err      error                `json:"err"`
}

// MakeGetByNameEndpoint returns an endpoint that invokes GetByName on the service.
func MakeGetByNameEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByNameRequest)
		response, err := s.GetByName(ctx, req.Request)
		return GetByNameResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByNameResponse) Failed() error {
	return r.Err
}

// GetByTimeZoneRequest collects the request parameters for the GetByTimeZone method.
type GetByTimeZoneRequest struct {
	Request *area.GetByTimeZoneRequest `json:"request"`
}

// GetByTimeZoneResponse collects the response parameters for the GetByTimeZone method.
type GetByTimeZoneResponse struct {
	Response *area.GetByTimeZoneReply `json:"response"`
	Err      error                    `json:"err"`
}

// MakeGetByTimeZoneEndpoint returns an endpoint that invokes GetByTimeZone on the service.
func MakeGetByTimeZoneEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByTimeZoneRequest)
		response, err := s.GetByTimeZone(ctx, req.Request)
		return GetByTimeZoneResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByTimeZoneResponse) Failed() error {
	return r.Err
}

// GetByLanguageRequest collects the request parameters for the GetByLanguage method.
type GetByLanguageRequest struct {
	Request *area.GetByLanguageRequest `json:"request"`
}

// GetByLanguageResponse collects the response parameters for the GetByLanguage method.
type GetByLanguageResponse struct {
	Response *area.GetByLanguageReply `json:"response"`
	Err      error                    `json:"err"`
}

// MakeGetByLanguageEndpoint returns an endpoint that invokes GetByLanguage on the service.
func MakeGetByLanguageEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByLanguageRequest)
		response, err := s.GetByLanguage(ctx, req.Request)
		return GetByLanguageResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByLanguageResponse) Failed() error {
	return r.Err
}

// GetByStateRequest collects the request parameters for the GetByState method.
type GetByStateRequest struct {
	Request *area.GetByStateRequest `json:"request"`
}

// GetByStateResponse collects the response parameters for the GetByState method.
type GetByStateResponse struct {
	Response *area.GetByStateReply `json:"response"`
	Err      error                 `json:"err"`
}

// MakeGetByStateEndpoint returns an endpoint that invokes GetByState on the service.
func MakeGetByStateEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByStateRequest)
		response, err := s.GetByState(ctx, req.Request)
		return GetByStateResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByStateResponse) Failed() error {
	return r.Err
}

// GetByCreateTimeRequest collects the request parameters for the GetByCreateTime method.
type GetByCreateTimeRequest struct {
	Request *area.GetByCreateTimeRequest `json:"request"`
}

// GetByCreateTimeResponse collects the response parameters for the GetByCreateTime method.
type GetByCreateTimeResponse struct {
	Response *area.GetByCreateTimeReply `json:"response"`
	Err      error                      `json:"err"`
}

// MakeGetByCreateTimeEndpoint returns an endpoint that invokes GetByCreateTime on the service.
func MakeGetByCreateTimeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByCreateTimeRequest)
		response, err := s.GetByCreateTime(ctx, req.Request)
		return GetByCreateTimeResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByCreateTimeResponse) Failed() error {
	return r.Err
}

// GetByUpdateTimeRequest collects the request parameters for the GetByUpdateTime method.
type GetByUpdateTimeRequest struct {
	Request *area.GetByUpdateTimeRequest `json:"request"`
}

// GetByUpdateTimeResponse collects the response parameters for the GetByUpdateTime method.
type GetByUpdateTimeResponse struct {
	Response *area.GetByUpdateTimeReply `json:"response"`
	Err      error                      `json:"err"`
}

// MakeGetByUpdateTimeEndpoint returns an endpoint that invokes GetByUpdateTime on the service.
func MakeGetByUpdateTimeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByUpdateTimeRequest)
		response, err := s.GetByUpdateTime(ctx, req.Request)
		return GetByUpdateTimeResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByUpdateTimeResponse) Failed() error {
	return r.Err
}

// GetByDeleteTimeRequest collects the request parameters for the GetByDeleteTime method.
type GetByDeleteTimeRequest struct {
	Request *area.GetByDeleteTimeRequest `json:"request"`
}

// GetByDeleteTimeResponse collects the response parameters for the GetByDeleteTime method.
type GetByDeleteTimeResponse struct {
	Rsponse *area.GetByDeleteTimeReply `json:"rsponse"`
	Err     error                      `json:"err"`
}

// MakeGetByDeleteTimeEndpoint returns an endpoint that invokes GetByDeleteTime on the service.
func MakeGetByDeleteTimeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByDeleteTimeRequest)
		rsponse, err := s.GetByDeleteTime(ctx, req.Request)
		return GetByDeleteTimeResponse{
			Err:     err,
			Rsponse: rsponse,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByDeleteTimeResponse) Failed() error {
	return r.Err
}

// UpdateCodeRequest collects the request parameters for the UpdateCode method.
type UpdateCodeRequest struct {
	Request *area.UpdateCodeRequest `json:"request"`
}

// UpdateCodeResponse collects the response parameters for the UpdateCode method.
type UpdateCodeResponse struct {
	Response *area.UpdateCodeReply `json:"response"`
	Err      error                 `json:"err"`
}

// MakeUpdateCodeEndpoint returns an endpoint that invokes UpdateCode on the service.
func MakeUpdateCodeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCodeRequest)
		response, err := s.UpdateCode(ctx, req.Request)
		return UpdateCodeResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateCodeResponse) Failed() error {
	return r.Err
}

// UpdateNameRequest collects the request parameters for the UpdateName method.
type UpdateNameRequest struct {
	Request *area.UpdateNameRequest `json:"request"`
}

// UpdateNameResponse collects the response parameters for the UpdateName method.
type UpdateNameResponse struct {
	Response *area.UpdateNameReply `json:"response"`
	Err      error                 `json:"err"`
}

// MakeUpdateNameEndpoint returns an endpoint that invokes UpdateName on the service.
func MakeUpdateNameEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateNameRequest)
		response, err := s.UpdateName(ctx, req.Request)
		return UpdateNameResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateNameResponse) Failed() error {
	return r.Err
}

// UpdateTimeZoneRequest collects the request parameters for the UpdateTimeZone method.
type UpdateTimeZoneRequest struct {
	Request *area.UpdateTimeZoneRequest `json:"request"`
}

// UpdateTimeZoneResponse collects the response parameters for the UpdateTimeZone method.
type UpdateTimeZoneResponse struct {
	Response *area.UpdateTimeZoneReply `json:"response"`
	Err      error                     `json:"err"`
}

// MakeUpdateTimeZoneEndpoint returns an endpoint that invokes UpdateTimeZone on the service.
func MakeUpdateTimeZoneEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTimeZoneRequest)
		response, err := s.UpdateTimeZone(ctx, req.Request)
		return UpdateTimeZoneResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTimeZoneResponse) Failed() error {
	return r.Err
}

// UpdateLanguageRequest collects the request parameters for the UpdateLanguage method.
type UpdateLanguageRequest struct {
	Request *area.UpdateLanguageRequest `json:"request"`
}

// UpdateLanguageResponse collects the response parameters for the UpdateLanguage method.
type UpdateLanguageResponse struct {
	Response *area.UpdateLanguageReply `json:"response"`
	Err      error                     `json:"err"`
}

// MakeUpdateLanguageEndpoint returns an endpoint that invokes UpdateLanguage on the service.
func MakeUpdateLanguageEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateLanguageRequest)
		response, err := s.UpdateLanguage(ctx, req.Request)
		return UpdateLanguageResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateLanguageResponse) Failed() error {
	return r.Err
}

// UpdateStateRequest collects the request parameters for the UpdateState method.
type UpdateStateRequest struct {
	Request *area.UpdateStateRequest `json:"request"`
}

// UpdateStateResponse collects the response parameters for the UpdateState method.
type UpdateStateResponse struct {
	Response *area.UpdateStateReply `json:"response"`
	Err      error                  `json:"err"`
}

// MakeUpdateStateEndpoint returns an endpoint that invokes UpdateState on the service.
func MakeUpdateStateEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateStateRequest)
		response, err := s.UpdateState(ctx, req.Request)
		return UpdateStateResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateStateResponse) Failed() error {
	return r.Err
}

// DeleteByIdRequest collects the request parameters for the DeleteById method.
type DeleteByIdRequest struct {
	Request *area.DeleteByIdRequest `json:"request"`
}

// DeleteByIdResponse collects the response parameters for the DeleteById method.
type DeleteByIdResponse struct {
	Response *area.DeleteByIdReply `json:"response"`
	Err      error                 `json:"err"`
}

// MakeDeleteByIdEndpoint returns an endpoint that invokes DeleteById on the service.
func MakeDeleteByIdEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByIdRequest)
		response, err := s.DeleteById(ctx, req.Request)
		return DeleteByIdResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteByIdResponse) Failed() error {
	return r.Err
}

// DeleteByCodeRequest collects the request parameters for the DeleteByCode method.
type DeleteByCodeRequest struct {
	Request *area.DeleteByCodeRequest `json:"request"`
}

// DeleteByCodeResponse collects the response parameters for the DeleteByCode method.
type DeleteByCodeResponse struct {
	Response *area.DeleteByCodeReply `json:"response"`
	Err      error                   `json:"err"`
}

// MakeDeleteByCodeEndpoint returns an endpoint that invokes DeleteByCode on the service.
func MakeDeleteByCodeEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByCodeRequest)
		response, err := s.DeleteByCode(ctx, req.Request)
		return DeleteByCodeResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteByCodeResponse) Failed() error {
	return r.Err
}

// DeleteByNameRequest collects the request parameters for the DeleteByName method.
type DeleteByNameRequest struct {
	Request *area.DeleteByNameRequest `json:"request"`
}

// DeleteByNameResponse collects the response parameters for the DeleteByName method.
type DeleteByNameResponse struct {
	Response *area.DeleteByNameReply `json:"response"`
	Err      error                   `json:"err"`
}

// MakeDeleteByNameEndpoint returns an endpoint that invokes DeleteByName on the service.
func MakeDeleteByNameEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByNameRequest)
		response, err := s.DeleteByName(ctx, req.Request)
		return DeleteByNameResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteByNameResponse) Failed() error {
	return r.Err
}

// DeleteByStateRequest collects the request parameters for the DeleteByState method.
type DeleteByStateRequest struct {
	Request *area.DeleteByStateRequest `json:"request"`
}

// DeleteByStateResponse collects the response parameters for the DeleteByState method.
type DeleteByStateResponse struct {
	Response *area.DeleteByStateReply `json:"response"`
	Err      error                    `json:"err"`
}

// MakeDeleteByStateEndpoint returns an endpoint that invokes DeleteByState on the service.
func MakeDeleteByStateEndpoint(s service.AreaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByStateRequest)
		response, err := s.DeleteByState(ctx, req.Request)
		return DeleteByStateResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteByStateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Save implements Service. Primarily useful in a client.
func (e Endpoints) Save(ctx context.Context, request *area.SaveRequest) (response *area.SaveReply, err error) {
	request0 := SaveRequest{Request: request}
	response1, err := e.SaveEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(SaveResponse).Response, response1.(SaveResponse).Err
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, request *area.GetRequest) (response *area.GetReply, err error) {
	request0 := GetRequest{Request: request}
	response1, err := e.GetEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetResponse).Response, response1.(GetResponse).Err
}

// GetById implements Service. Primarily useful in a client.
func (e Endpoints) GetById(ctx context.Context, request *area.GetByIdRequest) (response *area.GetByIdReply, err error) {
	request0 := GetByIdRequest{Request: request}
	response1, err := e.GetByIdEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByIdResponse).Response, response1.(GetByIdResponse).Err
}

// GetByCode implements Service. Primarily useful in a client.
func (e Endpoints) GetByCode(ctx context.Context, request *area.GetByCodeRequest) (response *area.GetByCodeReply, err error) {
	request0 := GetByCodeRequest{Request: request}
	response1, err := e.GetByCodeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByCodeResponse).Response, response1.(GetByCodeResponse).Err
}

// GetByName implements Service. Primarily useful in a client.
func (e Endpoints) GetByName(ctx context.Context, request *area.GetByNameRequest) (response *area.GetByNameReply, err error) {
	request0 := GetByNameRequest{Request: request}
	response1, err := e.GetByNameEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByNameResponse).Response, response1.(GetByNameResponse).Err
}

// GetByTimeZone implements Service. Primarily useful in a client.
func (e Endpoints) GetByTimeZone(ctx context.Context, request *area.GetByTimeZoneRequest) (response *area.GetByTimeZoneReply, err error) {
	request0 := GetByTimeZoneRequest{Request: request}
	response1, err := e.GetByTimeZoneEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByTimeZoneResponse).Response, response1.(GetByTimeZoneResponse).Err
}

// GetByLanguage implements Service. Primarily useful in a client.
func (e Endpoints) GetByLanguage(ctx context.Context, request *area.GetByLanguageRequest) (response *area.GetByLanguageReply, err error) {
	request0 := GetByLanguageRequest{Request: request}
	response1, err := e.GetByLanguageEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByLanguageResponse).Response, response1.(GetByLanguageResponse).Err
}

// GetByState implements Service. Primarily useful in a client.
func (e Endpoints) GetByState(ctx context.Context, request *area.GetByStateRequest) (response *area.GetByStateReply, err error) {
	request0 := GetByStateRequest{Request: request}
	response1, err := e.GetByStateEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByStateResponse).Response, response1.(GetByStateResponse).Err
}

// GetByCreateTime implements Service. Primarily useful in a client.
func (e Endpoints) GetByCreateTime(ctx context.Context, request *area.GetByCreateTimeRequest) (response *area.GetByCreateTimeReply, err error) {
	request0 := GetByCreateTimeRequest{Request: request}
	response1, err := e.GetByCreateTimeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByCreateTimeResponse).Response, response1.(GetByCreateTimeResponse).Err
}

// GetByUpdateTime implements Service. Primarily useful in a client.
func (e Endpoints) GetByUpdateTime(ctx context.Context, request *area.GetByUpdateTimeRequest) (response *area.GetByUpdateTimeReply, err error) {
	request0 := GetByUpdateTimeRequest{Request: request}
	response1, err := e.GetByUpdateTimeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(GetByUpdateTimeResponse).Response, response1.(GetByUpdateTimeResponse).Err
}

// GetByDeleteTime implements Service. Primarily useful in a client.
func (e Endpoints) GetByDeleteTime(ctx context.Context, request *area.GetByDeleteTimeRequest) (rsponse *area.GetByDeleteTimeReply, err error) {
	request0 := GetByDeleteTimeRequest{Request: request}
	response, err := e.GetByDeleteTimeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(GetByDeleteTimeResponse).Rsponse, response.(GetByDeleteTimeResponse).Err
}

// UpdateCode implements Service. Primarily useful in a client.
func (e Endpoints) UpdateCode(ctx context.Context, request *area.UpdateCodeRequest) (response *area.UpdateCodeReply, err error) {
	request0 := UpdateCodeRequest{Request: request}
	response1, err := e.UpdateCodeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(UpdateCodeResponse).Response, response1.(UpdateCodeResponse).Err
}

// UpdateName implements Service. Primarily useful in a client.
func (e Endpoints) UpdateName(ctx context.Context, request *area.UpdateNameRequest) (response *area.UpdateNameReply, err error) {
	request0 := UpdateNameRequest{Request: request}
	response1, err := e.UpdateNameEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(UpdateNameResponse).Response, response1.(UpdateNameResponse).Err
}

// UpdateTimeZone implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTimeZone(ctx context.Context, request *area.UpdateTimeZoneRequest) (response *area.UpdateTimeZoneReply, err error) {
	request0 := UpdateTimeZoneRequest{Request: request}
	response1, err := e.UpdateTimeZoneEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(UpdateTimeZoneResponse).Response, response1.(UpdateTimeZoneResponse).Err
}

// UpdateLanguage implements Service. Primarily useful in a client.
func (e Endpoints) UpdateLanguage(ctx context.Context, request *area.UpdateLanguageRequest) (response *area.UpdateLanguageReply, err error) {
	request0 := UpdateLanguageRequest{Request: request}
	response1, err := e.UpdateLanguageEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(UpdateLanguageResponse).Response, response1.(UpdateLanguageResponse).Err
}

// UpdateState implements Service. Primarily useful in a client.
func (e Endpoints) UpdateState(ctx context.Context, request *area.UpdateStateRequest) (response *area.UpdateStateReply, err error) {
	request0 := UpdateStateRequest{Request: request}
	response1, err := e.UpdateStateEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(UpdateStateResponse).Response, response1.(UpdateStateResponse).Err
}

// DeleteById implements Service. Primarily useful in a client.
func (e Endpoints) DeleteById(ctx context.Context, request *area.DeleteByIdRequest) (response *area.DeleteByIdReply, err error) {
	request0 := DeleteByIdRequest{Request: request}
	response1, err := e.DeleteByIdEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(DeleteByIdResponse).Response, response1.(DeleteByIdResponse).Err
}

// DeleteByCode implements Service. Primarily useful in a client.
func (e Endpoints) DeleteByCode(ctx context.Context, request *area.DeleteByCodeRequest) (response *area.DeleteByCodeReply, err error) {
	request0 := DeleteByCodeRequest{Request: request}
	response1, err := e.DeleteByCodeEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(DeleteByCodeResponse).Response, response1.(DeleteByCodeResponse).Err
}

// DeleteByName implements Service. Primarily useful in a client.
func (e Endpoints) DeleteByName(ctx context.Context, request *area.DeleteByNameRequest) (response *area.DeleteByNameReply, err error) {
	request0 := DeleteByNameRequest{Request: request}
	response1, err := e.DeleteByNameEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(DeleteByNameResponse).Response, response1.(DeleteByNameResponse).Err
}

// DeleteByState implements Service. Primarily useful in a client.
func (e Endpoints) DeleteByState(ctx context.Context, request *area.DeleteByStateRequest) (response *area.DeleteByStateReply, err error) {
	request0 := DeleteByStateRequest{Request: request}
	response1, err := e.DeleteByStateEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response1.(DeleteByStateResponse).Response, response1.(DeleteByStateResponse).Err
}
