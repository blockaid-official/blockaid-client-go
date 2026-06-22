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

func TestSiteReport(t *testing.T) {
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
	_, err := client.Site.Report(context.TODO(), blockaidclientgo.SiteReportParams{
		Details: blockaidclientgo.F("Details about the report"),
		Event:   blockaidclientgo.F(blockaidclientgo.SiteReportParamsEventFalsePositive),
		Report: blockaidclientgo.F[blockaidclientgo.SiteReportParamsReportUnion](blockaidclientgo.SiteReportParamsReportRequestIDReport{
			RequestID: blockaidclientgo.F("11111111-1111-1111-1111-111111111111"),
			Type:      blockaidclientgo.F(blockaidclientgo.SiteReportParamsReportRequestIDReportTypeRequestID),
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

func TestSiteScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Site.Scan(context.TODO(), blockaidclientgo.SiteScanParams{
		URL: blockaidclientgo.F("https://app.uniswap.org"),
		Metadata: blockaidclientgo.F(blockaidclientgo.SiteScanParamsMetadata{
			Account: blockaidclientgo.F(blockaidclientgo.SiteScanParamsMetadataAccount{
				AccountID:                blockaidclientgo.F("account_id"),
				AccountAddresses:         blockaidclientgo.F([]string{"string"}),
				AccountCreationTimestamp: blockaidclientgo.F(time.Now()),
				UserAge:                  blockaidclientgo.F(int64(1)),
				UserCountryCode:          blockaidclientgo.F("user_country_code"),
			}),
			Connection: blockaidclientgo.F(blockaidclientgo.SiteScanParamsMetadataConnection{
				IPAddress:                blockaidclientgo.F("ip_address"),
				Origin:                   blockaidclientgo.F("https://example.com"),
				UserAgent:                blockaidclientgo.F("user_agent"),
				WalletconnectDescription: blockaidclientgo.F("walletconnect_description"),
				WalletconnectName:        blockaidclientgo.F("walletconnect_name"),
			}),
			Domain:  blockaidclientgo.F("domain"),
			NonDapp: blockaidclientgo.F(true),
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
