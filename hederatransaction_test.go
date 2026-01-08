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

func TestHederaTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Hedera.Transaction.Scan(context.TODO(), blockaidclientgo.HederaTransactionScanParams{
		AccountAddress: blockaidclientgo.F[any]("0.0.9352077"),
		Chain:          blockaidclientgo.F(blockaidclientgo.HederaTransactionScanParamsChainMainnet),
		Metadata: blockaidclientgo.F[blockaidclientgo.HederaTransactionScanParamsMetadataUnion](blockaidclientgo.HederaTransactionScanParamsMetadataHederaWalletRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.HederaTransactionScanParamsMetadataHederaWalletRequestMetadataTypeWallet),
			URL:  blockaidclientgo.F("https://example.com"),
		}),
		Transaction: blockaidclientgo.F("KmEKXQoVCgwIjvztygYQn6yo3QISBRiawrcDEgIYBhiAwtcvIgIIeDIVSEJBUiB0cmFuc2ZlciBleGFtcGxlciAKHgoNCgUYjee6BBD/p9a5BwoNCgUYha/rARCAqNa5BxIA"),
		Options:     blockaidclientgo.F([]blockaidclientgo.HederaTransactionScanParamsOption{blockaidclientgo.HederaTransactionScanParamsOptionSimulation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
