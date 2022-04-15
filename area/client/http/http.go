package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpoint1 "github.com/kokutas/vs/area/pkg/endpoint"
	http2 "github.com/kokutas/vs/area/pkg/http"
	service "github.com/kokutas/vs/area/pkg/service"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.AreaService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var saveEndpoint endpoint.Endpoint
	{
		saveEndpoint = http.NewClient("POST", copyURL(u, "/save"), encodeHTTPGenericRequest, decodeSaveResponse, options["Save"]...).Endpoint()
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = http.NewClient("POST", copyURL(u, "/get"), encodeHTTPGenericRequest, decodeGetResponse, options["Get"]...).Endpoint()
	}

	var getByIdEndpoint endpoint.Endpoint
	{
		getByIdEndpoint = http.NewClient("POST", copyURL(u, "/get-by-id"), encodeHTTPGenericRequest, decodeGetByIdResponse, options["GetById"]...).Endpoint()
	}

	var getByCodeEndpoint endpoint.Endpoint
	{
		getByCodeEndpoint = http.NewClient("POST", copyURL(u, "/get-by-code"), encodeHTTPGenericRequest, decodeGetByCodeResponse, options["GetByCode"]...).Endpoint()
	}

	var getByNameEndpoint endpoint.Endpoint
	{
		getByNameEndpoint = http.NewClient("POST", copyURL(u, "/get-by-name"), encodeHTTPGenericRequest, decodeGetByNameResponse, options["GetByName"]...).Endpoint()
	}

	var getByTimeZoneEndpoint endpoint.Endpoint
	{
		getByTimeZoneEndpoint = http.NewClient("POST", copyURL(u, "/get-by-time-zone"), encodeHTTPGenericRequest, decodeGetByTimeZoneResponse, options["GetByTimeZone"]...).Endpoint()
	}

	var getByLanguageEndpoint endpoint.Endpoint
	{
		getByLanguageEndpoint = http.NewClient("POST", copyURL(u, "/get-by-language"), encodeHTTPGenericRequest, decodeGetByLanguageResponse, options["GetByLanguage"]...).Endpoint()
	}

	var getByStateEndpoint endpoint.Endpoint
	{
		getByStateEndpoint = http.NewClient("POST", copyURL(u, "/get-by-state"), encodeHTTPGenericRequest, decodeGetByStateResponse, options["GetByState"]...).Endpoint()
	}

	var getByCreateTimeEndpoint endpoint.Endpoint
	{
		getByCreateTimeEndpoint = http.NewClient("POST", copyURL(u, "/get-by-create-time"), encodeHTTPGenericRequest, decodeGetByCreateTimeResponse, options["GetByCreateTime"]...).Endpoint()
	}

	var getByUpdateTimeEndpoint endpoint.Endpoint
	{
		getByUpdateTimeEndpoint = http.NewClient("POST", copyURL(u, "/get-by-update-time"), encodeHTTPGenericRequest, decodeGetByUpdateTimeResponse, options["GetByUpdateTime"]...).Endpoint()
	}

	var getByDeleteTimeEndpoint endpoint.Endpoint
	{
		getByDeleteTimeEndpoint = http.NewClient("POST", copyURL(u, "/get-by-delete-time"), encodeHTTPGenericRequest, decodeGetByDeleteTimeResponse, options["GetByDeleteTime"]...).Endpoint()
	}

	var updateCodeEndpoint endpoint.Endpoint
	{
		updateCodeEndpoint = http.NewClient("POST", copyURL(u, "/update-code"), encodeHTTPGenericRequest, decodeUpdateCodeResponse, options["UpdateCode"]...).Endpoint()
	}

	var updateNameEndpoint endpoint.Endpoint
	{
		updateNameEndpoint = http.NewClient("POST", copyURL(u, "/update-name"), encodeHTTPGenericRequest, decodeUpdateNameResponse, options["UpdateName"]...).Endpoint()
	}

	var updateTimeZoneEndpoint endpoint.Endpoint
	{
		updateTimeZoneEndpoint = http.NewClient("POST", copyURL(u, "/update-time-zone"), encodeHTTPGenericRequest, decodeUpdateTimeZoneResponse, options["UpdateTimeZone"]...).Endpoint()
	}

	var updateLanguageEndpoint endpoint.Endpoint
	{
		updateLanguageEndpoint = http.NewClient("POST", copyURL(u, "/update-language"), encodeHTTPGenericRequest, decodeUpdateLanguageResponse, options["UpdateLanguage"]...).Endpoint()
	}

	var updateStateEndpoint endpoint.Endpoint
	{
		updateStateEndpoint = http.NewClient("POST", copyURL(u, "/update-state"), encodeHTTPGenericRequest, decodeUpdateStateResponse, options["UpdateState"]...).Endpoint()
	}

	var deleteByIdEndpoint endpoint.Endpoint
	{
		deleteByIdEndpoint = http.NewClient("POST", copyURL(u, "/delete-by-id"), encodeHTTPGenericRequest, decodeDeleteByIdResponse, options["DeleteById"]...).Endpoint()
	}

	var deleteByCodeEndpoint endpoint.Endpoint
	{
		deleteByCodeEndpoint = http.NewClient("POST", copyURL(u, "/delete-by-code"), encodeHTTPGenericRequest, decodeDeleteByCodeResponse, options["DeleteByCode"]...).Endpoint()
	}

	var deleteByNameEndpoint endpoint.Endpoint
	{
		deleteByNameEndpoint = http.NewClient("POST", copyURL(u, "/delete-by-name"), encodeHTTPGenericRequest, decodeDeleteByNameResponse, options["DeleteByName"]...).Endpoint()
	}

	var deleteByStateEndpoint endpoint.Endpoint
	{
		deleteByStateEndpoint = http.NewClient("POST", copyURL(u, "/delete-by-state"), encodeHTTPGenericRequest, decodeDeleteByStateResponse, options["DeleteByState"]...).Endpoint()
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

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeSaveResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeSaveResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.SaveResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByCodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByCodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByCodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByNameResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByNameResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByNameResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByTimeZoneResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByTimeZoneResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByTimeZoneResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByLanguageResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByLanguageResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByLanguageResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByStateResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByStateResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByStateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByCreateTimeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByCreateTimeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByCreateTimeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByUpdateTimeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByUpdateTimeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByUpdateTimeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetByDeleteTimeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetByDeleteTimeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetByDeleteTimeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateCodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateCodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateCodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateNameResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateNameResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateNameResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTimeZoneResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTimeZoneResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTimeZoneResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateLanguageResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateLanguageResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateLanguageResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateStateResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateStateResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateStateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteByCodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteByCodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteByCodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteByNameResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteByNameResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteByNameResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteByStateResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteByStateResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteByStateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
