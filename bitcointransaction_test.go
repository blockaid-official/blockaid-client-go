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
		option.WithClientID("My Client ID"),
	)
	_, err := client.Bitcoin.Transaction.Scan(context.TODO(), blockaidclientgo.BitcoinTransactionScanParams{
		Options:     blockaidclientgo.F([]blockaidclientgo.BitcoinTransactionScanParamsOption{blockaidclientgo.BitcoinTransactionScanParamsOptionSimulation}),
		Transaction: blockaidclientgo.F("0100000001194ebd43f14daef7ea1479a5b694e0cbfe8f036bf8b3debaffb9e7e217b3abf70100000000ffffffff0200e1f505000000001976a9143ec6c3ed8dfc3ceabcc1cbdb0c5aef4e2d02873c88acdf84448f0c00000017a914e8a4c2bc45640654bb0b915f98c5e72508cff3768700000000"),
		Chain:       blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsChainBitcoin),
		Metadata: blockaidclientgo.F[blockaidclientgo.BitcoinTransactionScanParamsMetadataUnion](blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanParamsMetadataBitcoinInAppRequestMetadataTypeInApp),
		}),
		SkipUtxoCheck: blockaidclientgo.F(true),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
