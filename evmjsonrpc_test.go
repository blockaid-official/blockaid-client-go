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

func TestEvmJsonRpcScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.JsonRpc.Scan(context.TODO(), blockaidclientgo.EvmJsonRpcScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		Data: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsData{
			Method: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsDataMethodEthSignTypedDataV4),
			Params: blockaidclientgo.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}}),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsMetadata{
			Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
		}),
		AccountAddress:           blockaidclientgo.F("0x49c73c9d361c04769a452E85D343b41aC38e0EE4"),
		Block:                    blockaidclientgo.F[blockaidclientgo.EvmJsonRpcScanParamsBlockUnion](shared.UnionString("18370320")),
		Options:                  blockaidclientgo.F([]blockaidclientgo.EvmJsonRpcScanParamsOption{blockaidclientgo.EvmJsonRpcScanParamsOptionSimulation, blockaidclientgo.EvmJsonRpcScanParamsOptionValidation}),
		SimulateWithEstimatedGas: blockaidclientgo.F(true),
		StateOverride: blockaidclientgo.F(map[string]blockaidclientgo.EvmJsonRpcScanParamsStateOverride{
			"foo": {
				Balance:                 blockaidclientgo.F("balance"),
				Code:                    blockaidclientgo.F("code"),
				MovePrecompileToAddress: blockaidclientgo.F("movePrecompileToAddress"),
				Nonce:                   blockaidclientgo.F("nonce"),
				State: blockaidclientgo.F(map[string]string{
					"foo": "string",
				}),
				StateDiff: blockaidclientgo.F(map[string]string{
					"foo": "string",
				}),
			},
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
