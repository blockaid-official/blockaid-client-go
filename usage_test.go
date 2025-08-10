// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
)

func TestUsage(t *testing.T) {
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
	transactionScanResponse, err := client.Evm.JsonRpc.Scan(context.TODO(), blockaidclientgo.EvmJsonRpcScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		Data: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsData{
			Method: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsDataMethodEthSignTypedDataV4),
			Params: blockaidclientgo.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}}),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsMetadata{
			Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
		}),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", transactionScanResponse.Validation)
}
