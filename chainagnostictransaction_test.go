// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
)

func TestChainAgnosticTransactionScanWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blockaidclientgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ChainAgnostic.Transaction.Scan(context.TODO(), blockaidclientgo.ChainAgnosticTransactionScanParams{
		Data: blockaidclientgo.F(blockaidclientgo.ChainAgnosticTransactionScanParamsData{
			Amount: blockaidclientgo.F(1.000000),
			Asset: blockaidclientgo.F[blockaidclientgo.ChainAgnosticTransactionScanParamsDataAssetUnion](blockaidclientgo.ChainAgnosticTransactionScanParamsDataAssetAssetSymbol{
				Symbol: blockaidclientgo.F("BTC"),
			}),
			Chain: blockaidclientgo.F(blockaidclientgo.ChainAgnosticTransactionScanParamsDataChainBitcoin),
			To:    blockaidclientgo.F("bc1qanrfutwqh854g74lqrygr55jkgf99em4lpfm80"),
			From:  blockaidclientgo.F("from"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.ChainAgnosticTransactionScanParamsMetadata{
			Account: blockaidclientgo.F(blockaidclientgo.ChainAgnosticTransactionScanParamsMetadataAccount{
				AccountID:                blockaidclientgo.F("1"),
				AccountCreationTimestamp: blockaidclientgo.F(time.Now()),
				UserAge:                  blockaidclientgo.F(int64(1)),
				UserCountryCode:          blockaidclientgo.F("user_country_code"),
			}),
			Connection: blockaidclientgo.F(blockaidclientgo.ChainAgnosticTransactionScanParamsMetadataConnection{
				IPAddress: blockaidclientgo.F("1.1.1.1"),
				UserAgent: blockaidclientgo.F("1"),
			}),
		}),
		Options: blockaidclientgo.F([]blockaidclientgo.ChainAgnosticTransactionScanParamsOption{blockaidclientgo.ChainAgnosticTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
