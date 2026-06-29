// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaidclientgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/internal/testutil"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/blockaid-official/blockaid-client-go/shared"
)

func TestEvmTransactionRawScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.TransactionRaw.Scan(context.TODO(), blockaidclientgo.EvmTransactionRawScanParams{
		AccountAddress: blockaidclientgo.F("0x362320f3a3eeeb4c4699b1b9062a84B2612bcdba"),
		Chain:          blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainEthereum),
		Data:           blockaidclientgo.F("0x02f903f8018208488405f5e100850a9a03feb38302fa6a941111111254eeb25477b68fb85ed929f73a96058280b9038862e238bb00000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000002e00000000000000000000000000000000000000000000000000000000000000360000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ab3e25398a24d6af080000000000000000000000000000000000000000000000000000000070db68f000000000000000000000000000000000000000000000000000000c0c2f020e4000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000006e2a43be0b1d33b726f0ca3b8de60b3482b8b050000000000000000000000000b78ed0dd769e3fbd8e2b526f6f75dcccc7e2af4f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000773594000000000000000000000000000000000000000000000000b4b34aede8e617e060000000a4000000a4000000a4000000a400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000a4bf15fcd8000000000000000000000000303389f541ff2d620e42832f180a08e767b28e10000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000242cc2878d0064b6509600000000060000b78ed0dd769e3fbd8e2b526f6f75dcccc7e2af4f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041a946fd8d5b9e873563a1411cbdf290b8310d8cdddc94da3aebf95b16a6dc0bf56d736ece63e3906527b7dcf08aa845d6a5cd4e0d99c9994f617b6faa378317f71c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e26b9977c001a0386a78e05b5ab3b4badbae843aaa6ed379b2f4353aa730c8f360141e72cba692a036ed37a00c6364685e58bc0cd9cdd1c140d6690c3c0d216c1b67e3d262e2f0f9"),
		Metadata: blockaidclientgo.F(blockaidclientgo.EvmTransactionRawScanParamsMetadata{
			Account: blockaidclientgo.F(blockaidclientgo.EvmTransactionRawScanParamsMetadataAccount{
				AccountID:                blockaidclientgo.F("account_id"),
				AccountAddresses:         blockaidclientgo.F([]string{"string"}),
				AccountCreationTimestamp: blockaidclientgo.F(time.Now()),
				UserAge:                  blockaidclientgo.F(int64(1)),
				UserCountryCode:          blockaidclientgo.F("user_country_code"),
			}),
			Connection: blockaidclientgo.F(blockaidclientgo.EvmTransactionRawScanParamsMetadataConnection{
				IPAddress:                blockaidclientgo.F("ip_address"),
				Origin:                   blockaidclientgo.F("https://example.com"),
				UserAgent:                blockaidclientgo.F("user_agent"),
				WalletconnectDescription: blockaidclientgo.F("walletconnect_description"),
				WalletconnectName:        blockaidclientgo.F("walletconnect_name"),
			}),
			Domain:  blockaidclientgo.F("https://app.1inch.io"),
			NonDapp: blockaidclientgo.F(true),
		}),
		Block:                    blockaidclientgo.F[blockaidclientgo.EvmTransactionRawScanParamsBlockUnion](shared.UnionString("17718858")),
		Options:                  blockaidclientgo.F([]blockaidclientgo.EvmTransactionRawScanParamsOption{blockaidclientgo.EvmTransactionRawScanParamsOptionSimulation, blockaidclientgo.EvmTransactionRawScanParamsOptionValidation}),
		SimulateWithEstimatedGas: blockaidclientgo.F(true),
		StateOverride: blockaidclientgo.F(map[string]blockaidclientgo.EvmTransactionRawScanParamsStateOverride{
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
		TransactionHints: blockaidclientgo.F([]blockaidclientgo.EvmTransactionRawScanParamsTransactionHintUnion{blockaidclientgo.EvmTransactionRawScanParamsTransactionHintsCrossChainBridgeHint{
			Type:               blockaidclientgo.F(blockaidclientgo.EvmTransactionRawScanParamsTransactionHintsCrossChainBridgeHintTypeCrossChainBridge),
			DestinationAddress: blockaidclientgo.F("destination_address"),
			DestinationAsset: blockaidclientgo.F[blockaidclientgo.EvmTransactionRawScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetUnion](blockaidclientgo.EvmTransactionRawScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAsset{
				Type:     blockaidclientgo.F(blockaidclientgo.EvmTransactionRawScanParamsTransactionHintsCrossChainBridgeHintDestinationAssetCrossChainBridgeNativeAssetTypeNative),
				RawValue: blockaidclientgo.F("raw_value"),
				UsdPrice: blockaidclientgo.F("usd_price"),
			}),
			DestinationChain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		}}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
