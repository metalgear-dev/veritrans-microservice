package pkg

import (
	"context"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
)

// Service of the veritrans payment
type Service interface {
	// GetMDKToken function gets the MDK token from card information
	GetMDKToken(ctx context.Context, cardInfo *veritrans.ClientCardInfo) (string, error)
	// CreateAccount function creates a veritrans account
	CreateAccount(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// UpdateAccount function updates the veritrans account
	UpdateAccount(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// Authorize function executes the veritrans payment
	Authorize(ctx context.Context, param *veritrans.Params) error
	// Cancel function cancels the veritrans payment
	Cancel(ctx context.Context, param *veritrans.Params) error
}
