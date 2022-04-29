package endpoint

import "github.com/david1992121/veritrans-microservice/internal/veritrans"

// GetMDKTokenRequest struct
// veritrans.ClientCardInfo

// GetMDKTokenResponse struct
type GetMDKTokenResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// AccountRequest struct
// veritrans.AccountParam

// AccountResponse struct
type AccountResponse struct {
	Account *veritrans.Account `json:"account,omitempty"`
	Err     string             `json:"err"`
}

// PaymentRequest struct
// veritrans.AccountParam

// PaymentResponse struct
type PaymentResponse struct {
	Err string `json:"err"`
}
