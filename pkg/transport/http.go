package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler initializes the http handler
func NewHTTPHandler(ep endpoint.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/mdk/token", httptransport.NewServer(
		ep.GetMDKTokenEndpoint,
		decodeHTTPGetMDKTokenRequest,
		encodeResponse,
	))

	m.Handle("/account/create", httptransport.NewServer(
		ep.CreateAccountEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/account/update", httptransport.NewServer(
		ep.UpdateAccountEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/authorize", httptransport.NewServer(
		ep.AuthorizeEndpoint,
		decodeHTTPPaymentRequest,
		encodeResponse,
	))

	m.Handle("/cancel", httptransport.NewServer(
		ep.CancelEndpoint,
		decodeHTTPPaymentRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGetMDKTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req veritrans.ClientCardInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req veritrans.AccountParam
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req veritrans.Params
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
