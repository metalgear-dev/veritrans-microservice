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
	// CreateCard function adds a card into the account
	CreateCard(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// UpdateCard function adds a card into the account
	UpdateCard(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// DeleteCard function adds a card into the account
	DeleteCard(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// GetCard function adds a card into the account
	GetCard(ctx context.Context, accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// Authorize function executes the veritrans payment
	Authorize(ctx context.Context, param *veritrans.Params) error
	// Authorize function executes the veritrans payment
	Capture(ctx context.Context, param *veritrans.Params) error
	// Cancel function cancels the veritrans payment
	Cancel(ctx context.Context, param *veritrans.Params) error
}
