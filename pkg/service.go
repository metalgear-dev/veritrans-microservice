package pkg

import (
	"github.com/david1992121/veritrans-microservice/internal/veritrans"
)

// Service of the veritrans payment
type Service interface {
	// GetMDKToken function gets the MDK token from card information
	GetMDKToken(cardInfo *veritrans.ClientCardInfo) (string, error)
	// CreateAccount function creates a veritrans account
	CreateAccount(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// UpdateAccount function updates the veritrans account
	UpdateAccount(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// CreateCard function adds a card into the account
	CreateCard(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// UpdateCard function adds a card into the account
	UpdateCard(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// DeleteCard function adds a card into the account
	DeleteCard(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// GetCard function adds a card into the account
	GetCard(accountParam *veritrans.AccountParam) (*veritrans.Account, error)
	// Authorize function executes the veritrans payment
	Authorize(param *veritrans.Params) error
	// Authorize function executes the veritrans payment
	Capture(param *veritrans.Params) error
	// Cancel function cancels the veritrans payment
	Cancel(param *veritrans.Params) error
}
