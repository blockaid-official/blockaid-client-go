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

func TestTokenScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Token.Scan(context.TODO(), blockaidclientgo.TokenScanParams{
		Address: blockaidclientgo.F[blockaidclientgo.TokenScanParamsAddressUnion](shared.UnionString("0x833110ae77ad7d1f883ffc4f6eb0059143d2754c")),
		Chain:   blockaidclientgo.F(blockaidclientgo.TokenScanSupportedChainEthereum),
		Metadata: blockaidclientgo.F(blockaidclientgo.TokenScanParamsMetadata{
			Domain: blockaidclientgo.F("string"),
		}),
		TokenID: blockaidclientgo.F(int64(0)),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
