package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getMDKToken   grpctransport.Handler
	createAccount grpctransport.Handler
	updateAccount grpctransport.Handler
	createCard    grpctransport.Handler
	updateCard    grpctransport.Handler
	deleteCard    grpctransport.Handler
	getCard       grpctransport.Handler
	authorize     grpctransport.Handler
	capture       grpctransport.Handler
	cancel        grpctransport.Handler
}
