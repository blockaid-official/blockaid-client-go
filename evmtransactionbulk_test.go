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
	"github.com/blockaid-official/blockaid-client-go/shared"
)

func TestEvmTransactionBulkScanWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Evm.TransactionBulk.Scan(context.TODO(), blockaidclientgo.EvmTransactionBulkScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		Data: blockaidclientgo.F([]blockaidclientgo.EvmTransactionBulkScanParamsData{{
			From:     blockaidclientgo.F("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			Data:     blockaidclientgo.F("0x"),
			Gas:      blockaidclientgo.F("gas"),
			GasPrice: blockaidclientgo.F("gas_price"),
			To:       blockaidclientgo.F("0xA4e5961B58DBE487639929643dCB1Dc3848dAF5E"),
			Value:    blockaidclientgo.F("0x100000000000"),
		}, {
			From:     blockaidclientgo.F("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			Data:     blockaidclientgo.F("0x"),
			Gas:      blockaidclientgo.F("gas"),
			GasPrice: blockaidclientgo.F("gas_price"),
			To:       blockaidclientgo.F("0x0D524a5B52737C0a02880d5E84F7D20b8d66bfba"),
			Value:    blockaidclientgo.F("0xdeadbeef"),
		}}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("https://example.com"),
		}),
		Block:   blockaidclientgo.F[blockaidclientgo.EvmTransactionBulkScanParamsBlockUnion](shared.UnionString("20224477")),
		Options: blockaidclientgo.F([]blockaidclientgo.EvmTransactionBulkScanParamsOption{blockaidclientgo.EvmTransactionBulkScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
