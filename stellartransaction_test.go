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

func TestStellarTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Stellar.Transaction.Scan(context.TODO(), blockaidclientgo.StellarTransactionScanParams{
		StellarTransactionScanRequest: blockaidclientgo.StellarTransactionScanRequestParam{
			Chain:          blockaidclientgo.F(blockaidclientgo.StellarTransactionScanRequestChainPubnet),
			Options:        blockaidclientgo.F([]blockaidclientgo.StellarTransactionScanRequestOption{blockaidclientgo.StellarTransactionScanRequestOptionValidation}),
			AccountAddress: blockaidclientgo.F("string"),
			Transactions:   blockaidclientgo.F([]string{"string", "string", "string"}),
			Metadata: blockaidclientgo.F[blockaidclientgo.StellarTransactionScanRequestMetadataUnionParam](blockaidclientgo.StellarTransactionScanRequestMetadataStellarWalletRequestMetadataParam{
				Type: blockaidclientgo.F(blockaidclientgo.StellarTransactionScanRequestMetadataStellarWalletRequestMetadataTypeWallet),
				URL:  blockaidclientgo.F("string"),
			}),
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
