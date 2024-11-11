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

func TestBitcoinTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Bitcoin.Transaction.Scan(context.TODO(), blockaidclientgo.BitcoinTransactionScanParams{
		AccountAddress: blockaidclientgo.F("account_address"),
		Chain:          blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsChainBitcoin),
		Metadata: blockaidclientgo.F[blockaidclientgo.BitcoinTransactionScanParamsMetadataUnion](blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinblockaidApplicationRunnerAppSchemasTransactionScanningAPIWalletRequestMetadataTypeWallet),
			URL:  blockaidclientgo.F("url"),
		}),
		Transaction: blockaidclientgo.F("transaction"),
		Options:     blockaidclientgo.F([]blockaidclientgo.BitcoinTransactionScanParamsOption{blockaidclientgo.BitcoinTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
