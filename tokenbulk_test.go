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

func TestTokenBulkScanWithOptionalParams(t *testing.T) {
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
	_, err := client.TokenBulk.Scan(context.TODO(), blockaidclientgo.TokenBulkScanParams{
		Chain:  blockaidclientgo.F(blockaidclientgo.TokenScanSupportedChainArbitrum),
		Tokens: blockaidclientgo.F([]string{"0x66587563e933bbf3974b89156b47bb82b921eb35", "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"}),
		Metadata: blockaidclientgo.F(blockaidclientgo.TokenBulkScanParamsMetadata{
			Domain: blockaidclientgo.F("domain"),
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
