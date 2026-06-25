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

func TestEvmPostTransactionReport(t *testing.T) {
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
	_, err := client.Evm.PostTransaction.Report(context.TODO(), blockaidclientgo.EvmPostTransactionReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.EvmPostTransactionReportParamsEventFalseNegative),
		Report: blockaidclientgo.F[blockaidclientgo.EvmPostTransactionReportParamsReportUnion](blockaidclientgo.EvmPostTransactionReportParamsReportRequestIDReport{
			RequestID: blockaidclientgo.F("11111111-1111-1111-1111-111111111111"),
			Type:      blockaidclientgo.F(blockaidclientgo.EvmPostTransactionReportParamsReportRequestIDReportTypeRequestID),
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

func TestEvmPostTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.PostTransaction.Scan(context.TODO(), blockaidclientgo.EvmPostTransactionScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainEthereum),
		Data: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsData{
			TxHash: blockaidclientgo.F("0xc01780dadc107754b331250b4797606949cb3d0087facc0a737122d5e973c83c"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsMetadata{
			Account: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsMetadataAccount{
				AccountID:                blockaidclientgo.F("account_id"),
				AccountAddresses:         blockaidclientgo.F([]string{"string"}),
				AccountCreationTimestamp: blockaidclientgo.F(time.Now()),
				UserAge:                  blockaidclientgo.F(int64(1)),
				UserCountryCode:          blockaidclientgo.F("user_country_code"),
			}),
			Connection: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsMetadataConnection{
				IPAddress:                blockaidclientgo.F("ip_address"),
				Origin:                   blockaidclientgo.F("https://example.com"),
				UserAgent:                blockaidclientgo.F("user_agent"),
				WalletconnectDescription: blockaidclientgo.F("walletconnect_description"),
				WalletconnectName:        blockaidclientgo.F("walletconnect_name"),
			}),
			Domain:  blockaidclientgo.F("domain"),
			NonDapp: blockaidclientgo.F(true),
		}),
		Block:                    blockaidclientgo.F[blockaidclientgo.EvmPostTransactionScanParamsBlockUnion](shared.UnionInt(int64(0))),
		Options:                  blockaidclientgo.F([]blockaidclientgo.EvmPostTransactionScanParamsOption{blockaidclientgo.EvmPostTransactionScanParamsOptionValidation, blockaidclientgo.EvmPostTransactionScanParamsOptionSimulation}),
		SimulateWithEstimatedGas: blockaidclientgo.F(true),
		StateOverride: blockaidclientgo.F(map[string]blockaidclientgo.EvmPostTransactionScanParamsStateOverride{
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
		TransactionHints: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsTransactionHints{
			CrossChainBridge: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsTransactionHintsCrossChainBridge{
				DestinationAddress: blockaidclientgo.F("destination_address"),
				DestinationAsset: blockaidclientgo.F[blockaidclientgo.EvmPostTransactionScanParamsTransactionHintsCrossChainBridgeDestinationAssetUnion](blockaidclientgo.EvmPostTransactionScanParamsTransactionHintsCrossChainBridgeDestinationAssetCrossChainBridgeNativeAsset{
					Type:     blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsTransactionHintsCrossChainBridgeDestinationAssetCrossChainBridgeNativeAssetTypeNative),
					RawValue: blockaidclientgo.F("raw_value"),
					UsdPrice: blockaidclientgo.F("usd_price"),
				}),
				DestinationChain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
			}),
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
