// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaid_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/blockaid-client-go"
	"github.com/stainless-sdks/blockaid-client-go/internal/testutil"
	"github.com/stainless-sdks/blockaid-client-go/option"
)

func TestEvmTransactionBulkScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.TransactionBulk.Scan(context.TODO(), blockaid.EvmTransactionBulkScanParams{
		Chain: blockaid.F("ethereum"),
		Data: blockaid.F([]blockaid.EvmTransactionBulkScanParamsData{{
			From:     blockaid.F("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			To:       blockaid.F("0xA4e5961B58DBE487639929643dCB1Dc3848dAF5E"),
			Data:     blockaid.F("0x"),
			Value:    blockaid.F("0x100000000000"),
			Gas:      blockaid.F("string"),
			GasPrice: blockaid.F("string"),
		}, {
			From:     blockaid.F("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			To:       blockaid.F("0x0D524a5B52737C0a02880d5E84F7D20b8d66bfba"),
			Data:     blockaid.F("0x"),
			Value:    blockaid.F("0xdeadbeef"),
			Gas:      blockaid.F("string"),
			GasPrice: blockaid.F("string"),
		}}),
		Metadata: blockaid.F(blockaid.MetadataParam{
			Domain: blockaid.F("https://example.com"),
		}),
		Options: blockaid.F([]blockaid.EvmTransactionBulkScanParamsOption{blockaid.EvmTransactionBulkScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaid.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
