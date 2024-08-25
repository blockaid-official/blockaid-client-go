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
		option.WithClientID("My Client ID"),
	)
	_, err := client.Evm.PostTransaction.Report(context.TODO(), blockaidclientgo.EvmPostTransactionReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.EvmPostTransactionReportParamsEventFalsePositive),
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
		option.WithClientID("My Client ID"),
	)
	_, err := client.Evm.PostTransaction.Scan(context.TODO(), blockaidclientgo.EvmPostTransactionScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainArbitrum),
		Data: blockaidclientgo.F(blockaidclientgo.EvmPostTransactionScanParamsData{
			TxHash: blockaidclientgo.F("0xc01780dadc107754b331250b4797606949cb3d0087facc0a737122d5e973c83c"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("domain"),
		}),
		Block:   blockaidclientgo.F[blockaidclientgo.EvmPostTransactionScanParamsBlockUnion](shared.UnionInt(int64(0))),
		Options: blockaidclientgo.F([]blockaidclientgo.EvmPostTransactionScanParamsOption{blockaidclientgo.EvmPostTransactionScanParamsOptionValidation, blockaidclientgo.EvmPostTransactionScanParamsOptionSimulation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
