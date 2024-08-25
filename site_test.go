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
		option.WithClientID("My Client ID"),
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
		option.WithClientID("My Client ID"),
	)
	_, err := client.Site.Scan(context.TODO(), blockaidclientgo.SiteScanParams{
		URL: blockaidclientgo.F("https://app.uniswap.org"),
		Metadata: blockaidclientgo.F[blockaidclientgo.SiteScanParamsMetadataUnion](blockaidclientgo.SiteScanParamsMetadataCatalogRequestMetadata{
			Type: blockaidclientgo.F(blockaidclientgo.SiteScanParamsMetadataCatalogRequestMetadataTypeCatalog),
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
