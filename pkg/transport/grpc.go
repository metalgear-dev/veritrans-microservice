package transport

import (
	"context"

	"github.com/david1992121/veritrans-microservice/api/pb"
	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg"
	"github.com/david1992121/veritrans-microservice/pkg/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
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
	pb.UnimplementedVeritransServer
}

// GetGRPCServer returns the handler
func GetGRPCServer(logger log.Logger) pb.VeritransServer {
	service := pkg.NewLoggingMiddleware(logger, pkg.NewService(pkg.GetServiceConfig()))
	eps := endpoint.NewEndpointSet(service)
	return NewGRPCServer(eps)
}

// NewGRPCServer function intializes a new gRPC server
func NewGRPCServer(ep endpoint.Set) pb.VeritransServer {
	return &grpcServer{
		getMDKToken: grpctransport.NewServer(
			ep.GetMDKTokenEndpoint,
			decodeGRPCMDKRequest,
			encodeMDKResponse,
		),
		createAccount: grpctransport.NewServer(
			ep.CreateAccountEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		updateAccount: grpctransport.NewServer(
			ep.UpdateAccountEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		createCard: grpctransport.NewServer(
			ep.CreateCardEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		updateCard: grpctransport.NewServer(
			ep.UpdateCardEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		deleteCard: grpctransport.NewServer(
			ep.DeleteCardEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		getCard: grpctransport.NewServer(
			ep.GetCardEndpoint,
			decodeGRPCAccountRequest,
			encodeAccountResponse,
		),
		authorize: grpctransport.NewServer(
			ep.AuthorizeEndpoint,
			decodeGRPCPaymentRequest,
			encodePaymentResponse,
		),
		capture: grpctransport.NewServer(
			ep.CaptureEndpoint,
			decodeGRPCPaymentRequest,
			encodePaymentResponse,
		),
		cancel: grpctransport.NewServer(
			ep.CancelEndpoint,
			decodeGRPCPaymentRequest,
			encodePaymentResponse,
		),
	}
}

func (g *grpcServer) GetMDKToken(ctx context.Context, r *pb.GetMDKTokenRequest) (*pb.TokenReply, error) {
	_, rep, err := g.getMDKToken.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TokenReply), nil
}

func (g *grpcServer) CreateAccount(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.createAccount.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) UpdateAccount(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.updateAccount.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) CreateCard(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.createCard.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) UpdateCard(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.updateCard.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) DeleteCard(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.deleteCard.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) GetCard(ctx context.Context, r *pb.AccountRequest) (*pb.AccountReply, error) {
	_, rep, err := g.getCard.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AccountReply), nil
}

func (g *grpcServer) Authorize(ctx context.Context, r *pb.PaymentRequest) (*pb.PaymentReply, error) {
	_, rep, err := g.authorize.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PaymentReply), nil
}

func (g *grpcServer) Capture(ctx context.Context, r *pb.PaymentRequest) (*pb.PaymentReply, error) {
	_, rep, err := g.capture.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PaymentReply), nil
}

func (g *grpcServer) Cancel(ctx context.Context, r *pb.PaymentRequest) (*pb.PaymentReply, error) {
	_, rep, err := g.cancel.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PaymentReply), nil
}

func decodeGRPCMDKRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetMDKTokenRequest)
	clientCardInfo := veritrans.ClientCardInfo{
		CardNumber:   req.CardNumber,
		CardExpire:   req.CardExpire,
		SecurityCode: req.SecurityCode,
	}
	if req.CardHolderName != nil && *req.CardHolderName != "" {
		clientCardInfo.CardHolderName = *req.CardHolderName
	}
	return clientCardInfo, nil
}

func encodeMDKResponse(_ context.Context, endpointRes interface{}) (interface{}, error) {
	res := endpointRes.(endpoint.GetMDKTokenResponse)
	return &pb.TokenReply{Token: res.Token, Err: res.Err}, nil
}

func decodeGRPCAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AccountRequest)
	var accountParam veritrans.AccountParam
	accountParam.AccountID = req.AccountID
	if req.CardParam != nil {
		accountParam.CardParam = &veritrans.CardParam{}
		if req.CardParam.CardNumber != nil && *req.CardParam.CardNumber != "" {
			accountParam.CardParam.CardNumber = *(req.CardParam.CardNumber)
		}
		if req.CardParam.CardExpire != nil && *req.CardParam.CardExpire != "" {
			accountParam.CardParam.CardExpire = *req.CardParam.CardExpire
		}
		if req.CardParam.DefaultCard != nil && *req.CardParam.DefaultCard != "" {
			accountParam.CardParam.DefaultCard = *req.CardParam.DefaultCard
		}
		if req.CardParam.CardID != nil && *req.CardParam.CardID != "" {
			accountParam.CardParam.CardID = *req.CardParam.CardID
		}
	}
	return accountParam, nil
}

func encodeAccountResponse(_ context.Context, endpointRes interface{}) (interface{}, error) {
	res := endpointRes.(endpoint.AccountResponse)
	var accountReply pb.AccountReply
	if res.Account != nil {
		accountReply.Account = &pb.AccountReply_AccountInfo{
			AccountID: res.Account.AccountID,
		}
		for _, cardItem := range res.Account.CardInfo {
			accountReply.Account.CardInfo = append(accountReply.Account.CardInfo, &pb.AccountReply_AccountInfo_CardInfo{
				CardID:      cardItem.CardID,
				CardExpire:  cardItem.CardExpire,
				CardNumber:  cardItem.CardNumber,
				DefaultCard: cardItem.DefaultCard,
			})
		}
	}
	accountReply.Err = res.Err
	return &accountReply, nil
}

func decodeGRPCPaymentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PaymentRequest)
	var param veritrans.Params
	param.OrderID = req.OrderID
	param.Amount = req.Amount
	if req.Jpo != nil && *req.Jpo != "" {
		param.JPO = *req.Jpo
	}
	if req.PayNowIDParam != nil {
		if req.PayNowIDParam.Token != "" {
			param.PayNowIDParam.Token = req.PayNowIDParam.Token
		}
		if req.PayNowIDParam.AccountParam != nil && req.PayNowIDParam.AccountParam.AccountID != "" {
			param.PayNowIDParam.AccountParam = &veritrans.AccountParam{
				AccountID: req.PayNowIDParam.AccountParam.AccountID,
			}
		}
	}
	return param, nil
}

func encodePaymentResponse(_ context.Context, endpointRes interface{}) (interface{}, error) {
	res := endpointRes.(endpoint.PaymentResponse)
	var paymentReply pb.PaymentReply
	paymentReply.Err = res.Err
	return &paymentReply, nil
}
