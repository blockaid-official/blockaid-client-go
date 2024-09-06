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
		Transaction: blockaidclientgo.F("0100000001cca758769ae4cdecc89d9defb6baf74dd997b8ee656a737412f736b66b5771280000000000ffffffff0b544b0600000000001976a91439309fb127bba85d50c3379be9ac2d05edff0d7388ac544b0600000000001976a9146606e4efd46f43aa0bafe74f61402366329edda688ac544b0600000000001976a914163f1a3342956eda716d056be582f317ffbfc2ef88ac544b0600000000001976a91404246d84e0125114def10c5a6c7e02d1cadf189688ac544b0600000000001976a914cfb9b949ef347daa8afaa92538be145dd7f4907288ac544b0600000000001976a914070e1228753fbf00f3ec25ae94e186e3414541a288ac544b0600000000001976a914a3b1289cda60e3fad45f03636096516f838a61b388ac544b0600000000001976a9142457ae343ff877bd029491f7858baceb1c0a78a188ac544b0600000000001976a9145a27d6993f3583ecd6d0d134f9d9a69cb143c7d888ac544b0600000000001976a914b50cf380d7e52814854f4c35a7088eb963fb9f3688ac20a10700000000001976a914f5d7afc3df015ecfd309dd591acf1b8f1e0c4ec088ac00000000"),
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
