package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg"
	"github.com/david1992121/veritrans-microservice/pkg/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

// NewHTTPHandler returns the handler
func NewHTTPHandler(logger log.Logger) http.Handler {
	service := pkg.NewLoggingMiddleware(logger, pkg.NewService(pkg.GetServiceConfig()))
	eps := endpoint.NewEndpointSet(service)
	return GetHTTPHandler(eps)
}

// GetHTTPHandler initializes the http handler
func GetHTTPHandler(ep endpoint.Set) http.Handler {
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

	m.Handle("/card/create", httptransport.NewServer(
		ep.CreateCardEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/card/update", httptransport.NewServer(
		ep.UpdateCardEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/card/delete", httptransport.NewServer(
		ep.DeleteCardEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/card/get", httptransport.NewServer(
		ep.GetCardEndpoint,
		decodeHTTPAccountRequest,
		encodeResponse,
	))

	m.Handle("/authorize", httptransport.NewServer(
		ep.AuthorizeEndpoint,
		decodeHTTPPaymentRequest,
		encodeResponse,
	))

	m.Handle("/capture", httptransport.NewServer(
		ep.CaptureEndpoint,
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
