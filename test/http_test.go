package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/david1992121/veritrans-microservice/internal/veritrans"
	"github.com/david1992121/veritrans-microservice/pkg/endpoint"
	"github.com/david1992121/veritrans-microservice/pkg/transport"
	"github.com/go-kit/kit/log"
	assert "github.com/stretchr/testify/require"
)

var httpHandler http.Handler

func init() {
	logger := initLogger()
	httpHandler = transport.NewHTTPHandler(logger)
}

// TestMDK tests the request of mdk card token
func TestMDK(t *testing.T) {
	jsonStr := []byte(`{"card_number":"4111111111111111","card_expire":"12/22","security_code":"123"}`)
	req := httptest.NewRequest(http.MethodPost, "/mdk/token", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()

	httpHandler.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, 200)

	var resStruct endpoint.GetMDKTokenResponse
	err := json.Unmarshal([]byte(rec.Body.String()), &resStruct)
	assert.Nil(t, err)

	re := regexp.MustCompile(`[0-9a-z\-]{36}`)
	assert.Equal(t, true, re.Match([]byte(resStruct.Token)))
	assert.Equal(t, "", resStruct.Err)
}

// TestAccount function
func TestAccount(t *testing.T) {
	testAccountID := "test-account-001"
	jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s"}`, testAccountID))
	req := httptest.NewRequest(http.MethodPost, "/account/create", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()

	httpHandler.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, 200)

	var accountRes endpoint.AccountResponse
	err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
	assert.Nil(t, err)

	if accountRes.Account != nil {
		assert.Equal(t, testAccountID, accountRes.Account.AccountID)
	} else {
		assert.Equal(t, "入会中の会員です。", accountRes.Err)
	}
}

// TestCard function
func TestCard(t *testing.T) {
	testAccountID := "test-account-001"
	var cardID string
	var accountRes endpoint.AccountResponse

	// add card
	{
		cardNumber := "4111111111111111"
		cardNumberExpected := "411111********11"
		cardExpire := "12/23"
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s","cardParam":{"cardNumber":"%s","cardExpire":"%s","defaultCard":"1"}}`,
			testAccountID, cardNumber, cardExpire))
		req := httptest.NewRequest(http.MethodPost, "/card/create", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
		assert.Equal(t, accountRes.Account.AccountID, testAccountID)
		assert.Equal(t, 1, len(accountRes.Account.CardInfo))
		assert.Equal(t, accountRes.Account.CardInfo[0].CardNumber, cardNumberExpected)
		assert.Equal(t, accountRes.Account.CardInfo[0].CardExpire, cardExpire)
		cardID = accountRes.Account.CardInfo[0].CardID
	}

	// get card
	{
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s"}`, testAccountID))
		req := httptest.NewRequest(http.MethodPost, "/card/get", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)
		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
		assert.Equal(t, accountRes.Account.AccountID, testAccountID)
		assert.Equal(t, 1, len(accountRes.Account.CardInfo))
		assert.Equal(t, cardID, accountRes.Account.CardInfo[0].CardID)
	}

	// update card
	{
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s","cardParam":{"cardId":"%s","cardExpire":"12/24"}}`,
			testAccountID, cardID))
		req := httptest.NewRequest(http.MethodPost, "/card/update", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
	}

	// remove card
	{
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s","cardParam":{"cardId": "%s"}}`, testAccountID, cardID))
		req := httptest.NewRequest(http.MethodPost, "/card/delete", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
	}
}

// TestPayment function
func TestPayment(t *testing.T) {
	testAccountID := "test-account-001"
	var accountRes endpoint.AccountResponse
	var cardID string

	// add card
	{
		cardNumber := "4111111111111111"
		cardExpire := "12/23"
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s","cardParam":{"cardNumber":"%s","cardExpire":"%s","defaultCard":"1"}}`,
			testAccountID, cardNumber, cardExpire))
		req := httptest.NewRequest(http.MethodPost, "/card/create", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()
		httpHandler.ServeHTTP(rec, req)

		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
		assert.Equal(t, 1, len(accountRes.Account.CardInfo))
		cardID = accountRes.Account.CardInfo[0].CardID
	}

	// authorize with capture
	orderNumber := 0
	var paymentRes endpoint.PaymentResponse
	for {
		orderNumber = veritrans.GetRandomID(8)
		testOrderID := fmt.Sprintf("test-account-order-%d", orderNumber)
		jsonStr := []byte(fmt.Sprintf(`{"orderId":"%s", "amount":"100", "jpo": "10", "withCapture": "true", `+
			`"payNowIDParam":{"accountParam":{"accountId":"%s"}}}`, testOrderID, testAccountID))
		req := httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &paymentRes)
		assert.Nil(t, err)
		if paymentRes.Err != "" {
			continue
		}
		break
	}

	// authorize without capture
	for {
		orderNumber = veritrans.GetRandomID(8)
		testOrderID := fmt.Sprintf("test-account-order-%d", orderNumber)
		jsonStr := []byte(fmt.Sprintf(`{"orderId":"%s", "amount":"100", "jpo": "10", "withCapture": "false", `+
			`"payNowIDParam":{"accountParam":{"accountId":"%s"}}}`, testOrderID, testAccountID))
		req := httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &paymentRes)
		assert.Nil(t, err)
		if paymentRes.Err != "" {
			continue
		}
		break
	}

	// capture
	{
		testOrderID := fmt.Sprintf("test-account-order-%d", orderNumber)
		jsonStr := []byte(fmt.Sprintf(`{"orderId":"%s", "amount": "100"}`, testOrderID))
		req := httptest.NewRequest(http.MethodPost, "/capture", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &paymentRes)
		assert.Nil(t, err)
		assert.Equal(t, "", paymentRes.Err)
	}

	// remove card
	{
		jsonStr := []byte(fmt.Sprintf(`{"accountId":"%s","cardParam":{"cardId": "%s"}}`, testAccountID, cardID))
		req := httptest.NewRequest(http.MethodPost, "/card/delete", bytes.NewBuffer(jsonStr))
		rec := httptest.NewRecorder()

		httpHandler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, 200)

		err := json.Unmarshal([]byte(rec.Body.String()), &accountRes)
		assert.Nil(t, err)
		assert.Equal(t, "", accountRes.Err)
	}
}

func initLogger() log.Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	return logger
}
