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
		Options:     blockaidclientgo.F([]blockaidclientgo.BitcoinTransactionScanParamsOption{blockaidclientgo.BitcoinTransactionScanParamsOptionSimulation}),
		Transaction: blockaidclientgo.F("0100000002c5b20d466c445318b0028069042525cd7d18d167d48391ebb17948965edc51460000000000ffffffffcca758769ae4cdecc89d9defb6baf74dd997b8ee656a737412f736b66b5771280000000000ffffffff022a240000000000001976a914f5d7afc3df015ecfd309dd591acf1b8f1e0c4ec088ac68744b00000000001976a914f5d7afc3df015ecfd309dd591acf1b8f1e0c4ec088ac00000000"),
		Chain:       blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsChainBitcoin),
		Metadata: blockaidclientgo.F[blockaidclientgo.BitcoinTransactionScanParamsMetadataUnion](blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataTypeInApp),
		}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
