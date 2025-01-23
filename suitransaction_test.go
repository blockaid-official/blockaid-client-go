// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
)

func TestSuiTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Sui.Transaction.Scan(context.TODO(), blockaidclientgo.SuiTransactionScanParams{
		SuiTransactionScanRequest: blockaidclientgo.SuiTransactionScanRequestParam{
			AccountAddress: blockaidclientgo.F("account_address"),
			Chain:          blockaidclientgo.F(blockaidclientgo.SuiTransactionScanRequestChainMainnet),
			Metadata: blockaidclientgo.F[blockaidclientgo.SuiTransactionScanRequestMetadataUnionParam](blockaidclientgo.SuiTransactionScanRequestMetadataSuiWalletRequestMetadataParam{
				Type: blockaidclientgo.F(blockaidclientgo.SuiTransactionScanRequestMetadataSuiWalletRequestMetadataTypeWallet),
				URL:  blockaidclientgo.F("url"),
			}),
			Transaction: blockaidclientgo.F("transaction"),
			Options:     blockaidclientgo.F([]blockaidclientgo.SuiTransactionScanRequestOption{blockaidclientgo.SuiTransactionScanRequestOptionValidation}),
		},
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
