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

func TestEvmAddressReport(t *testing.T) {
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
	_, err := client.Evm.Address.Report(context.TODO(), blockaidclientgo.EvmAddressReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.EvmAddressReportParamsEventFalseNegative),
		Report: blockaidclientgo.F[blockaidclientgo.EvmAddressReportParamsReportUnion](blockaidclientgo.EvmAddressReportParamsReportRequestIDReport{
			RequestID: blockaidclientgo.F("11111111-1111-1111-1111-111111111111"),
			Type:      blockaidclientgo.F(blockaidclientgo.EvmAddressReportParamsReportRequestIDReportTypeRequestID),
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

func TestEvmAddressScan(t *testing.T) {
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
	_, err := client.Evm.Address.Scan(context.TODO(), blockaidclientgo.EvmAddressScanParams{
		ValidateAddress: blockaidclientgo.ValidateAddressParam{
			Address: blockaidclientgo.F("0x946D45c866AFD5b8F436d40E551D8E50A5B84230"),
			Chain:   blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChainEthereum),
			Metadata: blockaidclientgo.F(blockaidclientgo.ValidateAddressMetadataParam{
				Domain: blockaidclientgo.F("https://example.com"),
			}),
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
