package pkg

import (
	"context"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
)

// ServiceConfig struct
type ServiceConfig struct {
	MDKConfig        veritrans.MDKConfig
	ConnectionConfig veritrans.ConnectionConfig
}

type veritransService struct {
	MDKService     *veritrans.MDKService
	AccountService *veritrans.AccountService
	PaymentService *veritrans.PaymentService
}

// NewService initializes the veritrans service
func NewService(config *ServiceConfig) Service {
	mdkService := veritrans.NewMDKService(config.MDKConfig)

	paymentService, _ := veritrans.NewPaymentService(config.ConnectionConfig)
	accountService := veritrans.NewAccountService(config.ConnectionConfig)
	return &veritransService{
		MDKService:     mdkService,
		AccountService: accountService,
		PaymentService: paymentService,
	}
}

func (v *veritransService) GetMDKToken(_ context.Context, cardInfo *veritrans.ClientCardInfo) (string, error) {
	return v.MDKService.GetCardToken(cardInfo)
}

func (v *veritransService) CreateAccount(_ context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error) {
	return v.AccountService.CreateAccount(accountParam)
}

func (v *veritransService) UpdateAccount(_ context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error) {
	return v.AccountService.UpdateAccount(accountParam)
}

func (v *veritransService) Authorize(_ context.Context, param *veritrans.Params) error {
	_, err := v.PaymentService.Authorize(param, veritrans.PaymentServiceType(veritrans.PayCard))
	return err
}

func (v *veritransService) Cancel(_ context.Context, param *veritrans.Params) error {
	_, err := v.PaymentService.Cancel(param, veritrans.PaymentServiceType(veritrans.PayCard))
	return err
}
