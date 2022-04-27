package endpoint

import "github.com/david1992121/veritrans-microservice/internal/veritrans"

// GetMDKTokenRequest struct
// veritrans.ClientCardInfo

// GetMDKTokenResponse struct
type GetMDKTokenResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

// CreateAccountRequest struct
// veritrans.AccountParam

// CreateAccountResponse struct
type CreateAccountResponse struct {
	Account *veritrans.Account `json:"account,omitempty"`
	Err     string             `json:"err"`
}

// UpdateAccountRequest struct
// veritrans.AccountParam

// UpdateAccountResponse struct
type UpdateAccountResponse struct {
	Account *veritrans.Account `json:"account,omitempty"`
	Err     string             `json:"err"`
}

// AuthorizeRequest struct
// veritrans.AccountParam

// AuthorizeResponse struct
type AuthorizeResponse struct {
	Err string `json:"err"`
}

// CancelRequest struct
// veritrans.AccountParam

// CancelResponse struct
type CancelResponse struct {
	Err string `json:"err"`
}
