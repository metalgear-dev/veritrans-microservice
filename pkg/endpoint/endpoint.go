package endpoint

import (
	"context"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg"
	"github.com/go-kit/kit/endpoint"
)

// Set struct provides the endpoints
type Set struct {
	GetMDKTokenEndpoint   endpoint.Endpoint
	CreateAccountEndpoint endpoint.Endpoint
	UpdateAccountEndpoint endpoint.Endpoint
	AuthorizeEndpoint     endpoint.Endpoint
	CancelEndpoint        endpoint.Endpoint
}

// NewEndpointSet initializes the Set struct
func NewEndpointSet(svc pkg.Service) Set {
	return Set{
		GetMDKTokenEndpoint:   MakeGetMDKTokenEndpoint(svc),
		CreateAccountEndpoint: MakeCreateAccountEndpoint(svc),
		UpdateAccountEndpoint: MakeUpdateAccountEndpoint(svc),
		AuthorizeEndpoint:     MakeAuthorizeEndpoint(svc),
		CancelEndpoint:        MakeCancelEndpoint(svc),
	}
}

// MakeGetMDKTokenEndpoint returns the endpoint for mdk token request
func MakeGetMDKTokenEndpoint(svc pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(veritrans.ClientCardInfo)
		token, err := svc.GetMDKToken(ctx, &req)
		if err != nil {
			return GetMDKTokenResponse{Token: "", Err: err.Error()}, nil
		}
		return GetMDKTokenResponse{Token: token, Err: ""}, nil
	}
}

// MakeCreateAccountEndpoint returns the endpoint for account create request
func MakeCreateAccountEndpoint(svc pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(veritrans.AccountParam)
		account, err := svc.CreateAccount(ctx, &req)
		if err != nil {
			return account, nil
		}
		return CreateAccountResponse{Account: account, Err: ""}, nil
	}
}

// MakeUpdateAccountEndpoint returns the endpoint for acount update request
func MakeUpdateAccountEndpoint(svc pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(veritrans.AccountParam)
		account, err := svc.UpdateAccount(ctx, &req)
		if err != nil {
			return account, nil
		}
		return UpdateAccountResponse{Account: account, Err: ""}, nil
	}
}

// MakeAuthorizeEndpoint returns the endpoint for payment authorization request
func MakeAuthorizeEndpoint(svc pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(veritrans.Params)
		err := svc.Authorize(ctx, &req)
		if err != nil {
			return AuthorizeResponse{Err: err.Error()}, nil
		}
		return AuthorizeResponse{Err: ""}, nil
	}
}

// MakeCancelEndpoint returns the endpoint for payment cancel request
func MakeCancelEndpoint(svc pkg.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(veritrans.Params)
		err := svc.Cancel(ctx, &req)
		if err != nil {
			return CancelResponse{Err: err.Error()}, nil
		}
		return CancelResponse{Err: ""}, nil
	}
}
