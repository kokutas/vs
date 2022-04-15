package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/kokutas/vs/area/pkg/endpoint"
	"net/http"
)

// makeSaveHandler creates the handler logic
func makeSaveHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/save", http1.NewServer(endpoints.SaveEndpoint, decodeSaveRequest, encodeSaveResponse, options...))
}

// decodeSaveRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSaveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SaveRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSaveResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSaveResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetHandler creates the handler logic
func makeGetHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get", http1.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...))
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIdHandler creates the handler logic
func makeGetByIdHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-id", http1.NewServer(endpoints.GetByIdEndpoint, decodeGetByIdRequest, encodeGetByIdResponse, options...))
}

// decodeGetByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIdResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByCodeHandler creates the handler logic
func makeGetByCodeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-code", http1.NewServer(endpoints.GetByCodeEndpoint, decodeGetByCodeRequest, encodeGetByCodeResponse, options...))
}

// decodeGetByCodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByCodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByCodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByCodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByCodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByNameHandler creates the handler logic
func makeGetByNameHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-name", http1.NewServer(endpoints.GetByNameEndpoint, decodeGetByNameRequest, encodeGetByNameResponse, options...))
}

// decodeGetByNameRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByNameRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByNameResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByTimeZoneHandler creates the handler logic
func makeGetByTimeZoneHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-time-zone", http1.NewServer(endpoints.GetByTimeZoneEndpoint, decodeGetByTimeZoneRequest, encodeGetByTimeZoneResponse, options...))
}

// decodeGetByTimeZoneRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByTimeZoneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByTimeZoneRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByTimeZoneResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByTimeZoneResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByLanguageHandler creates the handler logic
func makeGetByLanguageHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-language", http1.NewServer(endpoints.GetByLanguageEndpoint, decodeGetByLanguageRequest, encodeGetByLanguageResponse, options...))
}

// decodeGetByLanguageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByLanguageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByLanguageRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByLanguageResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByLanguageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByStateHandler creates the handler logic
func makeGetByStateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-state", http1.NewServer(endpoints.GetByStateEndpoint, decodeGetByStateRequest, encodeGetByStateResponse, options...))
}

// decodeGetByStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByStateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByStateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByStateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByCreateTimeHandler creates the handler logic
func makeGetByCreateTimeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-create-time", http1.NewServer(endpoints.GetByCreateTimeEndpoint, decodeGetByCreateTimeRequest, encodeGetByCreateTimeResponse, options...))
}

// decodeGetByCreateTimeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByCreateTimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByCreateTimeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByCreateTimeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByCreateTimeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByUpdateTimeHandler creates the handler logic
func makeGetByUpdateTimeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-update-time", http1.NewServer(endpoints.GetByUpdateTimeEndpoint, decodeGetByUpdateTimeRequest, encodeGetByUpdateTimeResponse, options...))
}

// decodeGetByUpdateTimeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByUpdateTimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByUpdateTimeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByUpdateTimeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByUpdateTimeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByDeleteTimeHandler creates the handler logic
func makeGetByDeleteTimeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-by-delete-time", http1.NewServer(endpoints.GetByDeleteTimeEndpoint, decodeGetByDeleteTimeRequest, encodeGetByDeleteTimeResponse, options...))
}

// decodeGetByDeleteTimeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByDeleteTimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetByDeleteTimeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByDeleteTimeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByDeleteTimeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateCodeHandler creates the handler logic
func makeUpdateCodeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-code", http1.NewServer(endpoints.UpdateCodeEndpoint, decodeUpdateCodeRequest, encodeUpdateCodeResponse, options...))
}

// decodeUpdateCodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateCodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateCodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateCodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateCodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateNameHandler creates the handler logic
func makeUpdateNameHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-name", http1.NewServer(endpoints.UpdateNameEndpoint, decodeUpdateNameRequest, encodeUpdateNameResponse, options...))
}

// decodeUpdateNameRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateNameRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateNameResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTimeZoneHandler creates the handler logic
func makeUpdateTimeZoneHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-time-zone", http1.NewServer(endpoints.UpdateTimeZoneEndpoint, decodeUpdateTimeZoneRequest, encodeUpdateTimeZoneResponse, options...))
}

// decodeUpdateTimeZoneRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTimeZoneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTimeZoneRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTimeZoneResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTimeZoneResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateLanguageHandler creates the handler logic
func makeUpdateLanguageHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-language", http1.NewServer(endpoints.UpdateLanguageEndpoint, decodeUpdateLanguageRequest, encodeUpdateLanguageResponse, options...))
}

// decodeUpdateLanguageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateLanguageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateLanguageRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateLanguageResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateLanguageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateStateHandler creates the handler logic
func makeUpdateStateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-state", http1.NewServer(endpoints.UpdateStateEndpoint, decodeUpdateStateRequest, encodeUpdateStateResponse, options...))
}

// decodeUpdateStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateStateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateStateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateStateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteByIdHandler creates the handler logic
func makeDeleteByIdHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-by-id", http1.NewServer(endpoints.DeleteByIdEndpoint, decodeDeleteByIdRequest, encodeDeleteByIdResponse, options...))
}

// decodeDeleteByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteByIdResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteByCodeHandler creates the handler logic
func makeDeleteByCodeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-by-code", http1.NewServer(endpoints.DeleteByCodeEndpoint, decodeDeleteByCodeRequest, encodeDeleteByCodeResponse, options...))
}

// decodeDeleteByCodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteByCodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteByCodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteByCodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteByCodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteByNameHandler creates the handler logic
func makeDeleteByNameHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-by-name", http1.NewServer(endpoints.DeleteByNameEndpoint, decodeDeleteByNameRequest, encodeDeleteByNameResponse, options...))
}

// decodeDeleteByNameRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteByNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteByNameRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteByNameResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteByNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteByStateHandler creates the handler logic
func makeDeleteByStateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-by-state", http1.NewServer(endpoints.DeleteByStateEndpoint, decodeDeleteByStateRequest, encodeDeleteByStateResponse, options...))
}

// decodeDeleteByStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteByStateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteByStateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteByStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteByStateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
