package pkg

import (
	"encoding/json"
	"time"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// NewLoggingMiddleware function
func NewLoggingMiddleware(logger log.Logger, service Service) Service {
	return loggingMiddleware{
		logger: logger,
		next:   service,
	}
}

// GetMDKToken function
func (mw loggingMiddleware) GetMDKToken(cardInfo *veritrans.ClientCardInfo) (output string, err error) {
	cardString, _ := json.Marshal(cardInfo)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetMDKToken",
			"input", cardString,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetMDKToken(cardInfo)
	return
}

// CreateAccount function
func (mw loggingMiddleware) CreateAccount(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateAccount",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.CreateAccount(accountParam)
	return
}

// UpdateAccount function
func (mw loggingMiddleware) UpdateAccount(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "UpdateAccount",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.UpdateAccount(accountParam)
	return
}

// CreateCard function
func (mw loggingMiddleware) CreateCard(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateCard",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.CreateCard(accountParam)
	return
}

// UpdateCard function
func (mw loggingMiddleware) UpdateCard(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "UpdateCard",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.UpdateCard(accountParam)
	return
}

// DeleteCard function
func (mw loggingMiddleware) DeleteCard(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "DeleteCard",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.DeleteCard(accountParam)
	return
}

// GetCard function
func (mw loggingMiddleware) GetCard(accountParam *veritrans.AccountParam) (account *veritrans.Account, err error) {
	inputString, _ := json.Marshal(accountParam)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetCard",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	account, err = mw.next.GetCard(accountParam)
	return
}

// Authorize function
func (mw loggingMiddleware) Authorize(param *veritrans.Params) (err error) {
	inputString, _ := json.Marshal(param)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Authorize",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Authorize(param)
	return
}

// Cancel function
func (mw loggingMiddleware) Cancel(param *veritrans.Params) (err error) {
	inputString, _ := json.Marshal(param)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Cancel",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Cancel(param)
	return
}

// Capture function
func (mw loggingMiddleware) Capture(param *veritrans.Params) (err error) {
	inputString, _ := json.Marshal(param)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Capture",
			"input", inputString,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Capture(param)
	return
}
