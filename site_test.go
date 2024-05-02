// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blockaid_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/blockaid-client-go"
	"github.com/stainless-sdks/blockaid-client-go/internal/testutil"
	"github.com/stainless-sdks/blockaid-client-go/option"
)

func TestSiteScanWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blockaid.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Site.Scan(context.TODO(), blockaid.SiteScanParams{
		URL: blockaid.F("string"),
		Metadata: blockaid.F[blockaid.SiteScanParamsMetadataUnion](blockaid.SiteScanParamsMetadataCatalogRequestMetadata{
			Type: blockaid.F(blockaid.SiteScanParamsMetadataCatalogRequestMetadataTypeCatalog),
		}),
	})
	if err != nil {
		var apierr *blockaid.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
