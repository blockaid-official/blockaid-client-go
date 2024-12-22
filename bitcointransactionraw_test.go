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

func TestBitcoinTransactionRawReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Bitcoin.TransactionRaw.Report(context.TODO(), blockaidclientgo.BitcoinTransactionRawReportParams{
		Details: blockaidclientgo.F("details"),
		Event:   blockaidclientgo.F(blockaidclientgo.BitcoinTransactionRawReportParamsEventShouldBeMalicious),
		Report: blockaidclientgo.F[blockaidclientgo.BitcoinTransactionRawReportParamsReportUnion](blockaidclientgo.BitcoinTransactionRawReportParamsReportBitcoinAppealRequestID{
			ID:   blockaidclientgo.F("id"),
			Type: blockaidclientgo.F(blockaidclientgo.BitcoinTransactionRawReportParamsReportBitcoinAppealRequestIDTypeRequestID),
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

func TestBitcoinTransactionRawScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Bitcoin.TransactionRaw.Scan(context.TODO(), blockaidclientgo.BitcoinTransactionRawScanParams{
		BitcoinTransactionScanRequest: blockaidclientgo.BitcoinTransactionScanRequestParam{
			AccountAddress: blockaidclientgo.F("account_address"),
			Chain:          blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanRequestChainBitcoin),
			Metadata: blockaidclientgo.F[blockaidclientgo.BitcoinTransactionScanRequestMetadataUnionParam](blockaidclientgo.BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataParam{
				Type: blockaidclientgo.F(blockaidclientgo.BitcoinTransactionScanRequestMetadataBitcoinWalletRequestMetadataTypeWallet),
				URL:  blockaidclientgo.F("url"),
			}),
			Transaction: blockaidclientgo.F("transaction"),
			Options:     blockaidclientgo.F([]blockaidclientgo.BitcoinTransactionScanRequestOption{blockaidclientgo.BitcoinTransactionScanRequestOptionValidation}),
		},
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
