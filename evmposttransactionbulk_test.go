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

func TestEvmPostTransactionBulkScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.PostTransactionBulk.Scan(context.TODO(), blockaidclientgo.EvmPostTransactionBulkScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChain(blockaidclientgo.TransactionScanSupportedChainEthereum)),
		Data: blockaidclientgo.F([]blockaidclientgo.EvmPostTransactionBulkScanParamsData{{
			From:     blockaidclientgo.F("string"),
			To:       blockaidclientgo.F("string"),
			Data:     blockaidclientgo.F("string"),
			Value:    blockaidclientgo.F("string"),
			Gas:      blockaidclientgo.F("string"),
			GasPrice: blockaidclientgo.F("string"),
		}, {
			From:     blockaidclientgo.F("string"),
			To:       blockaidclientgo.F("string"),
			Data:     blockaidclientgo.F("string"),
			Value:    blockaidclientgo.F("string"),
			Gas:      blockaidclientgo.F("string"),
			GasPrice: blockaidclientgo.F("string"),
		}}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("string"),
		}),
		Options: blockaidclientgo.F([]blockaidclientgo.EvmPostTransactionBulkScanParamsOption{blockaidclientgo.EvmPostTransactionBulkScanParamsOptionValidation, blockaidclientgo.EvmPostTransactionBulkScanParamsOptionSimulation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
