package transport

import (
	"context"

	"github.com/david1992121/veritrans-microservice/api/pb"
	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg/endpoint"
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
	pb.UnimplementedVeritransServer
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
	accountParam.AccountID = req.AccountId
	if req.CardParam != nil {
		if req.CardParam.CardNumber != "" {
			accountParam.CardParam = &veritrans.CardParam{
				CardNumber: req.CardParam.CardNumber,
				CardExpire: req.CardParam.CardExpire,
			}
			if req.CardParam.DefaultCard != nil && *req.CardParam.DefaultCard != "" {
				accountParam.CardParam.DefaultCard = *req.CardParam.DefaultCard
			}
		}
	}
	return accountParam, nil
}

func encodeAccountResponse(_ context.Context, endpointRes interface{}) (interface{}, error) {
	res := endpointRes.(endpoint.AccountResponse)
	var accountReply pb.AccountReply
	accountReply.Account = &pb.AccountReply_AccountInfo{
		AccountId: res.Account.AccountID,
	}
	for _, cardItem := range res.Account.CardInfo {
		accountReply.Account.CardInfo = append(accountReply.Account.CardInfo, &pb.AccountReply_AccountInfo_CardInfo{
			CardId:      cardItem.CardID,
			CardExpire:  cardItem.CardExpire,
			CardNumber:  cardItem.CardNumber,
			DefaultCard: cardItem.DefaultCard,
		})
	}
	return &accountReply, nil
}

func decodeGRPCPaymentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PaymentRequest)
	var param veritrans.Params
	param.OrderID = req.OrderId
	param.Amount = req.Amount
	if req.Jpo != nil && *req.Jpo != "" {
		param.JPO = *req.Jpo
	}
	if req.PayNowIdParam != nil {
		if req.PayNowIdParam.Token != "" {
			param.PayNowIDParam.Token = req.PayNowIdParam.Token
		}
		if req.PayNowIdParam.AccountParam != nil && req.PayNowIdParam.AccountParam.AccountId != "" {
			param.PayNowIDParam.AccountParam = &veritrans.AccountParam{
				AccountID: req.PayNowIdParam.AccountParam.AccountId,
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
