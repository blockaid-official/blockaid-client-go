// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaid_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/blockaid-go"
	"github.com/stainless-sdks/blockaid-go/internal/testutil"
	"github.com/stainless-sdks/blockaid-go/option"
)

func TestEvmTransactionScanWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blockaid.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Evm.Transaction.Scan(context.TODO(), blockaid.EvmTransactionScanParams{
		AccountAddress: blockaid.F("string"),
		Chain:          blockaid.F("ethereum"),
		Data: blockaid.F(blockaid.EvmTransactionScanParamsData{
			From:     blockaid.F("0x5e1a0d484c5f0de722e82f9dca3a9d5a421d47cb"),
			To:       blockaid.F("0x0d524a5b52737c0a02880d5e84f7d20b8d66bfba"),
			Data:     blockaid.F("0x"),
			Value:    blockaid.F("0x1000000000000000"),
			Gas:      blockaid.F("string"),
			GasPrice: blockaid.F("string"),
		}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://boredapeyartclub.com"),
		}),
		Options: blockaid.F([]blockaid.EvmTransactionScanParamsOption{blockaid.EvmTransactionScanParamsOptionSimulation, blockaid.EvmTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaid.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
