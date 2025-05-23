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

func TestEvmTransactionReport(t *testing.T) {
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
	_, err := client.Evm.Transaction.Report(context.TODO(), blockaidclientgo.EvmTransactionReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.EvmTransactionReportParamsEventFalsePositive),
		Report: blockaidclientgo.F[blockaidclientgo.EvmTransactionReportParamsReportUnion](blockaidclientgo.EvmTransactionReportParamsReportRequestIDReport{
			RequestID: blockaidclientgo.F("11111111-1111-1111-1111-111111111111"),
			Type:      blockaidclientgo.F(blockaidclientgo.EvmTransactionReportParamsReportRequestIDReportTypeRequestID),
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

func TestEvmTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Evm.Transaction.Scan(context.TODO(), blockaidclientgo.EvmTransactionScanParams{
		AccountAddress: blockaidclientgo.F("account_address"),
		Chain:          blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		Data: blockaidclientgo.F(blockaidclientgo.EvmTransactionScanParamsData{
			From: blockaidclientgo.F("0x5e1a0d484c5f0de722e82f9dca3a9d5a421d47cb"),
			AuthorizationList: blockaidclientgo.F([]blockaidclientgo.EvmTransactionScanParamsDataAuthorizationList{{
				Address: blockaidclientgo.F("address"),
				ChainID: blockaidclientgo.F("chainId"),
				Eoa:     blockaidclientgo.F("eoa"),
				Nonce:   blockaidclientgo.F("nonce"),
				R:       blockaidclientgo.F("r"),
				S:       blockaidclientgo.F("s"),
				YParity: blockaidclientgo.F("yParity"),
			}}),
			Data:     blockaidclientgo.F("0x"),
			Gas:      blockaidclientgo.F("gas"),
			GasPrice: blockaidclientgo.F("gas_price"),
			To:       blockaidclientgo.F("0x0d524a5b52737c0a02880d5e84f7d20b8d66bfba"),
			Value:    blockaidclientgo.F("0x1000000000000000"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
		}),
		Block:   blockaidclientgo.F[blockaidclientgo.EvmTransactionScanParamsBlockUnion](shared.UnionString("21211118")),
		Options: blockaidclientgo.F([]blockaidclientgo.EvmTransactionScanParamsOption{blockaidclientgo.EvmTransactionScanParamsOptionSimulation, blockaidclientgo.EvmTransactionScanParamsOptionValidation}),
		StateOverride: blockaidclientgo.F(map[string]blockaidclientgo.EvmTransactionScanParamsStateOverride{
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
