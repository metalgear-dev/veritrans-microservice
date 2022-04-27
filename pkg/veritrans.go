package pkg

import (
	"context"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
)

type veritransService struct{}

// NewService initializes the veritrans service
func NewService() Service {
	return &veritransService{}
}

func (v *veritransService) GetMDKToken(_ context.Context, cardInfo *veritrans.ClientCardInfo) (string, error) {
	// Get MDK token from the card information
	return "token", nil
}

func (v *veritransService) CreateAccount(_ context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error) {
	// Create the veritrans account using the account information
	return nil, nil
}

func (v *veritransService) UpdateAccount(_ context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error) {
	// Update the veritrans account
	return nil, nil
}

func (v *veritransService) Authorize(_ context.Context, param *veritrans.Params) error {
	// Execute the payment using veritrans account or MDK token
	return nil
}

func (v *veritransService) Cancel(_ context.Context, param *veritrans.Params) error {
	// Cancel the payment
	return nil
}
