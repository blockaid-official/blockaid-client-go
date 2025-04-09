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
			Method: blockaidclientgo.F("eth_signTypedData_v4"),
			Params: blockaidclientgo.F([]interface{}{"0x49c73c9d361c04769a452E85D343b41aC38e0EE4", `{"domain":{"chainId":1,"name":"Aave interest bearing WETH","version":"1","verifyingContract":"0x030ba81f1c18d280636f32af80b9aad02cf0854e"},"message":{"owner":"0x49c73c9d361c04769a452E85D343b41aC38e0EE4","spender":"0xa74cbd5b80f73b5950768c8dc467f1c6307c00fd","value":"115792089237316195423570985008687907853269984665640564039457584007913129639935","nonce":"0","deadline":"1988064000","holder":"0x49c73c9d361c04769a452E85D343b41aC38e0EE4"},"primaryType":"Permit","types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Permit":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"nonce","type":"uint256"},{"name":"deadline","type":"uint256"}]}}`}),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
		}),
		AccountAddress: blockaidclientgo.F("0x49c73c9d361c04769a452E85D343b41aC38e0EE4"),
		Block:          blockaidclientgo.F[blockaidclientgo.EvmJsonRpcScanParamsBlockUnion](shared.UnionString("18370320")),
		Options:        blockaidclientgo.F([]blockaidclientgo.EvmJsonRpcScanParamsOption{blockaidclientgo.EvmJsonRpcScanParamsOptionSimulation, blockaidclientgo.EvmJsonRpcScanParamsOptionValidation}),
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
