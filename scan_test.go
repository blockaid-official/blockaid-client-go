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
)

func TestScanReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Scan.Report(context.TODO(), blockaidclientgo.ScanReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.ScanReportParamsEventFalsePositive),
		Metadata: blockaidclientgo.F(blockaidclientgo.ScanReportParamsMetadata{
			Account: blockaidclientgo.F(blockaidclientgo.ScanReportParamsMetadataAccount{
				AccountID:                blockaidclientgo.F("user-abc-123"),
				AccountAddresses:         blockaidclientgo.F([]string{"string"}),
				AccountCreationTimestamp: blockaidclientgo.F(time.Now()),
				UserAge:                  blockaidclientgo.F(int64(30)),
				UserCountryCode:          blockaidclientgo.F("US"),
			}),
			Connection: blockaidclientgo.F(blockaidclientgo.ScanReportParamsMetadataConnection{
				IPAddress:                blockaidclientgo.F("192.168.1.1"),
				Origin:                   blockaidclientgo.F("https://example.com"),
				UserAgent:                blockaidclientgo.F("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)"),
				WalletconnectDescription: blockaidclientgo.F("walletconnect_description"),
				WalletconnectName:        blockaidclientgo.F("walletconnect_name"),
			}),
			Domain: blockaidclientgo.F("https://app.uniswap.org"),
		}),
		RequestID: blockaidclientgo.F("11111111-1111-1111-1111-111111111111"),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScanStatus(t *testing.T) {
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
	_, err := client.Scan.Status(context.TODO(), blockaidclientgo.ScanStatusParams{
		RequestID: blockaidclientgo.F("7f959417-76c1-4c4d-89e8-5fdedab76a8d"),
		Status:    blockaidclientgo.F(blockaidclientgo.ScanStatusParamsStatusAccepted),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
