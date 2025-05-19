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

func TestSuiPostTransactionScanWithOptionalParams(t *testing.T) {
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
	_, err := client.Sui.PostTransaction.Scan(context.TODO(), blockaidclientgo.SuiPostTransactionScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.SuiPostTransactionScanParamsChainMainnet),
		Data: blockaidclientgo.F(blockaidclientgo.SuiPostTransactionScanParamsData{
			TxHash: blockaidclientgo.F("tx_hash"),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.SuiPostTransactionScanParamsMetadata{
			Domain:  blockaidclientgo.F("domain"),
			NonDapp: blockaidclientgo.F(true),
		}),
		Options: blockaidclientgo.F([]blockaidclientgo.SuiPostTransactionScanParamsOption{blockaidclientgo.SuiPostTransactionScanParamsOptionValidation}),
	})
	if err != nil {
		var apierr *blockaidclientgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
