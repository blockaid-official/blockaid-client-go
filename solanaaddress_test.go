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

func TestSolanaAddressScanWithOptionalParams(t *testing.T) {
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
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Solana.Address.Scan(context.TODO(), blockaidclientgo.SolanaAddressScanParams{
		AddressScanRequestSchema: blockaidclientgo.AddressScanRequestSchemaParam{
			Address: blockaidclientgo.F("2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S"),
			Metadata: blockaidclientgo.F(blockaidclientgo.AddressScanRequestSchemaMetadataParam{
				URL: blockaidclientgo.F("https://example.com"),
			}),
			Chain: blockaidclientgo.F("mainnet"),
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
