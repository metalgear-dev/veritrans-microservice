package test

import (
	"context"
	"log"
	"net"
	"regexp"
	"testing"

	"github.com/david1992121/veritrans-microservice/api/pb"
	"github.com/david1992121/veritrans-microservice/pkg/transport"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	logger := initLogger()

	listener = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	pb.RegisterVeritransServer(server, transport.GetGRPCServer(logger))
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Server exited with error; %s", err.Error())
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func getClient() (context.Context, pb.VeritransClient, error) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
		return nil, nil, err
	}
	return ctx, pb.NewVeritransClient(conn), nil
}

// TestGRPCMKS function
func TestGRPCMDK(t *testing.T) {
	ctx, client, err := getClient()
	assert.Nil(t, err)

	resp, err := client.GetMDKToken(ctx, &pb.GetMDKTokenRequest{
		CardNumber:   "4111111111111111",
		CardExpire:   "12/23",
		SecurityCode: "123",
	})
	assert.Nil(t, err)

	re := regexp.MustCompile(`[0-9a-z\-]{36}`)
	assert.Equal(t, true, re.Match([]byte(resp.Token)))
	assert.Equal(t, "", resp.Err)
}

// TestGRPCAccount function
func TestGRPCAccount(t *testing.T) {
	ctx, client, err := getClient()
	assert.Nil(t, err)

	testAccountID := "test-grpc-account-01"
	resp, err := client.CreateAccount(ctx, &pb.AccountRequest{
		AccountID: testAccountID,
	})
	assert.Nil(t, err)

	if resp.Account != nil {
		assert.Equal(t, testAccountID, resp.Account.AccountID)
	} else {
		assert.Equal(t, "入会中の会員です。", resp.Err)
	}
}

// TestGRPCCard function
func TestGRPCCard(t *testing.T) {
	ctx, client, err := getClient()
	assert.Nil(t, err)

	testAccountID := "test-grpc-account-01"
	var cardID string

	// add card
	{
		cardNumber := "4111111111111111"
		cardNumberExpected := "411111********11"
		cardExpire := "12/23"
		defaultCard := "1"

		resp, err := client.CreateCard(ctx, &pb.AccountRequest{
			AccountID: testAccountID,
			CardParam: &pb.AccountRequest_CardParam{
				CardExpire:  &cardExpire,
				CardNumber:  &cardNumber,
				DefaultCard: &defaultCard,
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "", resp.Err)
		assert.Equal(t, resp.Account.AccountID, testAccountID)
		assert.Equal(t, 1, len(resp.Account.CardInfo))
		assert.Equal(t, resp.Account.CardInfo[0].CardNumber, cardNumberExpected)
		assert.Equal(t, resp.Account.CardInfo[0].CardExpire, cardExpire)
		cardID = resp.Account.CardInfo[0].CardID
	}

	// get card
	{
		resp, err := client.GetCard(ctx, &pb.AccountRequest{
			AccountID: testAccountID,
		})

		assert.Nil(t, err)
		assert.Equal(t, "", resp.Err)
		assert.Equal(t, resp.Account.AccountID, testAccountID)
		assert.Equal(t, 1, len(resp.Account.CardInfo))
		assert.Equal(t, cardID, resp.Account.CardInfo[0].CardID)
	}

	// update card
	{
		newCardExpire := "12/24"
		resp, err := client.UpdateCard(ctx, &pb.AccountRequest{
			AccountID: testAccountID,
			CardParam: &pb.AccountRequest_CardParam{
				CardID:     &cardID,
				CardExpire: &newCardExpire,
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, "", resp.Err)
	}

	// remove card
	{
		resp, err := client.DeleteCard(ctx, &pb.AccountRequest{
			AccountID: testAccountID,
			CardParam: &pb.AccountRequest_CardParam{
				CardID: &cardID,
			},
		})
		assert.Nil(t, err)
		assert.Equal(t, "", resp.Err)
	}
}
